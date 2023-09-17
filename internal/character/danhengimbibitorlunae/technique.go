package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Technique key.Attack = "danhengimbibitorlunae-technique"

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Point,
		Source: c.id,
	})
	c.engine.Attack(info.Attack{
		Key:        Technique,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_IMAGINARY,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 1.2,
		},
		StanceDamage: 30.0,
		EnergyGain:   0.0,
	})
}
