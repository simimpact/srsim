package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltPrimary  key.Attack = "jingliu-ult-primary"
	UltAdjacent key.Attack = "jingliu-ult-adjacent"
	UltE1       key.Attack = "jingliu-ult-e1"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.NormalUlt(target, state)
	c.afterUlt = true
}

func (c *char) NormalUlt(target key.TargetID, state info.ActionState) {
	if c.isEnhanced {
		c.addTalentBuff()
	}
	c.engine.Attack(info.Attack{
		Key:          UltPrimary,
		Source:       c.id,
		Targets:      []key.TargetID{target},
		DamageType:   model.DamageType_ICE,
		AttackType:   model.AttackType_ULT,
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()]},
		StanceDamage: 60,
		EnergyGain:   5,
	})
	c.engine.Attack(info.Attack{
		Key:          UltAdjacent,
		Source:       c.id,
		Targets:      c.engine.AdjacentTo(target),
		DamageType:   model.DamageType_ICE,
		AttackType:   model.AttackType_ULT,
		BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()] / 2.0},
		StanceDamage: 60,
	})
	if len(c.engine.AdjacentTo(target)) == 0 && c.info.Eidolon >= 1 {
		c.engine.Attack(info.Attack{
			Key:        UltE1,
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_ICE,
			AttackType: model.AttackType_ULT,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 1},
		})
	}

	state.EndAttack()
	c.gainSyzygy()
}
