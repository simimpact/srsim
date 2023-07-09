package sampo

import (
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
	// if c.info.Eidolon >= 4 && c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_POISON) {
	// 	//TODO: implement sampo E4
	// }

	targets := c.engine.Enemies()
	c.engine.Attack(info.Attack{
		Key:        Skill,
		HitIndex:   i,
		Source:     c.id,
		Targets:    []key.TargetID{targets[c.engine.Rand().Intn(len(targets))]},
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_SKILL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
		},
		StanceDamage: stanceDamage,
		EnergyGain:   6,
	})
}
