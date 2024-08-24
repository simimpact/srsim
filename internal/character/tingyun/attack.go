package tingyun

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "tingyun-normal"

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	// A4 implemented the same way as game with add and remove
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   A4,
		Source: c.id,
	})

	c.engine.Attack(info.Attack{
		Key:        Normal,
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		Targets:      []key.TargetID{target},
		Source:       c.id,
		EnergyGain:   20,
		StanceDamage: 30,
	})

	c.engine.RemoveModifier(c.id, A4)

	c.engine.EndAttack()
}
