package db

import (
	"errors"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	t.Run("Given a correct configuration, Redis DB will connect", func(t *testing.T) {
		_, err := NewRedisClient(NewRedisDBConfiguration())

		if !errors.Is(err, nil) {
			t.Errorf("%v error was expected, but %v error was obtained", nil, err)
		}
	})
}
