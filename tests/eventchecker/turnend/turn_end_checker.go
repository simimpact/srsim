package turnend

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/tests/eventchecker"
)

func ExpectFor() eventchecker.EventChecker {
	return func(e handler.Event) (bool, error) {
		_, ok := e.(event.TurnEnd)
		if !ok {
			return false, nil
		}
		return true, nil
	}
}
