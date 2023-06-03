package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Service) emitHPChangeEvents(
	target, source key.TargetID, oldRatio float64, newRatio float64, maxHP float64) error {
	s.event.HPChange.Emit(event.HPChangeEvent{
		Target:     target,
		OldHPRatio: oldRatio,
		NewHPRatio: newRatio,
		OldHP:      maxHP * oldRatio,
		NewHP:      maxHP * newRatio,
	})

	if newRatio > 0 {
		return nil
	}

	// if event gets canceled, do not want to emit the death event
	if s.event.LimboWaitHeal.Emit(event.LimboWaitHealEvent{Target: target}) {
		return nil
	}

	s.event.TargetDeath.Emit(event.TargetDeathEvent{
		Target: target,
		Killer: source,
	})
	return nil
}
