package pela

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	oldModCount := c.engine.ModifierStatusCount(target, model.StatusType_STATUS_BUFF)

	c.engine.DispelStatus(target, info.Dispel{
		Status: model.StatusType_STATUS_BUFF,
		Order:  model.DispelOrder_LAST_ADDED,
		Count:  1,
	})

	c.e4(target)

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_ICE,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   30.0,
	})

	state.EndAttack()

	if c.engine.ModifierStatusCount(target, model.StatusType_STATUS_BUFF) < oldModCount {
		c.e2()
		c.a6()
	}
}
