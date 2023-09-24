package hook

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Skill key.Attack = "hook-skill"
const EnhancedSkill key.Attack = "hook-enhanced-skill"

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	if c.engine.HasModifier(c.id, SkillEnhancement) {
		c.EnhancedSkill(target, state)
	} else {
		c.NormalSkill(target, state)
	}

	burnDur := 2
	if c.info.Eidolon >= 2 {
		burnDur += 1
	}

	c.engine.AddModifier(target, info.Modifier{
		Name:   common.Burn,
		Source: c.id,
		State: &common.BurnState{
			DamagePercentage: skillBurnDot[c.info.SkillLevelIndex()],
			DamageValue:      0,
		},
		Chance:   1,
		Duration: burnDur,
	})

	c.engine.EndAttack()
}

func (c *char) NormalSkill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Skill,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillNormal[c.info.SkillLevelIndex()],
		},
		EnergyGain:   30,
		StanceDamage: 60,
	})

}

func (c *char) EnhancedSkill(target key.TargetID, state info.ActionState) {
	if c.info.Eidolon >= 1 {
		//c.engine.AddModifier()
	}

	//Main target
	c.engine.Attack(info.Attack{
		Key:        EnhancedSkill,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillEnhanceMain[c.info.SkillLevelIndex()],
		},
		EnergyGain:   30,
		StanceDamage: 60,
	})

	//Adjacent targets
	c.engine.Attack(info.Attack{
		Key:        EnhancedSkill,
		Targets:    c.engine.AdjacentTo(target),
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillEnhanceMain[c.info.SkillLevelIndex()],
		},
		EnergyGain:   30,
		StanceDamage: 60,
	})

	//Remove the enhancement modifier
	c.engine.RemoveModifier(c.id, SkillEnhancement)

}
