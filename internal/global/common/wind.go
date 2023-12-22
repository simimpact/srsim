package common

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	WindShear      = "common-wind-shear"
	BreakWindShear = "break-wind-shear"
)

type WindShearState struct {
	DamagePercentage float64
}

func init() {
	modifier.Register(WindShear, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          5,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_POISON,
		},
		Listeners: modifier.Listeners{
			OnPhase1: windShearPhase1,
		},
	})

	// break wind shear
	modifier.Register(BreakWindShear, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          5,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_POISON,
		},
		Listeners: modifier.Listeners{
			OnPhase1: breakWindShearPhase1,
		},
	})
}

func windShearPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*WindShearState)
	if !ok {
		panic("incorrect state used for wind shear modifier")
	}

	// perform wind shear damage
	mod.Engine().Attack(info.Attack{
		Key:        WindShear,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_WIND,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.DamagePercentage * mod.Count(),
		},
		UseSnapshot: true,
	})
}

func breakWindShearPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*WindShearState)
	if !ok {
		panic("incorrect state used for wind shear modifier")
	}

	// perform break wind shear damage
	mod.Engine().Attack(info.Attack{
		Key:        BreakWindShear,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_WIND,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: state.DamagePercentage * mod.Count(),
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}
