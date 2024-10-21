package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Normal key.Attack = "ruanmei-normal"

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Normal,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_NORMAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30.0,
		EnergyGain:   20.0,
	})
}
