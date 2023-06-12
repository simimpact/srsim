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
	A2 key.Modifier = "dan-heng-a2"
	A4 key.Modifier = "dan-heng-a4"
	A6 key.Modifier = "dan-heng-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			// set aggro down if starting at less than 50% HP
			OnAdd: func(mod *modifier.ModifierInstance) {
				if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
					mod.SetProperty(prop.AggroPercent, -0.5)
				}
			},

			// update aggro down based on new HP
			OnHPChange: func(mod *modifier.ModifierInstance, e event.HPChangeEvent) {
				if e.NewHPRatio <= 0.5 {
					mod.SetProperty(prop.AggroPercent, -0.5)
				} else {
					mod.SetProperty(prop.AggroPercent, 0)
				}
			},
		},
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
			OnBeforeHit: func(mod *modifier.ModifierInstance, e event.HitStartEvent) {
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
	if c.info.Traces["1002101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
		})
	}

	if c.info.Traces["1002103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}

// attempt to apply A4 for 2 turns w/ 50% chance
func (c *char) a4() {
	if c.info.Traces["1002102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     A4,
			Source:   c.id,
			Chance:   0.5,
			Duration: 2,
			Stats:    info.PropMap{prop.SPDPercent: 0.2},
		})
	}
}
