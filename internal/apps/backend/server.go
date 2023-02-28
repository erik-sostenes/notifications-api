package backend

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/apps/backend/handlers"
	"os"
	"strings"
)

const defaultPort = "8082"

// Server contains the configuration of server to start and register all http handlers
type Server struct {
	port     string
	engine   *echo.Echo
	handlers handlers.Handlers
}

// NewServer returns an instance of Server
func NewServer(engine *echo.Echo, handlers handlers.Handlers) *Server {
	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	return &Server{
		port:     port,
		engine:   engine,
		handlers: handlers,
	}
}

// Start initialize the server with all http handlers
func (s *Server) Start() error {
	s.setRoutes()

	return s.engine.Start(fmt.Sprintf(":%v", "8082"))
}

// Routes register all endpoints
//
// configure the middlewares CORS, Logger and Recover
func (s *Server) setRoutes() {
	s.engine.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	group := s.engine.Group("/v1/eat-fast/food-order-api")

	group.GET("/status", HealthCheck())
	group.PUT("/create-order/:id", s.handlers.OrderHandler.Create())
	group.GET("/handshake", s.handlers.OrderNotifier.Notify())
}
