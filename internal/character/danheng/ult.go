package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
		})
	}

	// base DMG% + added DMG% if enemy has a speed down modifier
	dmg := ultWindDMG[c.info.UltLevelIndex()]
	if c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_SPEED_DOWN) {
		dmg += ultSlowDMG[c.info.UltLevelIndex()]
	}

	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_WIND,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: dmg,
		},
		StanceDamage: 90.0,
		EnergyGain:   5.0,
	})

	state.EndAttack()
	c.a4()

	if c.info.Eidolon >= 4 {
		c.engine.RemoveModifier(c.id, E4)
	}
}
