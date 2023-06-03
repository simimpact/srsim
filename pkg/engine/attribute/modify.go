package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (s *Service) SetHP(target, source key.TargetID, amt float64) error {
	attr, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}

	oldRatio := attr.HPRatio
	stats := s.Stats(target)
	attr.HPRatio = amt / stats.MaxHP()

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	} else if attr.HPRatio < 0 {
		attr.HPRatio = 0
	}

	return s.emitHPChangeEvents(target, source, oldRatio, attr.HPRatio, stats.MaxHP())
}

func (s *Service) ModifyHPByAmount(target, source key.TargetID, amt float64) error {
	attr, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}

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

	return s.emitHPChangeEvents(target, source, oldRatio, attr.HPRatio, stats.MaxHP())
}

func (s *Service) ModifyHPByRatio(target, source key.TargetID, data info.ModifyHPByRatio) error {
	attr, ok := s.targets[target]
	if !ok {
		return fmt.Errorf("unknown target: %v", target)
	}

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
		return s.SetHP(target, source, data.Floor)
	}

	// TODO: unsure if there are limits on min and max
	if attr.HPRatio > 1 {
		attr.HPRatio = 1.0
	}

	return s.emitHPChangeEvents(target, source, oldRatio, attr.HPRatio, stats.MaxHP())
}

func (s *Service) SetStance(target key.TargetID, amt float64) error {
	return nil
}

func (s *Service) ModifyStance(target key.TargetID, amt float64) error {
	return nil
}

func (s *Service) SetEnergy(target key.TargetID, amt float64) error {
	return nil
}

func (s *Service) ModifyEnergy(target key.TargetID, amt float64) error {
	return nil
}
