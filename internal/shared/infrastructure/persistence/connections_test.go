package persistence

import (
	"errors"
	"testing"
)

func TestNewMongoClient(t *testing.T) {
	t.Run("Given a correct configuration, MongoDB will connect", func(t *testing.T) {
		_, err := NewMongoClient(NewMongoDBConfiguration())

		if !errors.Is(err, nil) {
			t.Errorf("%v error was expected, but %v error was obtained", nil, err)
		}
	})
}

func TestNewRedisClient(t *testing.T) {
	t.Run("Given a correct configuration, Redis DB will connect", func(t *testing.T) {
		_, err := NewRedisClient(NewRedisDBConfiguration())

		if !errors.Is(err, nil) {
			t.Errorf("%v error was expected, but %v error was obtained", nil, err)
		}
	})
}
