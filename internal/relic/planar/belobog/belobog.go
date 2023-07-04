package belobog

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "belobog-of-the-architects"
)

// Increases the wearer's DEF by 15%. When the wearer's Effect Hit Rate is 50% or higher,
// the wearer gains an extra 15% DEF.
func init() {
	relic.Register(key.BelobogOfTheArchitects, relic.Config{
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
	if stats.EffectHitRate() >= 0.5 {
		mod.SetProperty(prop.DEFPercent, 0.30)
	} else {
		mod.SetProperty(prop.DEFPercent, 0.15)
	}
}
