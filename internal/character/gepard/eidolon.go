package gepard

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	E2    key.Modifier = "gepard-e2"
	E4    key.Modifier = "gepard-e4"
	E4Res key.Modifier = "gepard-e4-res"
)

func init() {
	modifier.Register(E4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				targets := mod.Engine().Characters()

				for _, trg := range targets {
					mod.Engine().AddModifier(trg, info.Modifier{
						Name:   E4Res,
						Source: mod.Owner(),
					})
				}

				mod.Engine().Events().CharacterAdded.Subscribe(func(e event.CharacterAddedEvent) {
					mod.Engine().AddModifier(e.Id, info.Modifier{
						Name:   E4Res,
						Source: mod.Owner(),
					})
				})
			},
			OnBeforeDying: func(mod *modifier.ModifierInstance) {
				targets := mod.Engine().Characters()

				for _, trg := range targets {
					mod.Engine().RemoveModifier(trg, E4Res)
				}
			},
		},
	})

	modifier.Register(E4Res, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.EffectRES, 0.2)
			},
		},
	})
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
			if event.Modifier.Name == common.Freeze && c.engine.HasModifier(event.Target, E2Tracker) {
				c.engine.AddModifier(event.Target, info.Modifier{
					Name:     E2,
					Source:   c.id,
					Stats:    info.PropMap{prop.SPDPercent: -0.2},
					Duration: 1,
				})

				c.engine.RemoveModifier(event.Target, E2Tracker)
			}
		})
	}
}

func (c *char) e4() {
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
		})
	}
}
