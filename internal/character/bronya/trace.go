package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 key.Modifier = "bronya-a2"
	A4 key.Modifier = "bronya-a4"
	A6 key.Modifier = "bronya-a6"
)

func init() {

	// A2 Register
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.ModifierInstance, e event.HitStartEvent) {
				if e.Hit.AttackType == model.AttackType_NORMAL {
					e.Hit.Attacker.AddProperty(prop.CritChance, 1)
				}
			},
		},
	})

	// A4 Register
	modifier.Register(A4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.DEFPercent, 0.2)
			},
		},
		Duration: 2,
	})

	// A6 Register
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.AllDamagePercent, 0.1)
			},
			OnBeforeDying: func(mod *modifier.ModifierInstance) {
				if mod.Owner() == mod.Source() {
					targets := mod.Engine().Characters()

					for _, trg := range targets {
						mod.Engine().RemoveModifier(trg, A6)
					}
				}
			},
		},
	})
}

func (c *char) initTraces() {

	// A2
	if c.info.Traces["1101101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
		})
	}

	// A4
	if c.info.Traces["1101102"] {
		targets := c.engine.Characters()

		for _, trg := range targets {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   A4,
				Source: c.id,
			})
		}

		c.engine.Events().CharacterAdded.Subscribe(func(e event.CharacterAddedEvent) {
			c.engine.AddModifier(e.ID, info.Modifier{
				Name:   A4,
				Source: c.id,
			})
		})
	}

	// A6
	if c.info.Traces["1101103"] {
		targets := c.engine.Characters()

		for _, trg := range targets {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   A6,
				Source: c.id,
			})
		}

		c.engine.Events().CharacterAdded.Subscribe(func(e event.CharacterAddedEvent) {
			c.engine.AddModifier(e.ID, info.Modifier{
				Name:   A6,
				Source: c.id,
			})
		})
	}
}
