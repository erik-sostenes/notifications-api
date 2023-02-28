package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
	domain2 "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
)

var _ event.Bus[domain.OrderCreatedEvent] = (*DomainEventBus)(nil)

// DomainEventBus implements event.Bus interface
// redis streams are used to publish and consume domain events
type DomainEventBus struct {
	RDB    *redis.Client
	Stream string
}

// Publish method that publishes all domain events contained in the slice []event.Event
func (e *DomainEventBus) Publish(ctx context.Context, events []domain.OrderCreatedEvent) (err error) {
	for _, evt := range events {
		args := redis.XAddArgs{
			Stream: e.Stream,
			Values: map[string]any{
				"event_id":     evt.ID(),
				"type":         evt.Type(),
				"occurred_on":  evt.OccurredOn(),
				"aggregate_id": evt.AggregateID(),
				"data":         evt.Data(),
				"meta_data":    evt.MetaData(),
			},
		}

		if err = e.RDB.XAdd(ctx, &args).Err(); err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

// Consumer method that consumes all domain events only from subscribed streams
// each time a message is received, a handler will be executed concurrently
func (e *DomainEventBus) Consumer(ctx context.Context, events chan<- []byte) {
	streamId := "0"

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				args := redis.XReadArgs{
					Streams: []string{e.Stream, streamId},
					Count:   1,
					Block:   0,
				}

				values, err := e.RDB.XRead(ctx, &args).Result()
				if err != nil {
					messageErr, _ := json.Marshal(domain2.Map{"error": err})
					events <- messageErr
					continue
				}

				for i, v := range values {
					streamId = v.Messages[i].ID
					data, _ := json.Marshal(&v.Messages[i].Values)
					events <- data
				}
			}
		}
	}()
}
