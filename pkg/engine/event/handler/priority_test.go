package handler_test

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
)

func TestPriorityEmitNoSubscription(t *testing.T) {
	var h handler.PriorityEventHandler[int]
	h.Emit(10)
}

func TestPriorityListeners(t *testing.T) {
	var h handler.PriorityEventHandler[int]

	value := 0
	h.Subscribe(func(event int) {
		value = event / value
	}, 2)

	h.Subscribe(func(event int) {
		value += event
	}, 1)

	h.Emit(4)
	if value != 1 {
		t.Errorf("Value %d does not equal expected 1", value)
	}
}
