package handler_test

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/stretchr/testify/assert"
)

type testCancelEvent struct {
	cancelled bool
}

func (e testCancelEvent) Cancelled() handler.CancellableEvent {
	e.cancelled = true
	return e
}

func TestCancelableEmitNoSubscription(t *testing.T) {
	var h handler.CancelableEventHandler[testCancelEvent]
	assert.False(t, h.Emit(testCancelEvent{cancelled: false}))
}

func TestCancelableListeners(t *testing.T) {
	var h handler.CancelableEventHandler[testCancelEvent]

	h.Subscribe(func(event testCancelEvent) bool {
		assert.Fail(t, "the 2nd priority listener should never have been called")
		return false
	}, 2)

	callCount := 0

	h.Subscribe(func(event testCancelEvent) bool {
		assert.Equal(t, 1, callCount)
		callCount += 1
		return true
	}, 1)

	h.Subscribe(func(event testCancelEvent) bool {
		callCount += 1
		return false
	}, 0)

	h.Emit(testCancelEvent{cancelled: false})
	assert.Equal(t, 2, callCount)
}
