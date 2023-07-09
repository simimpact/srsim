package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (s *Service) SetHP(target, source key.TargetID, amt float64, isDamage bool) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	oldRatio := attr.HPRatio
	stats := s.Stats(target)
	attr.HPRatio = amt / stats.MaxHP()

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	} else if attr.HPRatio < 0 {
		attr.HPRatio = 0
	}

	return s.emitHPChangeEvents(target, source, oldRatio, attr.HPRatio, stats.MaxHP(), isDamage)
}

func (s *Service) ModifyHPByAmount(target, source key.TargetID, amt float64, isDamage bool) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	oldRatio := attr.HPRatio
	stats := s.Stats(target)

	newHP := stats.CurrentHP() + amt
	attr.HPRatio = newHP / stats.MaxHP()

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	} else if attr.HPRatio < 0 {
		attr.HPRatio = 0
	}

	return s.emitHPChangeEvents(target, source, oldRatio, attr.HPRatio, stats.MaxHP(), isDamage)
}

func (s *Service) ModifyHPByRatio(target, source key.TargetID, data info.ModifyHPByRatio, isDamage bool) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	oldRatio := attr.HPRatio

	switch data.RatioType {
	case model.ModifyHPRatioType_CURRENT_HP:
		attr.HPRatio += data.Ratio * attr.HPRatio
	case model.ModifyHPRatioType_MAX_HP:
		attr.HPRatio += data.Ratio
	default:
		return fmt.Errorf("unknown ratio type: %v", data.RatioType)
	}

	stats := s.Stats(target)
	if stats.CurrentHP() < data.Floor {
		return s.SetHP(target, source, data.Floor, isDamage)
	}

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	}

	return s.emitHPChangeEvents(target, source, oldRatio, attr.HPRatio, stats.MaxHP(), isDamage)
}

func (s *Service) SetStance(target, source key.TargetID, amt float64) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	prev := attr.Stance
	attr.Stance = amt
	if attr.Stance > attr.MaxStance {
		attr.Stance = attr.MaxStance
	} else if attr.Stance < 0 {
		attr.Stance = 0
	}

	return s.emitStanceChange(target, source, prev, attr.Stance)
}

func (s *Service) ModifyStance(target, source key.TargetID, amt float64) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	stats := s.Stats(target)
	newStance := attr.Stance + amt*(1+stats.GetProperty(prop.AllStanceDMGPercent))
	return s.SetStance(target, source, newStance)
}

func (s *Service) SetEnergy(target key.TargetID, amt float64) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	prev := attr.Energy
	attr.Energy = amt
	if attr.Energy > attr.MaxEnergy {
		attr.Energy = attr.MaxEnergy
	} else if attr.Energy < 0 {
		attr.Energy = 0
	}

	return s.emitEnergyChange(target, prev, attr.Energy)
}

func (s *Service) ModifyEnergy(target key.TargetID, amt float64) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	attr := t.attributes

	stats := s.Stats(target)
	return s.SetEnergy(target, attr.Energy+amt*(1+stats.EnergyRegen()))
}

func (s *Service) ModifyEnergyFixed(target key.TargetID, amt float64) error {
	t, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}
	return s.SetEnergy(target, t.attributes.Energy+amt)
}
