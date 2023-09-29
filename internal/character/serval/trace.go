package serval

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A2:
//
//	Skill has a 20% increased base chance to Shock enemies.
//
// A4:
//
//	At the start of the battle, immediately regenerates 15 Energy.
//
// A6:
//
//	Upon defeating an enemy, ATK is increased by 20% for 2 turn(s).
const (
	A2 key.Modifier = "serval-a2"
	A4              = "serval-a4"
	A6              = "serval-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: beforeActionA2,
			OnAfterAction:  removeEHRSkillA2,
		},
		CanModifySnapshot: true,
	})
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: A6Buff,
		},
	})
}

func initTraces(c *char) {
	if c.info.Traces["102"] {
		c.engine.ModifyEnergyFixed(info.ModifyAttribute{
			Key:    A4,
			Target: c.id,
			Source: c.id,
			Amount: 15.0,
		})
	}
}

// A2
func a2(c *char) {
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
			Stats:  info.PropMap{prop.EffectHitRate: 0.2},
		})
	}
}
func beforeActionA2(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType == model.AttackType_SKILL {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   A2,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.EffectHitRate: 0.2},
		})
	}
}
func removeEHRSkillA2(mod *modifier.Instance, e event.ActionEnd) {
	mod.Engine().RemoveModifier(mod.Owner(), A2)
}

// A6

func (c *char) a6() {
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}

func A6Buff(mod *modifier.Instance, target key.TargetID) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     A6,
		Source:   mod.Owner(),
		Duration: 2,
		Stats:    info.PropMap{prop.ATKPercent: 0.2},
	})
}
