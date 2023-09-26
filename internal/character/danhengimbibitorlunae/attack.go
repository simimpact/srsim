package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Attack1      key.Attack = "danhengimbibitorlunae-enhancedattack-1"
	Attack2      key.Attack = "danhengimbibitorlunae-enhancedattack-2"
	Attack3      key.Attack = "danhengimbibitorlunae-enhancedattack-3"
	Attack       key.Attack = "danhengimbibitorlunae-normalattack"
	AttackReason key.Reason = "danhengimbibitorlunae-attack"
)

var attackHitsNormal = []float64{0.3, 0.7}
var attackHitsEnhanced1 = []float64{0.33, 0.33, 0.34}
var attackHitsEnhanced2 = []float64{0.2, 0.2, 0.2, 0.2, 0.2}
var adjacentHitsEnhanced2 = []float64{0, 0, 0, 0.5, 0.5}
var attackHitsEnhanced3 = []float64{0.142, 0.142, 0.142, 0.142, 0.142, 0.142, 0.148}
var adjacentHitsEnhanced3 = []float64{0, 0, 0, 0.25, 0.25, 0.25, 0.25}

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	if !c.engine.HasModifier(c.id, EnhanceLevel) {
		c.NormalAttack(target, state)
		c.engine.ModifySP(info.ModifySP{
			Key:    AttackReason,
			Source: c.id,
			Amount: 1,
		})
	}

	level := int(c.engine.ModifierStackCount(c.id, c.id, EnhanceLevel))
	c.engine.RemoveModifier(c.id, EnhanceLevel)

	pointUse := level
	pointHas := 0
	if c.engine.HasModifier(c.id, Point) {
		pointHas = int(c.engine.ModifierStackCount(c.id, c.id, Point))
	}
	c.engine.RemoveModifier(c.id, Point)
	if pointUse > pointHas {
		pointUse -= pointHas
	} else {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   Point,
			Source: c.id,
			Count:  float64(pointHas - pointUse),
		})
		pointUse = 0
	}
	c.engine.ModifySP(info.ModifySP{
		Key:    AttackReason,
		Source: c.id,
		Amount: -pointUse,
	})

	if level == 1 {
		c.EnhancedAttack1(target, state)
	}
	if level == 2 {
		c.EnhancedAttack2(target, state)
	}
	if level == 3 {
		// add e6 buff
		if c.engine.HasModifier(c.id, E6Count) {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E6Effect,
				Source: c.id,
				Stats:  info.PropMap{prop.ImaginaryPEN: 20 * c.engine.ModifierStackCount(c.id, c.id, E6Count)},
			})
		}
		c.EnhancedAttack3(target, state)
		// reset count,remove e6
		c.engine.RemoveModifier(c.id, E6Effect)
		c.engine.RemoveModifier(c.id, E6Count)
	}
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
			Key:          Attack2,
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
		c.engine.Attack(info.Attack{
			Key:          Attack2,
			HitIndex:     i,
			Source:       c.id,
			Targets:      c.engine.AdjacentTo(target),
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk2[c.info.AttackLevelIndex()] * 3 / 19},
			StanceDamage: 30,
			HitRatio:     adjacentHitsEnhanced2[i],
		})
		c.AddTalent()
	}
}
func (c *char) EnhancedAttack3(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range attackHitsEnhanced3 {
		if i >= 3 {
			c.AddSkill()
		}
		c.engine.Attack(info.Attack{
			Key:          Attack3,
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
		c.engine.Attack(info.Attack{
			Key:          Attack3,
			HitIndex:     i,
			Source:       c.id,
			Targets:      c.engine.AdjacentTo(target),
			DamageType:   model.DamageType_IMAGINARY,
			AttackType:   model.AttackType_NORMAL,
			BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: enhancedAtk3[c.info.AttackLevelIndex()] * 9 / 25},
			StanceDamage: 60,
			HitRatio:     adjacentHitsEnhanced3[i],
		})
		c.AddTalent()
	}
}
