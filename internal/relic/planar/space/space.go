package space

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod = key.Modifier("space-sealing-station")
)

// 2pc:
// Increases the wearer's ATK by 12%. When the wearer's SPD reaches 120 or higher,
// the wearer's ATK increases by an extra 12%.
func init() {
	relic.Register(key.SpaceSealingStation, relic.Config{
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
	if stats.SPD() >= 120 {
		mod.SetProperty(prop.ATKPercent, 0.24)
	} else {
		mod.SetProperty(prop.ATKPercent, 0.12)
	}
}
