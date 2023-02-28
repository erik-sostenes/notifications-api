package domain

import (
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
)

// OrderCreatedEvent implements the event.Event interface
var _ event.Event = OrderCreatedEvent{}

// OrderCreatedEventType represents the type of domain event
const OrderCreatedEventType event.Type = "eatfast.event.order.created"

// OrderCreatedEvent represents a domain event when a new order has been created,
// is composed of event.DomainEvent which represents the base of an event.Event
type OrderCreatedEvent struct {
	event.DomainEvent
	data domain.Map
}

// NewOrderCreatedEvent adds the data and returns a new instance of OrderCreatedEvent
func NewOrderCreatedEvent(id OrderId, createAt OrderCreateAt, status OrderStatus, price Price, address Address,
	requestedTime OrderRequestedTime, isProduct OrderIsProduct, isSubscription OrderIsSubscription,
	typeSubscription OrderTypeSubscription, userId OrderUserId, foodDishesId []OrderFoodDishesId,
) OrderCreatedEvent {

	return OrderCreatedEvent{
		DomainEvent: event.NewDomainEvent(id.Value),
		data: domain.Map{
			"id":        id.Value,
			"create_at": createAt.Value,
			"status":    status,
			"price": domain.Map{
				"amount":   price.PriceAmount.Value(),
				"currency": price.PriceCurrency,
			},
			"address": domain.Map{
				"id":           address.AddressId.Value,
				"country":      address.AddressCountry.Value,
				"state":        address.AddressLongitude.Value,
				"municipality": address.AddressMunicipality.Value,
				"latitude":     address.AddressLatitude.Value,
				"longitude":    address.AddressLongitude.Value,
			},
			"requested_time":    requestedTime.Value,
			"is_product":        isProduct.Value,
			"is_subscription":   isSubscription.Value,
			"type_subscription": typeSubscription,
			"user_id":           userId.Value,
			"food_dishes_id":    foodDishesId,
		},
	}
}

func (e OrderCreatedEvent) Type() event.Type {
	return OrderCreatedEventType
}

func (e OrderCreatedEvent) Data() domain.Map {
	return e.data
}
