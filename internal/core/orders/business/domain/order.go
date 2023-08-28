package domain

import (
	"context"

	"github.com/erik-sostenes/notifications-api/pkg/bus/event"
	"github.com/erik-sostenes/notifications-api/pkg/common"
)

// Order represents a Value Object with the values of a food order
type Order struct {
	OrderId               OrderId
	OrderCreateAt         OrderCreateAt
	OrderStatus           OrderStatus
	OrderPrice            Price
	OrderAddress          Address
	OrderRequestedTime    OrderRequestedTime
	OrderIsProduct        OrderIsProduct
	OrderIsSubscription   OrderIsSubscription
	OrderTypeSubscription OrderTypeSubscription
	OrderUserId           OrderUserId

	events event.DomainEventRecorderInMemory[OrderCreatedEvent]
}

func NewOrder(
	id,
	createAt,
	status string,
	price Price,
	address Address,
	requestedTime string,
	isProduct,
	isSubscription bool,
	typeSubscription,
	userId string,
) (Order, error) {
	orderId, err := NewOrderId(id)
	if err != nil {
		return Order{}, err
	}

	orderCreateAt, err := NewOrderCreateAt(createAt)
	if err != nil {
		return Order{}, err
	}

	orderStatus := OrderStatus(status)
	if err != nil {
		return Order{}, err
	}

	orderRequestedTime, err := NewOrderRequestedTime(requestedTime)
	if err != nil {
		return Order{}, err
	}

	orderIsProduct, err := NewOrderIsProduct(isProduct)
	if err != nil {
		return Order{}, err
	}

	orderIsSubscription, err := NewOrderIsSubscription(isSubscription)
	if err != nil {
		return Order{}, err
	}

	orderTypeSubscription := OrderTypeSubscription(typeSubscription)
	if err != nil {
		return Order{}, err
	}

	orderUserId, err := NewOrderUserId(userId)
	if err != nil {
		return Order{}, err
	}

	order := Order{
		OrderId:               orderId,
		OrderCreateAt:         orderCreateAt,
		OrderStatus:           orderStatus,
		OrderPrice:            price,
		OrderAddress:          address,
		OrderRequestedTime:    orderRequestedTime,
		OrderIsProduct:        orderIsProduct,
		OrderIsSubscription:   orderIsSubscription,
		OrderTypeSubscription: orderTypeSubscription,
		OrderUserId:           orderUserId,
	}

	order.Record(
		context.Background(),
		NewOrderCreatedEvent(
			orderId,
			orderCreateAt,
			orderStatus,
			price,
			address,
			orderRequestedTime,
			orderIsProduct,
			orderIsSubscription,
			orderTypeSubscription,
			orderUserId,
		),
	)

	return order, err
}

type OrderId struct {
	value string
}

func NewOrderId(value string) (OrderId, error) {
	v, err := common.Identifier(value).Validate()

	return OrderId{v}, err
}

func (o OrderId) Value() string {
	return o.value
}

type OrderCreateAt struct {
	value int64
}

func NewOrderCreateAt(value string) (OrderCreateAt, error) {
	v, err := common.Timestamp(value).Validate()

	return OrderCreateAt{v}, err
}

func (o OrderCreateAt) Value() int64 {
	return o.value
}

type OrderStatus string

const (
	WAITING   OrderStatus = "WAITING"
	ACCEPTED  OrderStatus = "ACCEPTED"
	COMPLETED OrderStatus = "COMPLETED"
)

type OrderRequestedTime struct {
	value int64
}

func NewOrderRequestedTime(value string) (OrderRequestedTime, error) {
	v, err := common.Timestamp(value).Validate()

	return OrderRequestedTime{v}, err
}

func (o OrderRequestedTime) Value() int64 {
	return o.value
}

type OrderIsProduct struct {
	value bool
}

func NewOrderIsProduct(value bool) (OrderIsProduct, error) {
	return OrderIsProduct{value}, nil
}

func (o OrderIsProduct) Value() bool {
	return o.value
}

type OrderIsSubscription struct {
	value bool
}

func NewOrderIsSubscription(value bool) (OrderIsSubscription, error) {
	return OrderIsSubscription{value}, nil
}

func (o OrderIsSubscription) Value() bool {
	return o.value
}

type OrderTypeSubscription string

const (
	ANNUAL  OrderTypeSubscription = "ANNUAL"
	MONTHLY OrderTypeSubscription = "MONTHLY"
)

type OrderUserId struct {
	value string
}

func NewOrderUserId(value string) (OrderUserId, error) {
	v, err := common.Identifier(value).Validate()

	return OrderUserId{v}, err
}

func (o OrderUserId) Value() string {
	return o.value
}

func (o *Order) Record(ctx context.Context, evt OrderCreatedEvent) {
	_ = o.events.Record(ctx, &evt)
}

func (o *Order) PullEvents() []OrderCreatedEvent {
	return o.events.PullEvents()
}
