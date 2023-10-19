package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	NormalSkill           key.Attack = "jingliu-skill-normal"
	EnhancedSkillPrimary  key.Attack = "jingliu-skill-ehnance-primary"
	EnhancedSkillAdjacent key.Attack = "jingliu-skill-ehnance-adjacent"
	A4                    key.Reason = "jingliu-a4"
)

// 3*3616+7232+18080
// E1 3*(19106+6947)+(38213+13895)+(95332+34739)
// main 191060 E1 69740 2.75:1
var EnhancedSkillRatio = []float64{0.1, 0.1, 0.1, 0.2, 0.5}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	if !c.isEnhanced {
		c.NormalSkill(target, state)
		return
	}
	if len(c.engine.AdjacentTo(target)) == 0 && c.info.Eidolon >= 1 {
		c.EnhancedSkillE1(target, state)
		return
	}
	c.EnhancedSkill(target, state)
}

func (c *char) NormalSkill(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:          NormalSkill,
		Source:       c.id,
		Targets:      []key.TargetID{target},
		DamageType:   model.DamageType_ICE,
		AttackType:   model.AttackType_SKILL,
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()]},
		StanceDamage: 60,
		EnergyGain:   20,
	})
	state.EndAttack()
	if c.info.Traces["102"] {
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    A4,
			Target: c.id,
			Source: c.id,
			Amount: -0.1,
		})
	}
	c.gainSyzygy()
}

func (c *char) EnhancedSkill(target key.TargetID, state info.ActionState) {
	c.addTalentBuff()
	c.Syzygy -= 1

	for i, hitRatio := range attackRatio {
		c.engine.Attack(info.Attack{
			Key:          EnhancedSkillPrimary,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_SKILL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedSkill[c.info.SkillLevelIndex()]},
			StanceDamage: 60,
			EnergyGain:   30,
			HitRatio:     hitRatio,
		})
		c.engine.Attack(info.Attack{
			Key:          EnhancedSkillPrimary,
			HitIndex:     i,
			Source:       c.id,
			Targets:      c.engine.AdjacentTo(target),
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_SKILL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedSkill[c.info.SkillLevelIndex()] / 2.0},
			StanceDamage: 30,
			HitRatio:     hitRatio,
		})
	}
	state.EndAttack()
}

func (c *char) EnhancedSkillE1(target key.TargetID, state info.ActionState) {
	c.addTalentBuff()
	c.Syzygy -= 1

	for i, hitRatio := range attackRatio {
		c.engine.Attack(info.Attack{
			Key:          EnhancedSkillPrimary,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_ICE,
			AttackType:   model.AttackType_SKILL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedSkill[c.info.SkillLevelIndex()]},
			StanceDamage: 60,
			EnergyGain:   30,
			HitRatio:     hitRatio,
		})
		c.engine.Attack(info.Attack{
			Key:        EnhancedSkillPrimary,
			HitIndex:   i,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_ICE,
			AttackType: model.AttackType_SKILL,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 1},
			HitRatio:   hitRatio,
		})
	}
	state.EndAttack()
}
