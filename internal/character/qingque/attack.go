package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	atk := c.getAttack()
	atk(target, false)
	if c.tiles[0] == 4 {
		c.engine.RemoveModifier(c.id, Talent)
	}
	if c.engine.HasModifier(c.id, Autarky) {
		c.engine.InsertAbility(info.Insert{
			Execute:  func() { atk(target, true) },
			Source:   c.id,
			Priority: info.CharInsertAttackSelf,
			AbortFlags: []model.BehaviorFlag{
				model.BehaviorFlag_STAT_CTRL,
				model.BehaviorFlag_DISABLE_ACTION,
			},
		})
		c.engine.RemoveModifier(c.id, Autarky)
	}
	if c.tiles[0] == 4 {
		c.tiles = []int{0, 0, 0}
		c.suits[0] = ""
		c.unusedSuits = []string{"Wan", "Tong", "Tiao"}
		c.a6()
		if c.info.Eidolon >= 6 {
			c.engine.ModifySP(1)
		}
	} else {
		c.engine.ModifySP(1)
	}
	state.EndAttack()
}

type attackFunc func(target key.TargetID, isInsert bool)

func (c *char) getAttack() attackFunc {
	if c.tiles[0] == 4 {
		return c.enhancedAttack
	}
	return c.basicAttack
}
func (c *char) basicAttack(target key.TargetID, isInsert bool) {
	aType := model.AttackType_NORMAL
	energy := 20.0
	if isInsert {
		aType = model.AttackType_INSERT
		energy = 0.0
	}
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_QUANTUM,
		AttackType: aType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30.0,
		EnergyGain:   energy,
	})
}

func (c *char) enhancedAttack(target key.TargetID, isInsert bool) {
	aType := model.AttackType_NORMAL
	energy := 20.0
	if isInsert {
		aType = model.AttackType_INSERT
		energy = 0.0
	}
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_QUANTUM,
		AttackType: aType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 2.4 * atk[c.info.AttackLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   energy,
	})
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    c.engine.AdjacentTo(target),
		DamageType: model.DamageType_QUANTUM,
		AttackType: aType,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30.0,
	})
}
