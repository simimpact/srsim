package event_test

import (
	"testing"

	. "github.com/simimpact/srsim/pkg/engine/event"
)

func TestPriorityEmitNoSubscription(t *testing.T) {
	var handler PriorityEventHandler[int]
	handler.Emit(10)
}

func TestPriorityListeners(t *testing.T) {
	var handler PriorityEventHandler[int]

	value := 0
	handler.Subscribe(func(event int) {
		value = event / value
	}, 2)

	handler.Subscribe(func(event int) {
		value += event
	}, 1)

	handler.Emit(4)
	if value != 1 {
		t.Errorf("Value %d does not equal expected 1", value)
	}
}