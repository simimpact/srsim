package shield

import (
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

type activeShields []*Instance

type Absorb interface {
	AbsorbDamage(target key.TargetID, damage float64) float64
}

type Manager struct {
	event *event.System
	attr  attribute.Getter

	targets map[key.TargetID]activeShields
}

func New(event *event.System, attr attribute.Getter) *Manager {
	return &Manager{
		event:   event,
		attr:    attr,
		targets: make(map[key.TargetID]activeShields, 10),
	}
}

func (mgr *Manager) HasShield(target key.TargetID, shieldKey key.Shield) bool {
	for _, shield := range mgr.targets[target] {
		if shield.name == shieldKey {
			return true
		}
	}
	return false
}

func (mgr *Manager) IsShielded(target key.TargetID) bool {
	if shields, ok := mgr.targets[target]; ok {
		return len(shields) > 0
	}
	return false
}

func (mgr *Manager) MaxShield(target key.TargetID) float64 {
	maxShield := 0.0
	for _, shield := range mgr.targets[target] {
		if shield.hp > maxShield {
			maxShield = shield.hp
		}
	}
	return maxShield
}
