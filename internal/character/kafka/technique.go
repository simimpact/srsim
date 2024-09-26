package kafka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	technique = "kafka-talent"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        technique,
		Targets:    c.engine.Enemies(),
		Source:     c.id,
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.5,
		},
	})

	c.engine.EndAttack()

	c.applyShock(c.engine.Enemies())
}
