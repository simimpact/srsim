package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A2:
//  When current HP percentage is 50% or lower, reduces the chance of being attacked by enemies.
// A4:
//  For every Sword Stance triggered, the DMG dealt by Sword Stance increases by 2.5%. Stacks up to
//  10 time(s).
// A6:
//  After using Basic ATK or Skill, if there are enemies on the field with Weakness Break,
//  Sushang's action is Advanced Forward by 15%.

const (
	A2Check key.Modifier = "sushang-a2-check"
	A2Buff  key.Modifier = "sushang-a2-buff"
	A4Mod   key.Modifier = "sushang-a4-mod"
	A4Buff  key.Modifier = "sushang-a4-buff"
)

func init() {
	// checks if we need to add/remove the A2 buff
	modifier.Register(A2Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: a2HPCheck,
			OnHPChange: func(mod *modifier.ModifierInstance, e event.HPChangeEvent) {
				a2HPCheck(mod)
			},
		},
	})

	// A2 aggro down buff
	modifier.Register(A2Buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})

	// applies a4 buff
	modifier.Register(A4Mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: a4OnBeforeHitAll,
		},
	})

	// A4 dmg buff
	modifier.Register(A4Buff, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		StatusType:        model.StatusType_STATUS_BUFF,
		MaxCount:          10,
		CountAddWhenStack: 1,
	})
}

// add A2 on init
func (c *char) initTraces() {
	if c.info.Traces["1206101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2Check,
			Source: c.id,
		})
	}
}

func a2HPCheck(mod *modifier.ModifierInstance) {
	if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   A2Buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AggroPercent: -0.5},
		})
	} else {
		mod.Engine().RemoveModifier(mod.Owner(), A2Buff)
	}
}

func a4OnBeforeHitAll(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	if mod.Engine().HasModifier(mod.Owner(), A4Buff) {
		stacks := mod.Engine().GetModifiers(mod.Owner(), A4Buff)[0].Count
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, stacks*0.025)
	}
}

func (c *char) a4AddStack() {
	if c.engine.HasModifier(c.id, A4Buff) {
		stacks := c.engine.GetModifiers(c.id, A4Buff)[0].Count
		if stacks == 10 {
			return
		}
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:   A4Buff,
		Source: c.id,
	})

}

func (c *char) a6() {
	if c.info.Traces["1206103"] {
		for _, enemy := range c.engine.Enemies() {
			if c.engine.Stats(enemy).Stance() == 0 {
				c.engine.ModifyCurrentGaugeCost(-0.15)
				break
			}
		}
	}
}
