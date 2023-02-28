package orders

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/services"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/command"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
)

// orderHandler implements the OrderHandler interface
var _ OrderHandler = orderHandler{}

// OrderHandler handles http requests and response to manage the processes of a food order
type OrderHandler interface {
	// Create http handler that maps the primitive data from the request body to
	// a command and sends it to a command bus
	Create() echo.HandlerFunc
}

// orderHandler handles the tasks specified by the OrderHandler interface
type orderHandler struct {
	command.Bus[services.CreateOrderCommand]
}

// NewOrderHandler returns an instance of OrderHandler
func NewOrderHandler(bus command.Bus[services.CreateOrderCommand]) OrderHandler {
	return orderHandler{bus}
}

// orderNotifier implements the OrderNotifier interface
var _ OrderNotifier = (*orderNotifier)(nil)

// OrderNotifier manages all notifications in real time
type OrderNotifier interface {
	// Notify method that notifies the client in real time when any message is received
	Notify() echo.HandlerFunc
}

// orderNotifier handles the tasks specified by the OrderNotifier interface
type orderNotifier struct {
	event.Handler
}

// NewOrderNotifier returns an instance of OrderNotifier, injects required values
func NewOrderNotifier(handler event.Handler) OrderNotifier {
	return &orderNotifier{handler}
}
