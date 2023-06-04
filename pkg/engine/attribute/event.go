package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Service) emitHPChangeEvents(target, source key.TargetID, oldRatio, newRatio, maxHP float64) error {
	if oldRatio == newRatio {
		return nil
	}

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

func (s *Service) emitStanceChange(target, source key.TargetID, prev, new float64) error {
	if prev == new {
		return nil
	}

	s.event.StanceChange.Emit(event.StanceChangeEvent{
		Target:    target,
		OldStance: prev,
		NewStance: new,
	})

	if new == 0 {
		s.event.StanceBreak.Emit(event.StanceBreakEvent{
			Target: target,
			Source: source,
		})
	} else if prev == 0 {
		s.event.StanceBreakEnd.Emit(event.StanceBreakEndEvent{
			Target: target,
		})
	}
	return nil
}

func (s *Service) emitEnergyChange(target key.TargetID, prev, new float64) error {
	if prev != new {
		s.event.EnergyChange.Emit(event.EnergyChangeEvent{
			Target:    target,
			OldEnergy: prev,
			NewEnergy: new,
		})
	}
	return nil
}
