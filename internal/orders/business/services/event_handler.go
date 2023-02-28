package services

import (
	"context"
	"encoding/json"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
	domain2 "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
	"sync"
	"time"
)

// eventHandler implements the event.Handler interface
var _ event.Handler = (*eventHandler)(nil)

// eventHandler consumes all domain events of the event.Bus
type eventHandler struct {
	sync.Mutex
	event.Bus[domain.OrderCreatedEvent]
}

// NewEventHandler returns an instance of
func NewEventHandler(bus event.Bus[domain.OrderCreatedEvent]) event.Handler {
	return &eventHandler{
		Bus: bus,
	}
}

// Handle consumes all messages, processes the message and sends it to http handler
func (h *eventHandler) Handle(ctx context.Context, message chan<- []byte) {
	events := make(chan []byte)
	defer close(events)

	h.Consumer(ctx, events)

	for evt := range events {
		signalCH := make(chan struct{}, 1)

		signalCH <- struct{}{}
		go func(evt []byte) {
			h.Lock()
			defer h.Unlock()

			var m domain2.Map
			_ = json.Unmarshal(evt, &m)

			data := m["data"]
			by, _ := json.Marshal(&data)
			message <- by

			time.Sleep(time.Second * 1)

			<-signalCH
		}(evt)
		close(signalCH)
	}
}
