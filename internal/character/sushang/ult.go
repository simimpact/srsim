package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult = "sushang-ult"
)

func init() {
	modifier.Register(Ult, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Source:     c.id,
		Targets:    []key.TargetID{target},
		DamageType: model.DamageType_PHYSICAL,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		StanceDamage: 90.0,
		EnergyGain:   5.0,
		HitRatio:     1.0,
	})

	state.EndAttack()
	c.engine.SetGauge(info.ModifyAttribute{
		Key:    Ult,
		Target: c.id,
		Source: c.id,
		Amount: 0,
	})

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Ult,
		Source: c.id,
		Stats: info.PropMap{
			prop.ATKPercent: ultAtkBuff[c.info.UltLevelIndex()],
		},
		TickImmediately: true,
		Duration:        2,
	})
}
