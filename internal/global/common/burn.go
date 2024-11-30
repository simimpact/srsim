package common

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Burn      = "common-burn"
	BreakBurn = "break-burn"
)

type BurnState struct {
	DamagePercentage    float64
	DamageValue         float64
	DEFDamagePercentage float64
}

type BreakBurnState struct {
	BreakBaseMulti float64
}

func init() {
	// common burn
	modifier.Register(Burn, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		CanDispel:         true,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_BURN,
		},
		Listeners: modifier.Listeners{
			OnPhase1: burnPhase1,
		},
	})

	// break burn
	modifier.Register(BreakBurn, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		CanDispel:         true,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_BURN,
		},
		Listeners: modifier.Listeners{
			OnPhase1: breakBurnPhase1,
		},
	})
}

func burnPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*BurnState)
	if !ok {
		panic("incorrect state used for burn modifier")
	}

	// perform burn damage
	mod.Engine().Attack(info.Attack{
		Key:        Burn,
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

func breakBurnPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*BreakBurnState)
	if !ok {
		panic("incorrect state used for burn modifier")
	}

	// perform break burn damage
	mod.Engine().Attack(info.Attack{
		Key:        BreakBurn,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: state.BreakBaseMulti,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}

func (b BurnState) TriggerDot(mod info.Modifier, ratio float64, engine engine.Engine, target key.TargetID) {
	// perform burn damage
	engine.Attack(info.Attack{
		Key:        Burn,
		Source:     mod.Source,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: b.DamagePercentage * ratio,
			model.DamageFormula_BY_DEF: b.DEFDamagePercentage * ratio,
		},
		DamageValue: b.DamageValue * ratio,
		UseSnapshot: true,
	})
}

func (b BreakBurnState) TriggerDot(mod info.Modifier, ratio float64, engine engine.Engine, target key.TargetID) {
	// perform burn damage
	engine.Attack(info.Attack{
		Key:        BreakBurn,
		Source:     mod.Source,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_FIRE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: b.BreakBaseMulti * ratio,
		},
		UseSnapshot: true,
	})
}
