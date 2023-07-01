package shield

//go:generate mockgen -destination=../../mock/mock_shield.go -package=mock -mock_names ShieldAbsorb=MockShield github.com/simimpact/srsim/pkg/engine/shield ShieldAbsorb

import (
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

type activeShields map[key.Shield]*ShieldInstance

type ShieldAbsorb interface {
	AbsorbDamage(target key.TargetID, damage float64) float64
}

type Manager struct {
	event *event.System
	attr  attribute.AttributeGetter

	targets map[key.TargetID]activeShields
}

func New(event *event.System, attr attribute.AttributeGetter) *Manager {
	return &Manager{
		event:   event,
		attr:    attr,
		targets: make(map[key.TargetID]activeShields, 10),
	}
}

func (mgr *Manager) HasShield(target key.TargetID, shield key.Shield) bool {
	if shields, ok := mgr.targets[target]; ok {
		_, ok := shields[shield]
		return ok
	}
	return false
}

func (mgr *Manager) IsShielded(target key.TargetID) bool {
	if shields, ok := mgr.targets[target]; ok {
		return len(shields) > 0
	}
	return false
}
