package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique = "xueyi-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Source:       c.id,
		Key:          Technique,
		Targets:      c.engine.Enemies(),
		AttackType:   model.AttackType_MAZE,
		DamageType:   model.DamageType_QUANTUM,
		StanceDamage: 60,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.8,
		},
	})

	state.EndAttack()
}
