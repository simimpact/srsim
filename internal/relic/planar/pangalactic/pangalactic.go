package pangalactic

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "pan-galactic-commercial-enterprise"
)

// 2pc:
// Increases the wearer's Effect Hit Rate by 10%.
// Meanwhile, the wearer's ATK increases by an amount that is equal to
// 25% of the current Effect Hit Rate, up to a maximum of 25%.
func init() {
	relic.Register(key.PanGalactic, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.EffectHitRate: 0.10},
				CreateEffect: nil,
			},
			{
				MinCount: 2,
				Stats:    nil,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   mod,
						Source: owner,
					})
				},
			},
		},
	})

	// This is not 100% accurate to the game, as it only uses OnAdd and OnPropertyChange
	// Game hooks onto OnPhase1, OnPhase2, OnBeforeAction (all targets), OnAfterAction (all targets),
	// and OnBattleStart

	modifier.Register(mod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd:            onCheck,
			OnPropertyChange: onCheck,
		},
	})
}

func onCheck(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	atk := 0.25 * stats.GetProperty(prop.EffectHitRate)
	if atk >= 0.25 {
		atk = 0.25
	}
	mod.SetProperty(prop.ATKPercent, atk)
}
