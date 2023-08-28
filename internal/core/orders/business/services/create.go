package services

import (
	"context"

	"github.com/erik-sostenes/notifications-api/internal/core/orders/business/domain"
	"github.com/erik-sostenes/notifications-api/pkg/bus/command"
	"github.com/erik-sostenes/notifications-api/pkg/bus/event"
)

// CreateOrderCommand implements the command.Command interface
var _ command.Command = CreateOrderCommand{}

type Price struct {
	Amount   float64
	Currency string
}

type Address struct {
	Country      string
	State        string
	Municipality string
	Latitude,
	Longitude float64
}

// CreateOrderCommand represents the DTO with the primitive values
type CreateOrderCommand struct {
	Id       string
	CreateAt string
	Status   string
	Price
	Address
	RequestedTime string
	IsProduct,
	IsSubscription bool
	TypeSubscription string
	UserId           string
}

// CommandId returns the command type
func (CreateOrderCommand) CommandId() string {
	return "create_order_command"
}

// CreateOrderCommandHandler implements the command.Handler[CreateOrderCommand] interface
var _ command.Handler[CreateOrderCommand] = (*CreateOrderCommandHandler)(nil)

type CreateOrderCommandHandler struct {
	event.Publisher[domain.OrderCreatedEvent]
}

func NewCreateOrderCommandHandler(pub event.Publisher[domain.OrderCreatedEvent]) CreateOrderCommandHandler {
	return CreateOrderCommandHandler{
		Publisher: pub,
	}
}

// Handler executes the action of the command.Command = CreateOrderCommand
func (h CreateOrderCommandHandler) Handler(ctx context.Context, cmd *CreateOrderCommand) (err error) {
	price, err := domain.NewPrice(cmd.Price.Amount, cmd.Price.Currency)
	if err != nil {
		return
	}

	address, err := domain.NewAddress(cmd.Address.Country, cmd.Address.State, cmd.Address.Municipality, cmd.Address.Latitude, cmd.Address.Longitude)
	if err != nil {
		return
	}

	order, err := domain.NewOrder(cmd.Id, cmd.CreateAt, cmd.Status, price, address, cmd.RequestedTime, cmd.IsProduct, cmd.IsSubscription, cmd.TypeSubscription, cmd.UserId)
	if err != nil {
		return
	}

	return h.Publish(ctx, order.PullEvents())
}
