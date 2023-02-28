package orders

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/services"
	m "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/infrastructure/persistence"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/command"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/infrastructure/persistence"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const streamName = "eatfast.order.1.domain_event.order.create_order_event"
const failurePublishDomainEvent = "failure_publish_domainEvent"
const domainEventPublished = "domainEvent_published"

func TestOrderHandler_Create(t *testing.T) {
	var tsc = map[string]struct {
		*http.Request
		OrderHandler
		expectedStatusCode int
	}{
		"Given a valid non-existing food order, a status code 201 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/v1/eat-fast/food-order-api/create-order/1e737f50-07f1-4d1b-9c3a-62f4d38559a9",
				strings.NewReader(
					`{
							"create_at": "2022-11-21 19:51:39",
							"status": "WAITING",
							"price": {
								"amount": "45.62",
								"currency": "MX"
							},
							"address": {
								"country": "Mexico",
								"state": "HIDALGO",
								"municipality": "Tula de Allende Hidalgo",
								"latitude": "6.5568768",
								"longitude": "3.3488896"
							},
							"requested_time": "2022-11-21 19:51:39",
							"is_product": "true",
							"is_subscription": "false",
							"type_subs": "YEAR",
 							"user_id":"c2f91217-de8b-46fa-9168-132fe9285d87",
							"food_dishes_ids": ["d3527262-f415-41c8-9aee-38812af4e484", "d3527262-f415-41c8-9aee-38812af4e484"]
					}`,
				),
			),
			OrderHandler: func() OrderHandler {
				rdb := persistence.NewRedisDataBase(persistence.NewRedisDBConfiguration())
				mongoDB := persistence.NewMongoDataBase(persistence.NewMongoDBConfiguration())

				evtBus := m.DomainEventBus{
					RDB:    rdb,
					Stream: streamName,
				}

				evtRecord := m.EventRecorder{
					CollectionName: domainEventPublished,
					DB:             mongoDB,
				}

				commandBus := make(command.CommandBus[services.CreateOrderCommand])
				err := commandBus.RegisterHandler(services.CreateOrderCommand{}, services.CreateOrderCommandHandler{
					OrderService: services.NewOrderServiceManager(evtRecord, &evtBus),
				})
				if err != nil {
					panic(err)
				}

				return NewOrderHandler(&commandBus)
			}(),
			expectedStatusCode: http.StatusCreated,
		},
		"Given an invalid non-existing food order, a status code 422 is expected": {
			Request: httptest.NewRequest(http.MethodPut, "/v1/eat-fast/food-order-api/create-order/1e737f50-07f1-4d1b-9c3a-62f4d38559a9",
				strings.NewReader(
					`{
							"create_at": "2022-11-21 19:51:39",
							"status": "WAITING",
							"price": {
								"amount": "45.62",
								"currency": "MX"
							},
							"requested_time": "2022-11-21 19:51:39",
							"is_product": "true",
							"is_subscription": "false",
							"type_subs": "YEAR",
 							"user_id":"c2f91217-de8b-46fa-9168-132fe9285d87",
							"food_dishes_ids": ["d3527262-f415-41c8-9aee-38812af4e484", "d3527262-f415-41c8-9aee-38812af4e484"]
					}`,
				),
			),
			OrderHandler: func() OrderHandler {
				commandBus := make(command.CommandBus[services.CreateOrderCommand])
				err := commandBus.RegisterHandler(services.CreateOrderCommand{}, services.CreateOrderCommandHandler{})
				if err != nil {
					panic(err)
				}

				return NewOrderHandler(&commandBus)
			}(),
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			e.PUT("/v1/eat-fast/food-order-api/create-order/:id", ts.OrderHandler.Create())

			resp := httptest.NewRecorder()
			req := ts.Request
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e.ServeHTTP(resp, req)

			if resp.Code != ts.expectedStatusCode {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}
