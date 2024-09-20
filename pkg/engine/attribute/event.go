package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Service) emitHPChangeEvents(
	key key.Reason, target, source key.TargetID, oldRatio, newRatio, maxHP float64, isDamage bool) error {
	if oldRatio == newRatio {
		return nil
	}

	if isDamage {
		s.targets[target].lastAttacker = source
	}

	s.event.HPChange.Emit(event.HPChange{
		Key:                key,
		Target:             target,
		Source:             source,
		OldHPRatio:         oldRatio,
		NewHPRatio:         newRatio,
		OldHP:              maxHP * oldRatio,
		NewHP:              maxHP * newRatio,
		IsHPChangeByDamage: isDamage,
	})

	if newRatio > 0 {
		s.targets[target].state = info.Alive
		return nil
	}

	s.targets[target].state = info.Dead
	if s.event.LimboWaitHeal.Emit(event.LimboWaitHeal{Target: target}) {
		s.targets[target].state = info.Limbo
	}
	return nil
}

func (s *Service) emitStanceChange(
	key key.Reason, target, source key.TargetID, prevS, newS float64) error {
	s.event.StanceChange.Emit(event.StanceChange{
		Key:       key,
		Target:    target,
		Source:    source,
		OldStance: prevS,
		NewStance: newS,
	})

	return nil
}

func (s *Service) emitEnergyChange(
	key key.Reason, target, source key.TargetID, prevE, newE float64) error {
	if prevE != newE {
		s.event.EnergyChange.Emit(event.EnergyChange{
			Key:       key,
			Target:    target,
			Source:    source,
			OldEnergy: prevE,
			NewEnergy: newE,
		})
	}
	return nil
}

func (s *Service) emitSPChange(key key.Reason, source key.TargetID, prevSP, newSP int) error {
	if prevSP != newSP {
		s.event.SPChange.Emit(event.SPChange{
			Key:    key,
			Source: source,
			OldSP:  prevSP,
			NewSP:  newSP,
		})
	}
	return nil
}
