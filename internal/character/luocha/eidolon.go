package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1          = "luocha-e1"
	E2HealBoost = "luocha-e2-healboost"
	E2Shield    = "luocha-e2-shield"
	E4          = "luocha-e4"
	E6          = "luocha-e6"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E2HealBoost, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeDealHeal: applyE2HealBoost,
		},
	})

	modifier.Register(E2Shield, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
		Duration:   2,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.Engine().AddShield(E2Shield, info.Shield{
					Source:      mod.Source(),
					Target:      mod.Owner(),
					BaseShield:  info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_ATK: 0.18},
					ShieldValue: 240,
				})
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(E2Shield, mod.Owner())
			},
		},
	})

	modifier.Register(E4, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_FATIGUE},
	})

	modifier.Register(E6, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		CanDispel:  true,
		Duration:   2,
	})
}

func applyE2HealBoost(mod *modifier.Instance, e *event.HealStart) {
	e.Healer.AddProperty(E2HealBoost, prop.HealBoost, 0.3)
	mod.RemoveSelf()
}
