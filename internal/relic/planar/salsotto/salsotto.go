package salsotto

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
	mod = key.Modifier("inert-salsotto")
)

// 2pc:
// Increases the wearer's CRIT Rate by 8%.
// When the wearer's current CRIT Rate reaches 50% or higher, the wearer's Ultimate and follow-up attack DMG increases by 15%.

func init() {
	relic.Register(key.InertSalsotto, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.CritChance: 0.08},
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
			OnBeforeHit: onBeforeHit,
		},
	})
}

func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	stats := mod.OwnerStats()
	if stats.CritChance() >= 0.50 {
		if e.Hit.AttackType == model.AttackType_ULT || e.Hit.AttackType == model.AttackType_INSERT {
			e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.15)
		}
	}
}
