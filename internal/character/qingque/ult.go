package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Ult key.Attack = "qingque-ult"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   5,
	})

	c.tiles = []int{4, 0, 0}
	c.suits[0] = "Yu"
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		Stats:  info.PropMap{prop.ATKPercent: talent[c.info.TalentLevelIndex()]},
	})
	state.EndAttack()
}
