package turnend

import (
	"github.com/simimpact/srsim/akivali/eventchecker"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
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
