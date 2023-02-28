package dependency

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/apps/backend"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/apps/backend/handlers"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/apps/backend/handlers/orders"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/services"
	m "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/infrastructure/persistence"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/command"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/infrastructure/persistence"
)

const streamName = "eatfast.order.1.domain_event.order.create_order_event"
const failurePublishDomainEvent = "failure_publish_domainEvent"
const domainEventPublished = "domainEvent_published"

// NewInjector injects all the dependencies to start the app
func NewInjector() (err error) {
	engine := echo.New()

	rdb := persistence.NewRedisDataBase(persistence.NewRedisDBConfiguration())
	mongoDB := persistence.NewMongoDataBase(persistence.NewMongoDBConfiguration())

	evtBus := m.DomainEventBus{
		RDB:    rdb,
		Stream: streamName,
	}

	evtRecord := m.EventRecorder{
		CollectionName: "domainEvent_published",
		DB:             mongoDB,
	}

	orderHandler, err := injectsOrderHandlerDependencies(evtRecord, evtBus)
	if err != nil {
		return
	}

	orderNotifier := injectsOrderNotifierDependencies(evtBus)

	h := handlers.Handlers{
		OrderHandler:  orderHandler,
		OrderNotifier: orderNotifier,
	}

	return backend.NewServer(engine, h).Start()
}

func injectsOrderHandlerDependencies(evtRecord m.EventRecorder, evtBus m.DomainEventBus) (orders.OrderHandler, error) {
	commandHandler := services.CreateOrderCommandHandler{
		OrderService: services.NewOrderServiceManager(evtRecord, &evtBus),
	}

	commandBus := make(command.CommandBus[services.CreateOrderCommand])

	err := commandBus.RegisterHandler(services.CreateOrderCommand{}, commandHandler)
	if err != nil {
		return nil, err
	}

	return orders.NewOrderHandler(&commandBus), nil
}

func injectsOrderNotifierDependencies(evtBus m.DomainEventBus) orders.OrderNotifier {
	evtHandler := services.NewEventHandler(&evtBus)
	return orders.NewOrderNotifier(evtHandler)
}
