package battlestart

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/testframe/teststub"
)

func ExpectOnly() teststub.EventChecker {
	return func(e handler.Event) (bool, error) {
		_, ok := e.(event.BattleStart)
		if !ok {
			return false, fmt.Errorf("incorrect Event %T", e)
		}
		return true, nil
	}
}

func ExpectFor() teststub.EventChecker {
	return func(e handler.Event) (bool, error) {
		_, ok := e.(event.BattleStart)
		if !ok {
			return false, nil
		}
		return true, nil
	}
}
