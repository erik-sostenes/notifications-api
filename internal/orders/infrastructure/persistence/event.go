package persistence

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/bus/event"
	errors2 "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ event.DomainEventRecorder[domain.OrderCreatedEvent] = EventRecorder{}

type EventRecorder struct {
	CollectionName string
	DB             *mongo.Database
}

// Record method that persists a domain event in mongoDB, the collection where the event will be stored depends
// on whether it was published successfully or not
func (e EventRecorder) Record(ctx context.Context, evt *domain.OrderCreatedEvent) (err error) {
	collection := e.DB.Collection(e.CollectionName)

	if _, err = e.find(ctx, evt.AggregateID()); !errors.Is(err, mongo.ErrNoDocuments) {
		err = errors2.DuplicatedDomainEvent(fmt.Sprintf(
			"The domain event OrderCreatedEvent with aggregate_id '%s' already exists", evt.AggregateID()))
		return
	}

	_, err = collection.InsertOne(ctx, bson.M{
		"event_id":     evt.ID(),
		"type":         evt.Type(),
		"occurred_on":  evt.OccurredOn(),
		"aggregate_id": evt.AggregateID(),
		"data":         evt.Data(),
		"meta_data":    evt.MetaData(),
	})
	return
}

// find method that searches for a domain event by aggregated id, currently only used for integration testing
func (e EventRecorder) find(ctx context.Context, aggregateId string) (_ *domain.OrderCreatedEvent, err error) {
	collection := e.DB.Collection(e.CollectionName)

	filter := bson.M{"aggregate_id": aggregateId}

	err = collection.FindOne(ctx, filter).Decode(bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
	}
	return
}

// delete method that deletes a domain event by event id, currently only used for integration testing
func (e EventRecorder) delete(ctx context.Context, eventId string) (err error) {
	collection := e.DB.Collection(e.CollectionName)

	filter := bson.M{"event_id": eventId}

	_, err = collection.DeleteOne(ctx, filter)
	return
}
