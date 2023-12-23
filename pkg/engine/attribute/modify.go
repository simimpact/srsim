package attribute

import (
	"fmt"
	
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

func (s *Service) SetHP(data info.ModifyAttribute, isDamage bool) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	attr := t.attributes

	oldRatio := attr.HPRatio
	stats := s.Stats(data.Target)
	attr.HPRatio = data.Amount / stats.MaxHP()

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	} else if attr.HPRatio < 0 {
		attr.HPRatio = 0
	}

	return s.emitHPChangeEvents(
		data.Key, data.Target, data.Source, oldRatio, attr.HPRatio, stats.MaxHP(), isDamage)
}

func (s *Service) ModifyHPByAmount(data info.ModifyAttribute, isDamage bool) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	attr := t.attributes

	oldRatio := attr.HPRatio
	stats := s.Stats(data.Target)

	newHP := stats.CurrentHP() + data.Amount
	attr.HPRatio = newHP / stats.MaxHP()

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	} else if attr.HPRatio < 0 {
		attr.HPRatio = 0
	}

	return s.emitHPChangeEvents(
		data.Key, data.Target, data.Source, oldRatio, attr.HPRatio, stats.MaxHP(), isDamage)
}

func (s *Service) ModifyHPByRatio(data info.ModifyHPByRatio, isDamage bool) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
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

	stats := s.Stats(data.Target)
	if stats.CurrentHP() < data.Floor {
		return s.SetHP(info.ModifyAttribute{
			Key:    data.Key,
			Target: data.Target,
			Source: data.Source,
			Amount: data.Floor,
		}, isDamage)
	}

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	}

	return s.emitHPChangeEvents(
		data.Key, data.Target, data.Source, oldRatio, attr.HPRatio, stats.MaxHP(), isDamage)
}

func (s *Service) SetStance(data info.ModifyAttribute) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	attr := t.attributes

	if data.Amount == 0 {
		s.event.StanceBreak.Emit(event.StanceBreak{
			Key:    data.Key,
			Target: data.Target,
			Source: data.Source,
		})
	} else if attr.Stance == 0 {
		s.event.StanceReset.Emit(event.StanceReset{
			Key:    data.Key,
			Target: data.Target,
		})
	}

	prev := attr.Stance
	attr.Stance = data.Amount
	if attr.Stance > attr.MaxStance {
		attr.Stance = attr.MaxStance
	} else if attr.Stance < 0 {
		attr.Stance = 0
	}

	return s.emitStanceChange(data.Key, data.Target, data.Source, prev, attr.Stance)
}

func (s *Service) ModifyStance(data info.ModifyAttribute) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	attr := t.attributes

	stats := s.Stats(data.Target)
	newStance := attr.Stance + data.Amount*(1+stats.GetProperty(prop.AllStanceDMGPercent))
	return s.SetStance(info.ModifyAttribute{
		Key:    data.Key,
		Target: data.Target,
		Source: data.Source,
		Amount: newStance,
	})
}

func (s *Service) SetEnergy(data info.ModifyAttribute) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	attr := t.attributes

	prev := attr.Energy
	attr.Energy = data.Amount
	if attr.Energy > attr.MaxEnergy {
		attr.Energy = attr.MaxEnergy
	} else if attr.Energy < 0 {
		attr.Energy = 0
	}

	return s.emitEnergyChange(data.Key, data.Target, data.Source, prev, attr.Energy)
}

func (s *Service) ModifyEnergy(data info.ModifyAttribute) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	attr := t.attributes

	stats := s.Stats(data.Target)
	return s.SetEnergy(info.ModifyAttribute{
		Key:    data.Key,
		Target: data.Target,
		Source: data.Source,
		Amount: attr.Energy + data.Amount*(1+stats.EnergyRegen()),
	})
}

func (s *Service) ModifyEnergyFixed(data info.ModifyAttribute) error {
	t, ok := s.targets[data.Target]
	if !ok {
		return fmt.Errorf("unknown target: %v", data.Target)
	}
	return s.SetEnergy(info.ModifyAttribute{
		Key:    data.Key,
		Target: data.Target,
		Source: data.Source,
		Amount: t.attributes.Energy + data.Amount,
	})
}

func (s *Service) ModifySP(data info.ModifySP) error {
	old := s.sp
	s.sp += data.Amount
	if s.sp > 5 {
		s.sp = 5
	} else if s.sp < 0 {
		s.sp = 0
	}
	return s.emitSPChange(data.Key, data.Source, old, s.sp)
}
