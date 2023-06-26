package battlestart

import (
	"errors"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/testframe/teststub"
)

func ExpectOnly() teststub.EventChecker {
	return func(e handler.Event) (bool, error) {
		_, ok := e.(event.BattleStartEvent)
		if !ok {
			return false, errors.New("incorrect Event")
		}
		return true, nil
	}
}
