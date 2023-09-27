package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Attack1         key.Attack = "danhengimbibitorlunae-enhancedattack-1"
	Attack2Primary  key.Attack = "danhengimbibitorlunae-enhancedattack-2-primary"
	Attack2Adjacent key.Attack = "danhengimbibitorlunae-enhancedattack-2-adjacent"
	Attack3Primary  key.Attack = "danhengimbibitorlunae-enhancedattack-3-primary"
	Attack3Adjacent key.Attack = "danhengimbibitorlunae-enhancedattack-3-adjacent"
	Attack          key.Attack = "danhengimbibitorlunae-normalattack"
	AttackReason    key.Reason = "danhengimbibitorlunae-attack"
)

var attackHitsNormal = []float64{0.3, 0.7}
var attackHitsEnhanced1 = []float64{0.33, 0.33, 0.34}
var attackHitsEnhanced2 = []float64{0.2, 0.2, 0.2, 0.2, 0.2}
var adjacentHitsEnhanced2 = []float64{0, 0, 0, 0.5, 0.5}
var attackHitsEnhanced3 = []float64{0.142, 0.142, 0.142, 0.142, 0.142, 0.142, 0.148}
var adjacentHitsEnhanced3 = []float64{0, 0, 0, 0.25, 0.25, 0.25, 0.25}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	if c.attackLevel == 0 {
		c.NormalAttack(target, state)
		c.engine.ModifySP(info.ModifySP{
			Key:    AttackReason,
			Source: c.id,
			Amount: 1,
		})
	}

	pointUse := c.attackLevel
	pointHas := c.point
	c.point = pointHas - pointUse
	if c.point < 0 {
		c.point = 0
	}
	pointUse -= pointHas
	if pointUse > 0 {
		c.engine.ModifySP(info.ModifySP{
			Key:    AttackReason,
			Source: c.id,
			Amount: -pointUse,
		})
	}
	switch c.attackLevel {
	case 1:
		c.EnhancedAttack1(target, state)
	case 2:
		c.EnhancedAttack2(target, state)
	case 3:
		// add e6 buff
		if c.info.Eidolon >= 6 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E6Effect,
				Source: c.id,
				Stats:  info.PropMap{prop.ImaginaryPEN: float64(c.E6Count) * 0.2},
			})
		}
		c.EnhancedAttack3(target, state)
		// reset count,remove e6
		c.engine.RemoveModifier(c.id, E6Effect)
		c.E6Count = 0
	}
	c.attackLevel = 0
	state.EndAttack()
}
func (c *char) NormalAttack(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHitsNormal {
		c.engine.Attack(info.Attack{
			Key:          Attack,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()]},
			StanceDamage: 30,
			EnergyGain:   20,
			HitRatio:     hitRatio,
		})
		c.AddTalent()
	}
}
func (c *char) EnhancedAttack1(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHitsEnhanced1 {
		c.engine.Attack(info.Attack{
			Key:          Attack1,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk1[c.info.AttackLevelIndex()]},
			StanceDamage: 60,
			EnergyGain:   30,
			HitRatio:     hitRatio,
		})
		c.AddTalent()
	}
}
func (c *char) EnhancedAttack2(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHitsEnhanced2 {
		if i >= 3 {
			c.AddSkill()
		}
		c.engine.Attack(info.Attack{
			Key:          Attack2Primary,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk2[c.info.AttackLevelIndex()]},
			StanceDamage: 90,
			EnergyGain:   35,
			HitRatio:     hitRatio,
		})
		if adjacentHitsEnhanced2[i] > 0 {
			c.engine.Attack(info.Attack{
				Key:          Attack2Adjacent,
				HitIndex:     i,
				Source:       c.id,
				Targets:      c.engine.AdjacentTo(target),
				DamageType:   model.DamageType_IMAGINARY,
				AttackType:   model.AttackType_NORMAL,
				BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk2[c.info.AttackLevelIndex()] * 3 / 19},
				StanceDamage: 30,
				HitRatio:     adjacentHitsEnhanced2[i],
			})
		}
		c.AddTalent()
	}
}
func (c *char) EnhancedAttack3(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHitsEnhanced3 {
		if i >= 3 {
			c.AddSkill()
		}
		c.engine.Attack(info.Attack{
			Key:          Attack3Primary,
			HitIndex:     i,
			Source:       c.id,
			Targets:      []key.TargetID{target},
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk3[c.info.AttackLevelIndex()]},
			StanceDamage: 120,
			EnergyGain:   40,
			HitRatio:     hitRatio,
		})
		if adjacentHitsEnhanced3[i] > 0 {
			c.engine.Attack(info.Attack{
				Key:          Attack3Adjacent,
				HitIndex:     i,
				Source:       c.id,
				Targets:      c.engine.AdjacentTo(target),
				DamageType:   model.DamageType_IMAGINARY,
				AttackType:   model.AttackType_NORMAL,
				BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk3[c.info.AttackLevelIndex()] * 9 / 25},
				StanceDamage: 60,
				HitRatio:     adjacentHitsEnhanced3[i],
			})
		}
		c.AddTalent()
	}
}
