package domain

import (
	"context"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
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
	OrderFoodDishesId     []OrderFoodDishesId

	events event.DomainEventRecorderInMemory[OrderCreatedEvent]
}

func NewOrder(id, createAt, status string, price Price, address Address, requestedTime, isProduct, isSubscription,
	typeSubscription, userId string, foodDishesId []string) (Order, error) {

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

	var orderFoodDishesId []OrderFoodDishesId
	for _, v := range foodDishesId {
		id, err := NewOrderFoodDishesId(v)
		if err != nil {
			return Order{}, err
		}
		orderFoodDishesId = append(orderFoodDishesId, id)
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
		OrderFoodDishesId:     orderFoodDishesId,
	}

	order.Record(context.TODO(), NewOrderCreatedEvent(orderId, orderCreateAt, orderStatus, price, address, orderRequestedTime, orderIsProduct,
		orderIsSubscription, orderTypeSubscription, orderUserId, orderFoodDishesId))

	return order, err
}

type OrderId struct {
	Value string
}

func NewOrderId(value string) (OrderId, error) {
	v, err := domain.Identifier(value).Validate()

	return OrderId{v}, err
}

type OrderCreateAt struct {
	Value int64
}

func NewOrderCreateAt(value string) (OrderCreateAt, error) {
	v, err := domain.Timestamp(value).Validate()

	return OrderCreateAt{v}, err
}

type OrderStatus string

const (
	WAITING   OrderStatus = "WAITING"
	ACCEPTED  OrderStatus = "ACCEPTED"
	COMPLETED OrderStatus = "COMPLETED"
)

type OrderRequestedTime struct {
	Value int64
}

func NewOrderRequestedTime(value string) (OrderRequestedTime, error) {
	v, err := domain.Timestamp(value).Validate()

	return OrderRequestedTime{v}, err
}

type OrderIsProduct struct {
	Value bool
}

func NewOrderIsProduct(value string) (OrderIsProduct, error) {
	v, err := domain.Bool(value).Validate()

	return OrderIsProduct{v}, err
}

type OrderIsSubscription struct {
	Value bool
}

func NewOrderIsSubscription(value string) (OrderIsSubscription, error) {
	v, err := domain.Bool(value).Validate()

	return OrderIsSubscription{v}, err
}

type OrderTypeSubscription string

const (
	ANNUAL  OrderTypeSubscription = "ANNUAL"
	MONTHLY OrderTypeSubscription = "MONTHLY"
)

type OrderUserId struct {
	Value string
}

func NewOrderUserId(value string) (OrderUserId, error) {
	v, err := domain.Identifier(value).Validate()

	return OrderUserId{v}, err
}

type OrderFoodDishesId struct {
	Value string
}

func NewOrderFoodDishesId(value string) (OrderFoodDishesId, error) {
	v, err := domain.Identifier(value).Validate()

	return OrderFoodDishesId{v}, err
}

func (o *Order) Record(ctx context.Context, evt OrderCreatedEvent) {
	_ = o.events.Record(ctx, &evt)
}

func (o *Order) PullEvents() []OrderCreatedEvent {
	return o.events.PullEvents()
}
