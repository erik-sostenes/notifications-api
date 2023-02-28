package orders

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/services"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/backend/handlers"
	"net/http"
)

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Address struct {
	Country      string `json:"country"`
	State        string `json:"state"`
	Municipality string `json:"municipality"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

// OrderRequest contains the body of http request
type OrderRequest struct {
	Id               string `param:"id"`
	CreateAt         string `json:"create_at"`
	Status           string `json:"status"`
	Price            `json:"price"`
	Address          `json:"address"`
	RequestedTime    string   `json:"requested_time"`
	IsProduct        string   `json:"is_product"`
	IsSubscription   string   `json:"is_subscription"`
	TypeSubscription string   `json:"type_subs"`
	UserId           string   `json:"user_id"`
	FoodDishesIds    []string `json:"food_dishes_ids"`
}

// Create a http handler that receives the body of the http request
// and binds it to the DTO(Data Transfer Object) OrderRequest and dispatches it to a command bus
func (h orderHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := OrderRequest{}
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		command := services.CreateOrderCommand{
			Id:       request.Id,
			CreateAt: request.CreateAt,
			Status:   request.Status,
			P: services.Price{
				Amount:   request.Price.Amount,
				Currency: request.Price.Currency,
			},
			A: services.Address{
				Country:      request.Address.Country,
				State:        request.Address.State,
				Municipality: request.Address.Municipality,
				Latitude:     request.Address.Latitude,
				Longitude:    request.Address.Longitude,
			},
			RequestedTime:    request.RequestedTime,
			IsProduct:        request.IsProduct,
			IsSubscription:   request.IsSubscription,
			TypeSubscription: request.TypeSubscription,
			UserId:           request.UserId,
			FoodDishesIds:    request.FoodDishesIds,
		}

		err := h.Dispatch(c.Request().Context(), command)
		if err != nil {
			return handlers.HandlerError(err)
		}

		return c.JSON(http.StatusCreated, "order created")
	}
}
