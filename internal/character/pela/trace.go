package pela

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A2:
// 	Deals 20% more DMG to debuffed enemies.
// A4:
//	When Pela is on the battlefield, all allies' Effect Hit Rate increases by 10%.
// A6:
//	Using Skill to remove buff(s) increases the DMG of the next attack by 20%.

const (
	A2 key.Modifier = "pela-a2"
	A4 key.Modifier = "pela-a4"
	A6 key.Modifier = "pela-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: func(mod *modifier.Instance, e event.HitStart) {
				if mod.Engine().ModifierStatusCount(e.Hit.Defender.ID(), model.StatusType_STATUS_DEBUFF) >= 1 {
					e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.2)
				}
			},
		},
	})

	modifier.Register(A4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.SetProperty(prop.EffectHitRate, 0.1)
			},
			OnBeforeDying: func(mod *modifier.Instance) {
				if mod.Owner() == mod.Source() {
					targets := mod.Engine().Characters()

					for _, trg := range targets {
						mod.Engine().RemoveModifier(trg, A4)
					}
				}
			},
		},
	})

	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: func(mod *modifier.Instance, e event.HitStart) {
				e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.2)
			},
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				mod.RemoveSelf()
			},
		},
	})
}

func (c *char) a6() {
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
		})
	}

	if c.info.Traces["102"] {
		c.engine.Events().CharactersAdded.Subscribe(func(e event.CharactersAdded) {
			for _, char := range e.Characters {
				c.engine.AddModifier(char.ID, info.Modifier{
					Name:   A4,
					Source: c.id,
				})
			}
		})
	}
}
