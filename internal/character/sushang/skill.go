package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var skillHits = []float64{0.3, 0.7}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// check if target is broken
	isBroken := false
	if c.engine.Stats(target).Stance() == 0 {
		isBroken = true
	}

	// 2 hits
	for _, hitRatio := range skillHits {
		c.engine.Attack(info.Attack{
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

	// sword stance hits
	chances := [3]bool{}

	if c.engine.HasModifier(c.id, UltBuff) {
		for i := 1; i <= 2; i++ {
			if isBroken || c.engine.Rand().Float64() < 0.33 {
				chances[i] = true
			}
		}
	}

	if isBroken || c.engine.Rand().Float64() < 0.33 {
		chances[0] = true
	}

	for i, chance := range chances {
		if chance {
			isExtra := false
			if i != 0 {
				isExtra = true
			}
			ssHit(c, target, isExtra)
		}
	}

	state.EndAttack()

	if isBroken && c.info.Eidolon >= 1 {
		c.engine.ModifySP(1)
	}

	c.a6()
}

func ssHit(c *char, target key.TargetID, isExtra bool) {
	// handle a4 buff
	if c.info.Traces["1206102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4Mod,
			Source: c.id,
		})
	}

	hitRatio := 1.0
	if isExtra {
		hitRatio = 0.5
	}

	c.engine.Attack(info.Attack{
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

	if c.info.Traces["1206102"] {
		c.engine.RemoveModifier(c.id, A4Mod)
		c.a4AddStack()
	}

	c.e2()
}
