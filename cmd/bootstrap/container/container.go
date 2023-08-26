package container

import (
	"github.com/erik-sostenes/notifications-api/pkg/server/health"
	m "github.com/erik-sostenes/notifications-api/pkg/server/middlewares"
	"github.com/erik-sostenes/notifications-api/pkg/server/route"
)

func Injector() (*route.RouteGroup, error) {
	group := route.NewGroup("/api/v1", m.CORS, m.Logger, m.Recovery)
	group.GET("/health", health.HealthCheck())

	return group, nil
}
