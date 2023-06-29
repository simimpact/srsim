package gepard

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 key.Modifier = "gepard-e2"
	E4 key.Modifier = "gepard-e4"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:      modifier.Replace,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_DOWN},
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Duration:      1,
	})

	modifier.Register(E4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.EffectRES, 0.2)
			},
			OnBeforeDying: func(mod *modifier.ModifierInstance) {
				if mod.Owner() == mod.Source() {
					targets := mod.Engine().Characters()

					for _, trg := range targets {
						mod.Engine().RemoveModifier(trg, E4)
					}
				}
			},
		},
	})
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
			if event.Modifier.Name == common.Freeze && c.engine.HasModifier(event.Target, E2Tracker) {
				c.engine.AddModifier(event.Target, info.Modifier{
					Name:   E2,
					Source: c.id,
					Stats:  info.PropMap{prop.SPDPercent: -0.2},
				})

				c.engine.RemoveModifier(event.Target, E2Tracker)
			}
		})
	}
}

func (c *char) e4() {
	if c.info.Eidolon >= 4 {
		targets := c.engine.Characters()

		for _, trg := range targets {
			c.engine.AddModifier(trg, info.Modifier{
				Name:   E4,
				Source: c.id,
			})
		}

		c.engine.Events().CharacterAdded.Subscribe(func(e event.CharacterAddedEvent) {
			c.engine.AddModifier(e.Id, info.Modifier{
				Name:   E4,
				Source: c.id,
			})
		})
	}
}
