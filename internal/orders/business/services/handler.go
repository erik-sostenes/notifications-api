package services

import (
	"context"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/ports"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/command"
)

// CreateOrderCommand implements the command.Command interface
var _ command.Command = CreateOrderCommand{}

type Price struct {
	Amount   string
	Currency string
}

type Address struct {
	Country      string
	State        string
	Municipality string
	Latitude     string
	Longitude    string
}

// CreateOrderCommand represents the DTO with the primitive values
type CreateOrderCommand struct {
	Id               string
	CreateAt         string
	Status           string
	P                Price
	A                Address
	RequestedTime    string
	IsProduct        string
	IsSubscription   string
	TypeSubscription string
	UserId           string
	FoodDishesIds    []string
}

// CommandId returns the command type
func (CreateOrderCommand) CommandId() string {
	return "create_order_command"
}

// CreateOrderCommandHandler implements the command.Handler[CreateOrderCommand] interface
var _ command.Handler[CreateOrderCommand] = (*CreateOrderCommandHandler)(nil)

type CreateOrderCommandHandler struct {
	OrderService ports.OrderServiceManager
}

// Handler executes the action of the command.Command = CreateOrderCommand
func (h CreateOrderCommandHandler) Handler(ctx context.Context, cmd CreateOrderCommand) error {
	address, err := domain.NewAddress(cmd.A.Country, cmd.A.State, cmd.A.Municipality, cmd.A.Latitude, cmd.A.Longitude)
	if err != nil {
		return err
	}

	price, err := domain.NewPrice(cmd.P.Amount, cmd.P.Currency)
	if err != nil {
		return err
	}

	order, err := domain.NewOrder(cmd.Id, cmd.CreateAt, cmd.Status, price, address, cmd.RequestedTime, cmd.IsProduct,
		cmd.IsSubscription, cmd.TypeSubscription, cmd.UserId, cmd.FoodDishesIds)
	if err != nil {
		return err
	}

	return h.OrderService.OrderCreator(ctx, &order)
}
