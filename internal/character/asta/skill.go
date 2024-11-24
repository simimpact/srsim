package asta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill = "asta-skill"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Skill,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		Targets:      []key.TargetID{target},
		Source:       c.id,
		EnergyGain:   6,
		StanceDamage: 30,
	})

	bounceCount := 4
	if c.info.Eidolon >= 1 {
		bounceCount++
	}

	for i := 0; i < bounceCount; i++ {
		target := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Enemies(),
			Filter: func(target key.TargetID) bool {
				return c.engine.HPRatio(target) > 0
			},
			Max: 1,
		})[0]
		c.engine.Attack(info.Attack{
			Key:        Skill,
			AttackType: model.AttackType_SKILL,
			DamageType: model.DamageType_FIRE,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
			},
			Targets:      []key.TargetID{target},
			Source:       c.id,
			EnergyGain:   6,
			StanceDamage: 15,
		})
	}

	c.engine.EndAttack()
}
