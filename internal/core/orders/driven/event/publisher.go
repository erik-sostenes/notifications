package event

import (
	"context"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/domain"
	"github.com/erik-sostenes/notifications-api/pkg/bus/event"
	"github.com/erik-sostenes/notifications-api/pkg/common"
	"github.com/redis/go-redis/v9"
)

var _ event.Publisher[domain.OrderCreatedEvent] = (*DomainEventPublisher)(nil)

// DomainEventPublisher implements event.Bus interface
// redis streams are used to publish and consume domain events
type DomainEventPublisher struct {
	RDB    *redis.Client
	Stream string
}

func NewDomainEventPublisher(rdb *redis.Client, stream string) *DomainEventPublisher {
	return &DomainEventPublisher{
		rdb,
		stream,
	}
}

// Publish method that publishes all domain events contained in the slice []event.Event
func (e *DomainEventPublisher) Publish(ctx context.Context, events []domain.OrderCreatedEvent) (err error) {
	for _, evt := range events {
		args := redis.XAddArgs{
			Stream: e.Stream,
			Values: map[string]any{
				"event_id":     evt.ID(),
				"type":         common.NewMarshalJSON(evt.Type()),
				"occurred_on":  evt.OccurredOn(),
				"aggregate_id": evt.AggregateID(),
				"data":         common.NewMarshalJSON(evt.Data()),
				"meta_data":    common.NewMarshalJSON(evt.MetaData()),
			},
		}

		if err = e.RDB.XAdd(ctx, &args).Err(); err != nil {
			return
		}
	}
	return
}
