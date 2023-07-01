package common

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Shock      key.Modifier = "common_shock"
	BreakShock key.Modifier = "break_shock"
)

// This is the state that should be passed in when folks call AddModifier
type ShockState struct {
	DamagePercentage float64
	DamageValue      float64
}

type BreakShockState struct {
	DamagePercentage float64
}

func init() {
	// common shock
	modifier.Register(Shock, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_ELECTRIC,
		},
		Listeners: modifier.Listeners{
			OnPhase1: shockPhase1,
		},
	})

	// break shock
	modifier.Register(BreakShock, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_ELECTRIC,
		},
		Listeners: modifier.Listeners{
			OnPhase1: breakShockPhase1,
		},
	})
}

func shockPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(ShockState)
	if !ok {
		panic("incorrect state used for shock modifier")
	}

	// perform shock damage
	mod.Engine().Attack(info.Attack{
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_THUNDER,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.DamagePercentage,
		},
		DamageValue: state.DamageValue,
		UseSnapshot: true,
	})
}

func breakShockPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(BreakShockState)
	if !ok {
		panic("incorrect state used for break shock modifier")
	}

	// perform break shock damage
	mod.Engine().Attack(info.Attack{
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_THUNDER,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: state.DamagePercentage,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}
