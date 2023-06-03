// package attribute provides attribute service which is used to keep track of
// character stats
package attribute

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

type AttributeGetter interface {
	Stats(target key.TargetID) *info.Stats
}

type AttributeModifier interface {
	AttributeGetter
	// TODO: ModifyHP, ModifyStance, ModifyEnergy
}

type Service struct {
	engine          engine.Engine
	modifierManager *modifier.Manager
	targets         map[key.TargetID]*info.Attributes
}

func New(engine engine.Engine, modManager *modifier.Manager) *Service {
	return &Service{
		engine:          engine,
		modifierManager: modManager,
		targets:         make(map[key.TargetID]*info.Attributes),
	}
}

func (s *Service) Stats(target key.TargetID) *info.Stats {
	mods := s.modifierManager.EvalModifiers(target)
	attr := s.targets[target]
	if attr == nil {
		attr = &info.Attributes{}
	}
	return info.NewStats(target, *attr, mods)
}

// TODO:
// 	- AddTarget specific functions
//	- BaseStats struct

// Metadata to have for stats (easy access):
//	- level
//	- weaknesses

// TODO: ChangeHP, return new HP (emit HPChangeEvent)
// TODO: ChangeStance, return new Stance (emit StanceChangeEvent)
// TODO: ChangeEnergy, return new Energy (emit EnergyChangeEvent)
