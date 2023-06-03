package combat

import (
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Manager struct {
	event *event.System
	attr  attribute.AttributeModifier
	// TODO: ShieldAbsorbDamage(target, amt) float64

	isInAttack   bool
	attacker     key.TargetID
	targets      []key.TargetID
	attackType   model.AttackType
	attackEffect model.AttackEffect
	damageType   model.DamageType
}

func New(event *event.System, attr attribute.AttributeModifier) *Manager {
	return &Manager{
		event:   event,
		attr:    attr,
		targets: make([]key.TargetID, 5),
	}
}
