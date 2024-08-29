package glamoth

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	name = "glamoth"
)

// 2pc:
// Increases the wearer's ATK by 12%. When the wearer's SPD is equal to or higher,
// than 135/160, the wearer deals 12%/18% more DMG.

func init() {
	relic.Register(key.Glamoth, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.ATKPercent: 0.12},
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   name,
						Source: owner,
					})
				},
			},
		},
	})

	modifier.Register(name, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:            onCheck,
			OnPropertyChange: onCheck,
		},
	})
}

func onCheck(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	if stats.SPD() >= 160 {
		mod.SetProperty(prop.AllDamagePercent, 0.18)
	} else if stats.SPD() >= 135 {
		mod.SetProperty(prop.AllDamagePercent, 0.12)
	}
}