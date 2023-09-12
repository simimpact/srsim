package rutilant

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
	name = "rutilant-arena"
)

// 2pc: Increases the wearer's CRIT Rate by 8%. When the wearer's CRIT Rate reaches 70% or higher, the wearer's Basic ATK and Skill DMG increase by 20%.

func init() {
	relic.Register(key.RutilantArena, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.CritChance: 0.08},
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
			OnBeforeHit: onBeforeHit,
		},
	})
}

func onBeforeHit(mod *modifier.Instance, e event.HitStart) {
	stats := mod.OwnerStats()
	if stats.CritChance() >= 0.70 {
		if e.Hit.AttackType == model.AttackType_NORMAL || e.Hit.AttackType == model.AttackType_SKILL {
			e.Hit.Attacker.AddProperty(name, prop.AllDamagePercent, 0.20)
		}
	}
}
