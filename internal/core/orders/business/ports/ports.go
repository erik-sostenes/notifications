package ports

import (
	"context"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/domain"
)

// ports right side
type (
	Saver interface {
		Save(context.Context, *domain.Order) error
	}
)
