package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Service) emitHPChangeEvents(
	target, source key.TargetID, oldRatio, newRatio, maxHP float64, isDamage bool) error {
	if oldRatio == newRatio {
		return nil
	}

	s.event.HPChange.Emit(event.HPChange{
		Target:             target,
		OldHPRatio:         oldRatio,
		NewHPRatio:         newRatio,
		OldHP:              maxHP * oldRatio,
		NewHP:              maxHP * newRatio,
		IsHPChangeByDamage: isDamage,
	})

	if newRatio > 0 {
		return nil
	}

	// if event gets canceled, do not want to emit the death event
	if s.event.LimboWaitHeal.Emit(event.LimboWaitHeal{Target: target}) {
		return nil
	}

	s.event.TargetDeath.Emit(event.TargetDeath{
		Target: target,
		Killer: source,
	})
	return nil
}

func (s *Service) emitStanceChange(target, source key.TargetID, prevS, newS float64) error {
	if prevS == newS {
		return nil
	}

	s.event.StanceChange.Emit(event.StanceChange{
		Target:    target,
		OldStance: prevS,
		NewStance: newS,
	})

	if newS == 0 {
		s.event.StanceBreak.Emit(event.StanceBreak{
			Target: target,
			Source: source,
		})
	} else if prevS == 0 {
		s.event.StanceReset.Emit(event.StanceReset{
			Target: target,
		})
	}
	return nil
}

func (s *Service) emitEnergyChange(target key.TargetID, prevE, newE float64) error {
	if prevE != newE {
		s.event.EnergyChange.Emit(event.EnergyChange{
			Target:    target,
			OldEnergy: prevE,
			NewEnergy: newE,
		})
	}
	return nil
}
