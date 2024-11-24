package valorous

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
	mod        key.Modifier = "wind-soaring-valorous"
	buffUltDmg key.Modifier = "wind-soaring-valorous-ult-buff"
)

// 2pc: Increases ATK by 12%.
// 4pc: Increases the wearer's CRIT Rate by 6%.
//
//	After the wearer uses follow-up attack,
//	increases DMG dealt by Ultimate by 36%, lasting for 1 turn(s).
func init() {
	relic.Register(key.WindSoaringValorous, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.ATKPercent: 0.12},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    info.PropMap{prop.CritChance: 0.06},
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
			OnBeforeAttack: onBeforeAttack,
		},
	})

	modifier.Register(buffUltDmg, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: doBuffUlt,
		},
	})
}

func onBeforeAttack(mod *modifier.Instance, e event.AttackStart) {
	if e.AttackType == model.AttackType_INSERT {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     buffUltDmg,
			Source:   mod.Owner(),
			Duration: 1,
		})
	}
}

func doBuffUlt(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_ULT {
		e.Hit.Attacker.AddProperty(key.Reason(buffUltDmg), prop.AllDamagePercent, 0.36)
	}
}
