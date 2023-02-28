package backend

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// HealthCheck http handler to check server health status
func HealthCheck() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "")
	}
}
