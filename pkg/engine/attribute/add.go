package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Service) AddTarget(target key.TargetID, attr info.Attributes) error {
	if _, dup := s.targets[target]; dup {
		return fmt.Errorf("target base stats already registered: %v", target)
	}
	if attr.BaseStats == nil {
		attr.BaseStats = info.NewPropMap()
	}
	if attr.BaseDebuffRES == nil {
		attr.BaseDebuffRES = info.NewDebuffRESMap()
	}
	if attr.Weakness == nil {
		attr.Weakness = info.NewWeaknessMap()
	}

	if attr.Energy > attr.MaxEnergy {
		attr.Energy = attr.MaxEnergy
	}

	if attr.HPRatio <= 0 {
		attr.HPRatio = 1.0
	}

	s.targets[target] = &attrTarget{
		attributes:   &attr,
		state:        info.Alive,
		lastAttacker: target,
	}
	return nil
}
