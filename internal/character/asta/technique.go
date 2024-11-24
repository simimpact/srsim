package asta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	astaTechnique = "asta-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        astaTechnique,
		Targets:    c.engine.Enemies(),
		Source:     c.id,
		AttackType: model.AttackType_MAZE,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.5,
		},
	})
}
