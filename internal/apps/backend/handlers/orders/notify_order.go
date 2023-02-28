package orders

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Notify manages notifications in real time, every time you receive a message
// SERVER-SENT EVENTS (SSE) protocol is used to notify the client
func (n *orderNotifier) Notify() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		flusher, ok := c.Response().Writer.(http.Flusher)
		if !ok {
			log.Println("server sent events not supported")
			return
		}

		message := make(chan []byte)

		defer func() {
			close(message)
			message = nil

			log.Println("Client close connection")
		}()

		go n.Handler.Handle(c.Request().Context(), message)

		for {
			select {
			case v := <-message:
				id := rand.NewSource(time.Now().Unix()).Int63()

				io.WriteString(c.Response().Writer, fmt.Sprintf("id: %v\nevent: handshake\ndata: %v", id, string(v)))
				io.WriteString(c.Response().Writer, "\n\n")

				flusher.Flush()
			case <-c.Request().Context().Done():
				return
			case <-time.After(1 * time.Second):
			}
		}
	}
}
