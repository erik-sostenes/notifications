package event

import (
	"context"
	"testing"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/domain"
	"github.com/erik-sostenes/notifications-api/pkg/bus/event"
	"github.com/erik-sostenes/notifications-api/pkg/db"
	"github.com/redis/go-redis/v9"
)

func Test_Publisher(t *testing.T) {
	type OrderFunc func() (domain.Order, error)

	tsc := map[string]struct {
		OrderFunc
		streamName string
		event.Publisher[domain.OrderCreatedEvent]
		rdb           *redis.Client
		expectedError error
	}{
		"Given a valid domain event, it will be published correctly": {
			OrderFunc: func() (order domain.Order, err error) {
				price, err := domain.NewPrice(45.62, "MX")
				if err != nil {
					return
				}

				address, err := domain.NewAddress("México", "HIDALGO", "Tula de Allende Hidalgo", 6.5568768, 3.3488896)
				if err != nil {
					return
				}

				return domain.NewOrder(
					"1e737f50-07f1-4d1b-9c3a-62f4d38559a9",
					"2022-11-21 19:51:39",
					"WAITING",
					price,
					address,
					"2022-11-21 19:51:39",
					true,
					false,
					"YEAR",
					"c2f91217-de8b-46fa-9168-132fe9285d87",
				)
			},
			streamName: "test.order.1.domain_event.order.create_order_event",
			rdb:        db.NewRedisDataBase(db.NewRedisDBConfiguration()),
			Publisher:  NewDomainEventPublisher(db.NewRedisDataBase(db.NewRedisDBConfiguration()), "test.order.1.domain_event.order.create_order_event"),
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			order, err := ts.OrderFunc()
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				_ = ts.rdb.Del(context.Background(), ts.streamName)
			})

			evts := order.PullEvents()
			if ts.expectedError != ts.Publisher.Publish(context.Background(), evts) {
				t.Fatalf("%v error was expected, but it was obtained %v", ts.expectedError, err)
			}

			values, err := ts.rdb.XRange(context.Background(), ts.streamName, "-", "+").Result()
			if err != nil {
				t.Fatal(err)
			}

			for _, v := range values {
				expected := evts[0].AggregateID()
				got := v.Values["aggregate_id"]

				if expected != got {
					t.Errorf("aggregate id was expected %v, but it was obtained %d", expected, got)
				}
			}
		})
	}
}
