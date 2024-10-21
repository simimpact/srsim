package handler_test

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
)

type mutableTestEvent struct {
	value int
}

func TestMutableEmitNoSubscription(t *testing.T) {
	var h handler.MutableEventHandler[mutableTestEvent]
	x := mutableTestEvent{value: 10}
	h.Emit(&x)
}

func TestMutableListeners(t *testing.T) {
	var h handler.MutableEventHandler[mutableTestEvent]

	h.Subscribe(func(event *mutableTestEvent) {
		event.value = 0
	}, 2)

	h.Subscribe(func(event *mutableTestEvent) {
		event.value *= event.value
	}, 1)

	evt := mutableTestEvent{value: 10}
	h.Emit(&evt)
	if evt.value != 0 {
		t.Errorf("Value %d does not equal expected 1", evt.value)
	}
}
