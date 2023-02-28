package services

import (
	"context"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/ports"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
)

// TODO adds comments

type OrderServiceManager struct {
	event.DomainEventRecorder[domain.OrderCreatedEvent]
	event.Bus[domain.OrderCreatedEvent]
}

func NewOrderServiceManager(
	event event.DomainEventRecorder[domain.OrderCreatedEvent], bus event.Bus[domain.OrderCreatedEvent],
) ports.OrderServiceManager {

	return &OrderServiceManager{
		event,
		bus,
	}
}

func (o *OrderServiceManager) OrderCreator(ctx context.Context, order *domain.Order) (err error) {
	evt := order.PullEvents()
	if err = o.Publish(ctx, evt); err != nil {
		if err = o.Record(ctx, &evt[0]); err != nil {
			return
		}
		return
	}

	if err = o.Record(ctx, &evt[0]); err != nil {
		return
	}

	return
}
