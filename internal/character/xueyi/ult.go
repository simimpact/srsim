package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	UltBuff  = "xueyi-ult-dmg-buff"
	Ultimate = "xueyi-ult"
)

func init() {
	modifier.Register(UltBuff, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeHit: ultBuffCallback,
		},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   UltBuff,
		Source: c.id,
	})

	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     E4,
			Source:   c.id,
			Duration: 2,
			Stats: info.PropMap{
				prop.BreakEffect: 0.4,
			},
		})
	}

	if c.info.Traces["102"] && c.engine.Stats(target).Stance() >= c.engine.MaxStance(target)*0.5 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4,
			Source: c.id,
		})
	}

	c.engine.Attack(info.Attack{
		Key:        Ultimate,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_ULT,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		EnergyGain: 5,
	})

	c.engine.ModifyStance(info.ModifyAttribute{
		Key:    Ultimate,
		Target: target,
		Source: c.id,
		Amount: -120,
	})

	state.EndAttack()

	c.engine.RemoveModifier(c.id, UltBuff)
	c.engine.RemoveModifier(c.id, A4)
}

func ultBuffCallback(mod *modifier.Instance, e event.HitStart) {
	multiplier := 1.0
	if e.Hit.StanceDamage > 30 {
		multiplier = e.Hit.StanceDamage / 30
	}
	stats, _ := mod.Engine().CharacterInfo(mod.Source())

	buff := ultbuff[stats.UltLevelIndex()] * multiplier
	if buff > ultbuff[stats.UltLevelIndex()]*4 {
		buff = ultbuff[stats.UltLevelIndex()] * 4
	}

	e.Hit.Attacker.AddProperty(
		UltBuff,
		prop.AllDamagePercent,
		buff,
	)
}
