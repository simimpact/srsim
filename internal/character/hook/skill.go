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
		c.EnhancedSkill(target)
	} else {
		c.NormalSkill(target)
	}

	c.applySkillBurn([]key.TargetID{target})
}

func (c *char) NormalSkill(target key.TargetID) {
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

	if c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BURN) {
		c.talentProc(target)
	} else {
		c.engine.EndAttack()
	}
}

// Special checks to mimic dm/also to avoid multiple procs of the energy gain
func (c *char) EnhancedSkill(target key.TargetID) {
	// Main target
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

	// Adjacent targets
	c.engine.Attack(info.Attack{
		Key:        EnhancedSkill,
		Targets:    c.engine.AdjacentTo(target),
		Source:     c.id,
		AttackType: model.AttackType_SKILL,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillEnhanceAdj[c.info.SkillLevelIndex()],
		},
		EnergyGain:   0,
		StanceDamage: 30,
	})

	talentCandidates := c.engine.Retarget(info.Retarget{
		Targets: append(c.engine.AdjacentTo(target), target),
		Filter: func(target key.TargetID) bool {
			return c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BURN)
		},
		IncludeLimbo: true,
	})

	// Cannot simply loop over every enemy from the retarget with talentProc
	// since that would trigger hp restore and energy restore multiple times
	if len(talentCandidates) > 0 {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    Talent,
			Target: c.id,
			Source: c.id,
			Amount: 5,
		})

		for _, t := range talentCandidates {
			c.talentPursuedDamage(t)
		}

		if c.info.Eidolon >= 4 {
			c.applySkillBurn(c.engine.AdjacentTo(target))
		}

		c.talentHeal()
	} else {
		c.engine.EndAttack()
	}

	// Remove the enhancement modifier
	c.engine.RemoveModifier(c.id, SkillEnhancement)

	// Remove the e1 modifier if it exists
	c.engine.RemoveModifier(c.id, E1)
}

// Apply hook's skill burn to all given targets
func (c *char) applySkillBurn(targets []key.TargetID) {
	burnDur := 2
	if c.info.Eidolon >= 2 {
		burnDur += 1
	}

	for _, target := range targets {
		c.engine.AddModifier(target, info.Modifier{
			Name:   common.Burn,
			Source: c.id,
			State: &common.BurnState{
				DamagePercentage:    skillBurnDot[c.info.SkillLevelIndex()],
				DamageValue:         0,
				DEFDamagePercentage: 0,
			},
			Chance:   1,
			Duration: burnDur,
		})
	}
}
