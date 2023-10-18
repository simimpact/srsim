package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "march7th-normal"

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Normal,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_ICE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
		},
		EnergyGain:   20,
		StanceDamage: 30,
	})
}
