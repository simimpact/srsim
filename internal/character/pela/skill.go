package pela

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.AbilityLevel.Skill-1],
		},
		StanceDamage: 60.0,
		EnergyGain:   30.0,
	})

	c.engine.DispelStatus(target, info.Dispel{
		Status: model.StatusType_STATUS_BUFF,
		Order:  model.DispelOrder_LAST_ADDED,
		Count:  1,
	})

	state.EndAttack()
}
