package common

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Burn      key.Modifier = "common_burn"
	BreakBurn key.Modifier = "break_burn"
)

type BurnState struct {
	DamagePercentage    float64
	DamageValue         float64
	DEFDamagePercentage float64
}

func init() {
	// common burn
	modifier.Register(Burn, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_BURN,
		},
		Listeners: modifier.Listeners{
			OnPhase1: burnPhase1,
		},
	})

	// TODO: break burn
}

func burnPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(BurnState)
	if !ok {
		panic("incorrect state used for burn modifier")
	}

	// perform burn damage
	mod.Engine().Attack(info.Attack{
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.DamagePercentage,
			model.DamageFormula_BY_DEF: state.DEFDamagePercentage,
		},
		DamageValue: state.DamageValue,
		UseSnapshot: true,
	})
}
