package container

import (
	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/services"
	"github.com/erik-sostenes/notifications-api/internal/core/orders/driven/event"
	"github.com/erik-sostenes/notifications-api/internal/core/orders/drives/handlers"
	"github.com/erik-sostenes/notifications-api/pkg/bus/command"
	"github.com/erik-sostenes/notifications-api/pkg/db"
	"github.com/erik-sostenes/notifications-api/pkg/server/health"
	m "github.com/erik-sostenes/notifications-api/pkg/server/middlewares"
	"github.com/erik-sostenes/notifications-api/pkg/server/route"
)

const streamName = "eatfast.order.1.domain_event.order.create_order_event"

func Injector() (*route.RouteGroup, error) {
	rdb := db.NewRedisDataBase(db.NewRedisDBConfiguration())

	publisher := event.NewDomainEventPublisher(rdb, streamName)

	orderCommandHandler := services.NewCreateOrderCommandHandler(&publisher)

	cmd := make(command.CommandBus[services.CreateOrderCommand])

	if err := cmd.Record(services.CreateOrderCommand{}, orderCommandHandler); err != nil {
		return nil, err
	}

	routes := route.NewGroup("/api/v1/orders", m.CORS, m.Logger, m.Recovery)
	routes.GET("/health", health.HealthCheck())
	routes.PUT("/create", handlers.HttpHandlerOrderCreator(&cmd))

	return routes, nil
}
