package thief

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const thief = "thief-of-shooting-meteor"

// 2pc: Increases Break Effect by 16%.
// 4pc: Increases the wearer's Break Effect by 16%.
//      After the wearer inflicts Weakness Break on an enemy, regenerates 3 Energy.

func init() {
	relic.Register(key.ThiefOfShootingMeteor, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.BreakEffect: 0.16},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    info.PropMap{prop.BreakEffect: 0.16},
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   thief,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(thief, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerBreak: onTriggerBreak,
		},
	})
}

func onTriggerBreak(mod *modifier.Instance, target key.TargetID) {
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    thief,
		Target: mod.Owner(),
		Source: mod.Owner(),
		Amount: 3.0,
	})
}
