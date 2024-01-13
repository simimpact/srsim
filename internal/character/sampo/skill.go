package sampo

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Skill key.Attack = "sampo-skill"

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.OnProjectileHit(target, 30, 0)

	bounces := 4
	if c.info.Eidolon >= 1 {
		bounces += 1
	}

	for i := 0; i < bounces; i++ {
		c.OnProjectileHit(target, 15, i+1)
	}

	state.EndAttack()
}

func (c *char) OnProjectileHit(target key.TargetID, stanceDamage float64, i int) {
	if c.info.Eidolon >= 4 && c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_POISON) {
		// Wind dot (kit dots)
		for _, dot := range c.engine.GetModifiers(target, common.WindShear) {
			dot.State.(common.WindShearState).TriggerDot(dot, 0.06, c.engine, target)
		}

		// Break wind dots
		for _, dot := range c.engine.GetModifiers(target, common.BreakWindShear) {
			dot.State.(common.BreakWindShearState).TriggerDot(dot, 0.06, c.engine, target)
		}
	}

	targets := []key.TargetID{target}
	if i > 0 {
		enemies := c.engine.Enemies()
		allDead := allTargetsDead(c.engine, enemies)

		targets = c.engine.Retarget(info.Retarget{
			Targets:      enemies,
			Max:          1,
			IncludeLimbo: true,
			Filter: func(t key.TargetID) bool {
				return allDead || c.engine.HPRatio(t) > 0
			},
		})
	}

	c.engine.Attack(info.Attack{
		Key:        Skill,
		HitIndex:   i,
		Source:     c.id,
		Targets:    targets,
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: stanceDamage,
		EnergyGain:   6,
	})
}

func allTargetsDead(engine engine.Engine, targets []key.TargetID) bool {
	for _, t := range targets {
		if engine.HPRatio(t) > 0 {
			return false
		}
	}
	return true
}
