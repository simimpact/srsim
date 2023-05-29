package space

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	mod = key.Modifier("space-sealing-station")
)

func init() {
	relic.Register(key.SpaceSealingStation, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{model.Property_ATK_PERCENT: 0.12},
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   mod,
						Source: owner,
					})
				},
			},
		},
	})

	// add +0.12 ATK% if SPD >= 120
	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				stats := mod.OwnerStats()
				if stats.SPD() >= 120 {
					mod.SetProperty(model.Property_ATK_PERCENT, 0.12)
				}
			},
			OnPropertyChange: func(mod *modifier.ModifierInstance) {
				stats := mod.OwnerStats()
				if stats.SPD() >= 120 {
					mod.SetProperty(model.Property_ATK_PERCENT, 0.12)
				} else {
					mod.SetProperty(model.Property_ATK_PERCENT, 0.0)
				}
			},
		},
	})
}
