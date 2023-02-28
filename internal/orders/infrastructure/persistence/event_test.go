package persistence

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/business/domain"
	d "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/orders/shared/business/domain"
	e "gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/business/domain/errors"
	"gitlab.com/eat-fast/back-end/eatfast-food-order-api/internal/shared/infrastructure/persistence"
	"testing"
)

type newOrderFunc func() (domain.Order, error)

const nameCollection = "recordTestDomainEvents"

func TestEventRecorder_Record(t *testing.T) {
	tsc := map[string]struct {
		EventRecorder
		newOrderFunc
		expectedError error
	}{
		"Given a valid non-existing domain event 'OrderCreatedEvent', it will be registered in the mongo collection": {
			EventRecorder: EventRecorder{
				nameCollection,
				persistence.NewMongoDataBase(persistence.NewMongoDBConfiguration()),
			},
			newOrderFunc: func() (domain.Order, error) {
				return d.CreateOrderRequestMother{}.RandomOrderOne()
			},
			expectedError: nil,
		},
		"Given an existing domain event 'OrderCreatedEvent', it will not be registered in the mongo collection": {
			EventRecorder: EventRecorder{
				nameCollection,
				persistence.NewMongoDataBase(persistence.NewMongoDBConfiguration()),
			},
			newOrderFunc: func() (domain.Order, error) {
				return d.CreateOrderRequestMother{}.RandomOrderTwo()
			},
			expectedError: e.DuplicatedDomainEvent(fmt.Sprintf(
				"The domain event OrderCreatedEvent with aggregate_id '%s' already exists", "1e737f50-07f1-4d1b-9c3a-62f4d38559a9")),
		},
	}

	newEventDomainDB := EventRecorder{
		nameCollection,
		persistence.NewMongoDataBase(persistence.NewMongoDBConfiguration()),
	}

	// SetUp prepare configuration before running integration tests all
	if err := func() error {
		order, err := d.CreateOrderRequestMother{}.RandomOrderTwo()
		if err != nil {
			return err
		}
		return newEventDomainDB.Record(context.TODO(), &order.PullEvents()[0])
	}(); err != nil {
		t.Fatal(err)
	}

	// Teardown reset configuration after running integration tests all
	t.Cleanup(func() {
		_ = newEventDomainDB.DB.Collection(nameCollection).Drop(context.TODO())
	})

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			order, err := ts.newOrderFunc()
			if err != nil {
				t.Error(err)
				t.SkipNow()
			}

			domainEvent := order.PullEvents()[0]

			t.Cleanup(func() {
				_ = ts.delete(context.TODO(), domainEvent.ID())
			})

			err = ts.Record(context.TODO(), &domainEvent)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("%v error was expected, but %v error was obtained", ts.expectedError, err)
			}
		})
	}
}
