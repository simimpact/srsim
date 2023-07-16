package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Skill = "arlan-skill"

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.ModifyHPByRatio(info.ModifyHPByRatio{
		Key:       Skill,
		Target:    c.id,
		Source:    c.id,
		Ratio:     -0.15,
		RatioType: model.ModifyHPRatioType_MAX_HP,
		Floor:     1,
	})

	c.e2()

	c.engine.Attack(info.Attack{
		Key:        Skill,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   30.0,
	})

	state.EndAttack()
}
