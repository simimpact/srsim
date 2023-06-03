package handler_test

import (
	"testing"

	. "github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/stretchr/testify/assert"
)

func TestCancelableEmitNoSubscription(t *testing.T) {
	var handler CancelableEventHandler[int]
	assert.False(t, handler.Emit(10))
}

func TestCancelableListeners(t *testing.T) {
	var handler CancelableEventHandler[int]

	handler.Subscribe(func(event int) bool {
		assert.Fail(t, "the 2nd priority listener should never have been called")
		return false
	}, 2)

	callCount := 0

	handler.Subscribe(func(event int) bool {
		assert.Equal(t, 1, callCount)
		callCount += 1
		return true
	}, 1)

	handler.Subscribe(func(event int) bool {
		callCount += 1
		return false
	}, 0)

	handler.Emit(4)
	assert.Equal(t, 2, callCount)
}
