package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type BaseStats struct {
	Level       int
	Stats       info.PropMap
	DebuffRES   info.DebuffRESMap
	Weakness    []model.DamageType
	StartEnergy float64
	MaxEnergy   float64
	MaxStance   float64
}

func (s *Service) AddTarget(target key.TargetID, base BaseStats) error {
	if _, dup := s.targets[target]; dup {
		return fmt.Errorf("target base stats already registered: %v", target)
	}
	if base.Stats == nil {
		base.Stats = info.NewPropMap()
	}
	if base.DebuffRES == nil {
		base.DebuffRES = info.NewDebuffRESMap()
	}
	if base.StartEnergy > base.MaxEnergy {
		base.StartEnergy = base.MaxEnergy
	}

	wMap := info.NewWeaknessMap()
	for _, w := range base.Weakness {
		wMap[w] = true
	}

	s.targets[target] = &info.Attributes{
		Level:         base.Level,
		BaseStats:     base.Stats,
		BaseDebuffRES: base.DebuffRES,
		Weakness:      wMap,
		MaxStance:     base.MaxStance,
		Stance:        base.MaxStance,
		MaxEnergy:     base.MaxEnergy,
		Energy:        base.StartEnergy,
		HPRatio:       1.0,
	}
	return nil
}
