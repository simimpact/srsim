package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type BaseStats struct {
	Stats       info.PropMap
	DebuffRES   info.DebuffRESMap
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

	s.targets[target] = &info.Attributes{
		BaseStats:     base.Stats,
		BaseDebuffRES: base.DebuffRES,
		MaxStance:     base.MaxStance,
		Stance:        base.MaxStance,
		MaxEnergy:     base.MaxEnergy,
		Energy:        base.StartEnergy,
		HPRatio:       1.0,
	}
	return nil
}
