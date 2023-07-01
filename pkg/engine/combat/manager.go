package combat

//go:generate go run generate.go

import (
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/shield"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Manager struct {
	event *event.System
	attr  attribute.AttributeModifier
	shld  shield.ShieldAbsorb

	isInAttack bool
	attackInfo attackInfo
}

type attackInfo struct {
	attacker   key.TargetID
	targets    []key.TargetID
	attackType model.AttackType
	damageType model.DamageType
}

func New(event *event.System, attr attribute.AttributeModifier, shld shield.ShieldAbsorb) *Manager {
	return &Manager{
		event: event,
		attr:  attr,
		shld:  shld,
	}
}
