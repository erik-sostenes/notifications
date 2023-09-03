package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/services"
	"github.com/erik-sostenes/notifications-api/internal/core/orders/driven/event"
	"github.com/erik-sostenes/notifications-api/pkg/bus/command"
)

func Test_HttpHandlerOrderCreator(t *testing.T) {
	type HandlerFunc func() (http.HandlerFunc, error)

	tsc := map[string]struct {
		*http.Request
		HandlerFunc
		expectedStatusCode int
	}{
		"Given a valid non-existing food order, a status code 201 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/api/v1/orders/create", strings.NewReader(
				`{
					"id": "1e737f50-07f1-4d1b-9c3a-62f4d38559a9",
                    "createAt": "2022-11-21 19:51:39",
					"status": "WAITING",
					"price": {
						"amount": 45.62,
						"currency": "MX"
					},
					"address": {
						"country": "Mexico",
						"state": "HIDALGO",
						"municipality": "Tula de Allende Hidalgo",
						"latitude": 6.5568768,
						"longitude": 3.3488896
					},
					"requestedTime": "2022-11-21 19:51:39",
					"isProduct": true,
					"isSubscription": false,
					"typeSubs": "YEAR",
 					"userId":"c2f91217-de8b-46fa-9168-132fe9285d87"
				}`,
			)),
			HandlerFunc: func() (http.HandlerFunc, error) {
				publisher := event.NewMockDomainEventPublisher("eatfast.order.1.domain_event.order.create_order_event")
				orderCommandHandler := services.NewCreateOrderCommandHandler(&publisher)

				cmd := make(command.CommandBus[services.CreateOrderCommand])

				if err := cmd.Record(services.CreateOrderCommand{}, orderCommandHandler); err != nil {
					return nil, err
				}

				return HttpHandlerOrderCreator(&cmd), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			req := ts.Request
			req.Header.Set("Content-type", "application/json; charset=utf-8")

			resp := httptest.NewRecorder()

			h, err := ts.HandlerFunc()
			if err != nil {
			}

			h.ServeHTTP(resp, req)

			if ts.expectedStatusCode != resp.Code {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}
