package event

import (
	"context"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/domain"

	"github.com/erik-sostenes/notifications-api/pkg/bus/event"
	"github.com/erik-sostenes/notifications-api/pkg/set"
)

var _ event.Publisher[domain.OrderCreatedEvent] = (*DomainEventPublisher)(nil)

// MockDomainEventPublisher is used for unit test and integration test mocks
type MockDomainEventPublisher struct {
	set.Set[string, domain.OrderCreatedEvent]
	Stream string
}

func NewMockDomainEventPublisher(stream string) MockDomainEventPublisher {
	return MockDomainEventPublisher{
		*set.NewSet[string, domain.OrderCreatedEvent](),
		stream,
	}
}

// Publish method that publishes all domain events contained in the slice []event.Event
func (e *MockDomainEventPublisher) Publish(ctx context.Context, events []domain.OrderCreatedEvent) (err error) {
	for _, event := range events {
		e.Add(e.Stream, event)
	}

	return
}
