package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// TODO: should have AddCharacter & AddEnemy w/ specific proto defs as input?
// TODO: starting energy
type BaseStats struct {
	Stats     info.PropMap
	DebuffRES info.DebuffRESMap
	MaxEnergy float64
	MaxStance float64
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

	s.targets[target] = &info.Attributes{
		BaseStats: base.Stats,
		MaxStance: base.MaxStance,
		Stance:    base.MaxStance,
		MaxEnergy: base.MaxEnergy,
		Energy:    base.MaxEnergy,
	}
	return nil
}
