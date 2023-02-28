package ports

import (
	"context"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
)

// TODO adds comments

type (
	// OrderServiceManager .
	OrderServiceManager interface {
		OrderCreator(ctx context.Context, order *domain.Order) error
	}
)
