package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A2:
// 	When current HP percentage is 50% or lower, reduces the chance of being attacked by enemies.
// A4:
//	After launching an attack, there is a 50% fixed chance to increase own SPD by 20% for 2 turn(s).
// A6:
//	Basic ATK deals 40% more DMG to Slowed enemies.

const (
	A2Check key.Modifier = "dan-heng-a2-check"
	A2Buff  key.Modifier = "dan-heng-a2-buff"
	A4      key.Modifier = "dan-heng-a4"
	A6      key.Modifier = "dan-heng-a6"
)

func init() {
	// checks if we need to add/remove the A2 buff
	modifier.Register(A2Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: a2HPCheck,
			OnHPChange: func(mod *modifier.Instance, e event.HPChange) {
				a2HPCheck(mod)
			},
		},
	})

	// A2 aggro down buff
	modifier.Register(A2Buff, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})

	// A4 metadata
	modifier.Register(A4, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_SPEED_UP,
		},
	})

	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.AttackType != model.AttackType_NORMAL {
					return
				}

				if mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_SPEED_DOWN) {
					e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.4)
				}
			},
		},
	})
}

// add A2 & A6 on init
func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2Check,
			Source: c.id,
		})
	}

	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}

func a2HPCheck(mod *modifier.Instance) {
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

// attempt to apply A4 for 2 turns w/ 50% fixed chance
func (c *char) a4() {
	if c.info.Traces["102"] && c.engine.Rand().Float64() < 0.5 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     A4,
			Source:   c.id,
			Duration: 2,
			Stats:    info.PropMap{prop.SPDPercent: 0.2},
		})
	}
}
