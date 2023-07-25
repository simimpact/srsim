package talia

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "talia-kingdom-of-banditry"
)

// 2pc:
// Increases the wearer's Break Effect by 16%.
// When the wearer's SPD reaches 145 or higher, the wearer's Break effect increases by an extra 20%.
func init() {
	relic.Register(key.TaliaKingdomOfBanditry, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   mod,
						Source: owner,
					})
				},
			},
		},
	})

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:            onCheck,
			OnPropertyChange: onCheck,
		},
	})
}

func onCheck(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	if stats.SPD() >= 145 {
		mod.SetProperty(prop.BreakEffect, 0.36)
	} else {
		mod.SetProperty(prop.BreakEffect, 0.16)
	}
}
