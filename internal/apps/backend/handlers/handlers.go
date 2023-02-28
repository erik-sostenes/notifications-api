package handlers

import "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/apps/backend/handlers/orders"

// Handlers contains all http handlers
type Handlers struct {
	orders.OrderHandler
	orders.OrderNotifier
}
