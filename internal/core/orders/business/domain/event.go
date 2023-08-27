package domain

import (
	"github.com/erik-sostenes/notifications-api/pkg/common"
	"github.com/erik-sostenes/notifications-api/pkg/domain/bus/event"
)

// OrderCreatedEvent implements the event.Event interface
var _ event.Event = OrderCreatedEvent{}

// OrderCreatedEventType represents the type of domain event
const OrderCreatedEventType event.Type = "eatfast.event.order.created"

// OrderCreatedEvent represents a domain event when a new order has been created,
// is composed of event.DomainEvent which represents the base of an event.Event
type OrderCreatedEvent struct {
	event.DomainEvent
	data common.Map
}

// NewOrderCreatedEvent adds the data and returns a new instance of OrderCreatedEvent
func NewOrderCreatedEvent(
	id OrderId,
	createAt OrderCreateAt,
	status OrderStatus,
	price Price,
	address Address,
	requestedTime OrderRequestedTime,
	isProduct OrderIsProduct,
	isSubscription OrderIsSubscription,
	typeSubscription OrderTypeSubscription,
	userId OrderUserId,
) OrderCreatedEvent {
	return OrderCreatedEvent{
		DomainEvent: event.NewDomainEvent(id.Value()),
		data: common.Map{
			"id":        id.Value(),
			"create_at": createAt.Value(),
			"status":    status,
			"price": common.Map{
				"amount":   price.PriceAmount.Value(),
				"currency": price.PriceCurrency,
			},
			"address": common.Map{
				"id":           address.AddressId.Value(),
				"country":      address.AddressCountry.Value(),
				"state":        address.AddressLongitude.Value(),
				"municipality": address.AddressMunicipality.Value(),
				"latitude":     address.AddressLatitude.Value(),
				"longitude":    address.AddressLongitude.Value(),
			},
			"requested_time":    requestedTime.Value(),
			"is_product":        isProduct.Value(),
			"is_subscription":   isSubscription.Value(),
			"type_subscription": typeSubscription,
			"user_id":           userId.Value(),
		},
	}
}

func (e OrderCreatedEvent) Type() event.Type {
	return OrderCreatedEventType
}

func (e OrderCreatedEvent) Data() common.Map {
	return e.data
}
