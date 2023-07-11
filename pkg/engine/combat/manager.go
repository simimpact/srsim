package combat

//go:generate go run generate.go

import (
	"math/rand"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/shield"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Manager struct {
	event  *event.System
	attr   attribute.Manager
	shld   shield.Absorb
	target engine.Target
	rdm    *rand.Rand

	isInAttack bool       `exhaustruct:"optional"`
	attackInfo attackInfo `exhaustruct:"optional"`
}

type attackInfo struct {
	attacker   key.TargetID
	targets    []key.TargetID
	attackType model.AttackType
	damageType model.DamageType
}

func New(event *event.System, attr attribute.Manager, shld shield.Absorb, target engine.Target, rdm *rand.Rand) *Manager {
	return &Manager{
		event:  event,
		attr:   attr,
		shld:   shld,
		target: target,
		rdm:    rdm,
	}
}
