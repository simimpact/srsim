// package attribute provides attribute service which is used to keep track of
// character stats
package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

type AttributeGetter interface {
	Stats(target key.TargetID) *info.Stats
	// TODO: getters for internal use (IsAlive, IsStanceBroken, IsMaxEnergy?)
}

type AttributeModifier interface {
	AttributeGetter

	SetHP(target, source key.TargetID, amt float64) error
	ModifyHPByRatio(target, source key.TargetID, data info.ModifyHPByRatio) error
	ModifyHPByAmount(target, source key.TargetID, amt float64) error

	SetEnergy(target key.TargetID, amt float64) error
	ModifyEnergy(target key.TargetID, amt float64) error

	SetStance(target key.TargetID, amt float64) error
	ModifyStance(target key.TargetID, amt float64) error
}

type Service struct {
	event   *event.System
	modEval modifier.ModifierEval
	targets map[key.TargetID]*info.Attributes
}

func New(event *event.System, modEval modifier.ModifierEval) *Service {
	return &Service{
		event:   event,
		modEval: modEval,
		targets: make(map[key.TargetID]*info.Attributes, 10),
	}
}

func (s *Service) Stats(target key.TargetID) *info.Stats {
	mods := s.modEval.EvalModifiers(target)
	attr := s.targets[target]
	if attr == nil {
		attr = &info.Attributes{}
	}
	return info.NewStats(target, attr, mods)
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
