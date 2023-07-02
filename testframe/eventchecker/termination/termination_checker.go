package termination

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/testframe/teststub"
)

func ExpectFor() teststub.EventChecker {
	return func(e handler.Event) (bool, error) {
		_, ok := e.(event.Termination)
		if !ok {
			return false, nil
		}
		return true, nil
	}
}
