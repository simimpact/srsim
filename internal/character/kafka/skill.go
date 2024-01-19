package kafka

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var hitSplitSkill = []float64{0.2, 0.3, 0.5}

var allTriggerableDots = []key.Modifier{common.Burn, common.BreakBurn, common.Bleed,
	common.BreakBleed, common.Shock, common.BreakShock,
	common.WindShear, common.BreakWindShear}

const Skill = "kafka-skill"

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// Main target
	for i, hitRatio := range hitSplitSkill {
		c.engine.Attack(info.Attack{
			Key:        Skill,
			HitIndex:   i,
			Targets:    []key.TargetID{target},
			Source:     c.id,
			AttackType: model.AttackType_NORMAL,
			DamageType: model.DamageType_THUNDER,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skillMain[c.info.SkillLevelIndex()],
			},
			EnergyGain:   30,
			StanceDamage: 60,
			HitRatio:     hitRatio,
		})
	}

	// Adjacent targets
	c.engine.Attack(info.Attack{
		Key:        Skill,
		HitIndex:   1,
		Targets:    c.engine.AdjacentTo(target),
		Source:     c.id,
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_THUNDER,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skillAdj[c.info.SkillLevelIndex()],
		},
		StanceDamage: 30,
		HitRatio:     1,
	})

	// Detonate dots on main target
	if c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT) {
		// Trigger all dots on target with ratio according to kafka skill level
		for _, triggerable := range allTriggerableDots {
			for _, dot := range c.engine.GetModifiers(target, triggerable) {
				dot.State.(common.TriggerableDot).TriggerDot(dot, skillDotDetonate[c.info.SkillLevelIndex()], c.engine, target)
			}
		}
	}

	c.engine.EndAttack()
}
