package event_test

import (
	"testing"

	. "github.com/simimpact/srsim/pkg/engine/event"
)

func TestEmitNoSubscription(t *testing.T) {
	var handler EventHandler[int]
	handler.Emit(10)
}

func TestSimpleListener(t *testing.T) {
	var handler EventHandler[int]

	value := 0
	handler.Subscribe(func(event int) {
		value += event
	})

	handler.Emit(4)
	if value != 4 {
		t.Errorf("Value %d does not equal expected 4", value)
	}
}

func TestMultipleListeners(t *testing.T) {
	var handler EventHandler[int]

	value := 0
	handler.Subscribe(func(event int) {
		value += event
	})

	handler.Subscribe(func(event int) {
		value -= event * 2
	})

	handler.Emit(10)
	if value != -10 {
		t.Errorf("Value %d does not equal expected -10", value)
	}
}

