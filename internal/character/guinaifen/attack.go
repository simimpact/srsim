package guinaifen

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "guinaifen-normal"

func (c *char) Attack(target key.TargetID, state info.ActionState) {
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

	c.engine.EndAttack()
}
