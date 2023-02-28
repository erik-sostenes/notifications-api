package handlers

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/errors"
	"net/http"
)

// HandlerError handles the type of error
//
// responds with http status code  along with the message
func HandlerError(err error) error {
	switch err.(type) {
	case errors.StatusBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": err})
	case errors.StatusUnprocessableEntity:
		return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{"error": err})
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{"error": err})
	}
}
