package genius

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	genius = "genius-of-brilliant-stars"
)

// 2pc : Increases Quantum DMG by 10%
// 4pc : When the wearer deals DMG to the target enemy, ignores 10% DEF.
//       If the target enemy has Quantum Weakness, the wearer additionally ignores 10% DEF.

func init() {
	relic.Register(key.GeniusOfBrilliantStars, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.QuantumDamagePercent: 0.1},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   genius,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(genius, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: addResPen,
		},
		CanModifySnapshot: true,
	})
}

func addResPen(mod *modifier.Instance, e event.HitStart) {
	defPenAmt := 0.1
	// +10% def pen if target is weak to quantum
	if e.Hit.Defender.IsWeakTo(model.DamageType_QUANTUM) {
		defPenAmt = 0.2
	}
	// add def pen amount to hit target as def "debuff"
	e.Hit.Defender.AddProperty(genius, prop.DEFPercent, -defPenAmt)
}
