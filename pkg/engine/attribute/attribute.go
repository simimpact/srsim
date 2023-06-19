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
	Stance(target key.TargetID) float64
	Energy(target key.TargetID) float64
	HPRatio(target key.TargetID) float64
}

type AttributeModifier interface {
	AttributeGetter

	AddTarget(target key.TargetID, base BaseStats) error

	SetHP(target, source key.TargetID, amt float64, isDamage bool) error
	ModifyHPByRatio(target, source key.TargetID, data info.ModifyHPByRatio, isDamage bool) error
	ModifyHPByAmount(target, source key.TargetID, amt float64, isDamage bool) error

	SetEnergy(target key.TargetID, amt float64) error
	ModifyEnergy(target key.TargetID, amt float64) error
	ModifyEnergyFixed(target key.TargetID, amt float64) error

	SetStance(target, source key.TargetID, amt float64) error
	ModifyStance(target, source key.TargetID, amt float64) error
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

func (s *Service) HPRatio(target key.TargetID) float64 {
	return s.targets[target].HPRatio
}

func (s *Service) Energy(target key.TargetID) float64 {
	return s.targets[target].Energy
}

func (s *Service) FullEnergy(target key.TargetID) bool {
	attr := s.targets[target]
	return attr.Energy >= attr.MaxEnergy
}

func (s *Service) Stance(target key.TargetID) float64 {
	return s.targets[target].Stance
}
