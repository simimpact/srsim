package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Attack = "trailblazerimaginary-skill"
	E1    key.Reason = "trailblazerimaginary-e1"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	stanceDamage := 30.0
	if c.info.Traces["102"] {
		stanceDamage = 60.0
	}
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_IMAGINARY,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillAttackRate[c.info.SkillLevelIndex()],
		},
		StanceDamage: stanceDamage,
		EnergyGain:   6,
	})
	extra := 4
	if c.info.Eidolon >= 6 {
		extra = 6
	}
	for i := 0; i < extra; i++ {
		c.Hit()
	}
	if c.info.Eidolon >= 1 && !c.E1Used {
		c.E1Used = true
		c.engine.ModifySP(info.ModifySP{
			Key:    E1,
			Source: c.id,
			Amount: 1,
		})
	}
	state.EndAttack()
}

func (c *char) Hit() {
	allTargetsDead := true
	for _, t := range c.engine.Enemies() {
		if c.engine.HPRatio(t) > 0 {
			allTargetsDead = false
			break
		}
	}
	targets := c.engine.Retarget(info.Retarget{
		Targets:      c.engine.Enemies(),
		Max:          1,
		IncludeLimbo: true,
		Filter: func(target key.TargetID) bool {
			return allTargetsDead || c.engine.HPRatio(target) > 0
		},
	})
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Source:     c.id,
		Targets:    targets,
		DamageType: model.DamageType_IMAGINARY,
		AttackType: model.AttackType_NORMAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillAttackRate[c.info.SkillLevelIndex()],
		},
		StanceDamage: 15.0,
		EnergyGain:   6,
	})
}
