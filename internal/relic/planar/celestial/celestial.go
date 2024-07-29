package celestial

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

// 2pc: Increases the wearer's CRIT DMG by 16%.
//      When the wearer's current CRIT DMG reaches 120% or higher, after entering battle,
//      the wearer's CRIT Rate increases by 60% until the end of their first attack.

// Multi Wave Support not implemented

const celestial = "celestial-differentiator"

func init() {
	relic.Register(key.CelestialDifferentiator, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.CritDMG: 0.16},
			},
			{
				MinCount: 2,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   celestial,
						Source: owner,
						Stats:  info.PropMap{prop.CritChance: 0.6},
					})
				},
			},
		},
	})
	modifier.Register(celestial, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: onAfterAttack,
		},
	})
}

func onAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}
