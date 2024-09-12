package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const technique = "luka-technique"

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	victim := c.engine.Retarget(info.Retarget{
		Targets:      c.engine.Enemies(),
		IncludeLimbo: false,
		Max:          1,
	})
	c.engine.Attack(info.Attack{
		Key:        technique,
		Targets:    victim,
		Source:     c.id,
		AttackType: model.AttackType_MAZE,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.5,
		},
	})

	// inflict same bleed as skill
	c.applyBleed(victim[0])

	c.incrementFightingSprit()
	state.EndAttack()
}
