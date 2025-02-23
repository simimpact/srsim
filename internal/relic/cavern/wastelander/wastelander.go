package wastelander

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
	wastelander     = "wastelander-of-banditry-desert"
	wastelandercr   = "wastelander-of-banditry-desert-cr"
	wastelandercdmg = "wastelander-of-banditry-desert-cdmg"
)

// 2pc: Increases Imaginary DMG by 10%.
// 4pc: When attacking debuffed enemies, the wearer's CRIT Rate increases by 10%,
//      and their CRIT DMG increases by 20% against Imprisoned enemies.

func init() {
	relic.Register(key.WastelanderOfBanditryDesert, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.ImaginaryDamagePercent: 0.1},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    nil,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   wastelander,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(wastelander, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})
}

func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	debuffCount := e.Hit.Defender.StatusCount(model.StatusType_STATUS_DEBUFF)
	if debuffCount >= 1 {
		e.Hit.Attacker.AddProperty(wastelandercr, prop.CritChance, 0.1)
		if mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_CONFINE) {
			e.Hit.Attacker.AddProperty(wastelandercdmg, prop.CritDMG, 0.2)
		}
	}
}
