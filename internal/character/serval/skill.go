package serval

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SkillPrimary  key.Attack = "serval-skill-primary"
	SkillAdjacent key.Attack = "serval-skill-adjacent"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        SkillPrimary,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   30,
	})

	c.engine.Attack(info.Attack{
		Key:        SkillAdjacent,
		Source:     c.id,
		Targets:    c.engine.AdjacentTo(target),
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillBlast[c.info.UltLevelIndex()],
		},
		StanceDamage: 30.0,
		EnergyGain:   0.0,
	})

	for _, trg := range c.engine.Enemies() {
		c.engine.AddModifier(trg, info.Modifier{
			Name: common.Shock,
			State: &common.ShockState{
				DamagePercentage: skillDot[c.info.SkillLevelIndex()],
			},
			Source:   c.id,
			Chance:   0.8,
			Duration: 2,
		})
	}
	state.EndAttack()
}
