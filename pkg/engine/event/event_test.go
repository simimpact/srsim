package event_test

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event"
)

func TestEventSystem(t *testing.T) {
	var sys event.System

	value := 0
	sys.Ping.Subscribe(func(event int) {
		value += event
	})

	sys.Ping.Emit(10)

	if value != 10 {
		t.Errorf("Expected value %q does not equal 10", value)
	}
}
