package turnstart

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/testframe/eventchecker"
)

func ExpectFor() eventchecker.EventChecker {
	return func(e handler.Event) (bool, error) {
		_, ok := e.(event.TurnStart)
		if !ok {
			return false, nil
		}
		return true, nil
	}
}

func CurrentTurnIs(id key.TargetID) eventchecker.EventChecker {
	return func(e handler.Event) (bool, error) {
		ev, ok := e.(event.TurnStart)
		if !ok {
			return false, fmt.Errorf("incorrect Event")
		}
		if ev.Active != id {
			return false, fmt.Errorf("expected: %d, current turn is %d", id, ev.Active)
		}
		return true, nil
	}
}
