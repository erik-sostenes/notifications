package event

import (
	"context"
	"time"

	"github.com/erik-sostenes/notifications-api/pkg/common"
)

type (

	// Publisher publish new events
	Publisher[V Event] interface {
		Publish(context.Context, []V) error
	}

	// Consumer consumes all the events the stream is subscribed to
	Consumer interface {
		Consume(context.Context) error
	}

	// Bus defines the expected behaviour from an event bus
	Bus[V Event] interface {
		Publisher[V]

		Consumer
	}
)

// Handler defines the expected behaviour from an event handler
type Handler interface {
	// Handle method that processes all received common events and sends them to the http handlers
	Handle(ctx context.Context, message []byte)
}

type (
	// Type represents the type of event common examples = eatfast.event.order.created, eatfast.event.order.removed
	Type string
)

// Event interface that implements all domain event must implement
type Event interface {
	// ID method that returns the event domain ID
	ID() string
	// Type method that returns the event domain Type
	Type() Type
	// OccurredOn method that returns the time in unix format when the domain event was created
	OccurredOn() int64
	// AggregateID method that returns the identifier of the DTO(Data Transfer Object) that was added to the domain event
	AggregateID() string
	// Data method that returns the payload of the DTO(Data Transfer Object) that was added to the domain event
	Data() common.Map
	// MetaData method that returns information extra when the domain event was created
	MetaData() common.Map
}

// DomainEvent represents the basis of a domain event, it implements the interfaces Event
type DomainEvent struct {
	eventID     string
	aggregateID string
	occurredOn  int64
	metaData    common.Map
}

// NewDomainEvent returns an instance of DomainEvent
func NewDomainEvent(aggregateID string) DomainEvent {
	return DomainEvent{
		eventID:     common.GenerateUuID(),
		aggregateID: aggregateID,
		occurredOn:  time.Now().Unix(),
		metaData: common.Map{
			"server_name": "eat-fast-food-order-api",
		},
	}
}

func (b DomainEvent) ID() string {
	return b.eventID
}

func (b DomainEvent) AggregateID() string {
	return b.aggregateID
}

func (b DomainEvent) OccurredOn() int64 {
	return b.occurredOn
}

func (b DomainEvent) MetaData() common.Map {
	return b.metaData
}

// DomainEventRecorder interface that will record domain events and store them in a database manager
// or in memory
type DomainEventRecorder[V Event] interface {
	// Record method that records event commons
	Record(context.Context, *V) error
}

// DomainEventRecorderInMemory is a slice of Event that implements the DomainEventRecorder
// interface and records event commons in memory
type DomainEventRecorderInMemory[V Event] []V

// Record method that records event commons in memory
func (e *DomainEventRecorderInMemory[V]) Record(_ context.Context, evt *V) error {
	*e = append(*e, *evt)
	return nil
}

// Flush method that flushes the slice that has the records of all event commons before pull
func (e *DomainEventRecorderInMemory[V]) Flush() {
	*e = make(DomainEventRecorderInMemory[V], 0)
}

// PullEvents method that pulls all event commons from the slice
func (e *DomainEventRecorderInMemory[V]) PullEvents() (evt []V) {
	evt = *e
	e.Flush()
	return evt
}
