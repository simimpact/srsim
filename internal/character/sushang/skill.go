package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1          key.Reason = "sushang-e1"
	Skill       key.Attack = "sushang-skill"
	SwordStance key.Attack = "sword-stance"
)

var skillHits = []float64{0.3, 0.7}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// check if target is broken
	isBroken := c.engine.Stats(target).Stance() == 0

	// 2 hits
	for i, hitRatio := range skillHits {
		c.engine.Attack(info.Attack{
			Key:        Skill,
			HitIndex:   i,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_PHYSICAL,
			AttackType: model.AttackType_SKILL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skill[c.info.SkillLevelIndex()],
			},
			StanceDamage: 60.0,
			EnergyGain:   30.0,
			HitRatio:     hitRatio,
		})
	}

	// sword stance hits (yes, this uses her own effect RES)
	chances := [3]bool{}
	stats := c.engine.Stats(c.id)
	hitChance := 0.33 * (1 + stats.EffectHitRate()) * (1 - stats.EffectRES())

	if c.engine.HasModifier(c.id, UltBuff) {
		for i := 0; i <= 1; i++ {
			if isBroken || c.engine.Rand().Float64() < hitChance {
				chances[i] = true
			}
		}
	}

	if isBroken || c.engine.Rand().Float64() < hitChance {
		chances[2] = true
	}

	for i, chance := range chances {
		if chance {
			ssHit(c, target, i)
		}
	}

	state.EndAttack()

	if isBroken && c.info.Eidolon >= 1 {
		c.engine.ModifySP(E1, 1)
	}

	c.a6()
}

func ssHit(c *char, target key.TargetID, index int) {
	// handle a4 buff
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4Mod,
			Source: c.id,
			State:  c.engine.ModifierStackCount(c.id, c.id, A4Buff),
		})
	}

	hitRatio := 1.0
	if index != 2 {
		hitRatio = 0.5
	}

	c.engine.Attack(info.Attack{
		Key:        SwordStance,
		HitIndex:   index,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_PHYSICAL,
		AttackType: model.AttackType_PURSUED,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ssDamage[c.info.SkillLevelIndex()],
		},
		StanceDamage: 0.0,
		EnergyGain:   0.0,
		HitRatio:     hitRatio,
	})

	c.e2()

	if c.info.Traces["102"] {
		c.engine.RemoveModifier(c.id, A4Mod)
		c.a4AddStack()
	}
}
