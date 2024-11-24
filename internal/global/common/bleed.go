package common

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Bleed      = "common-bleed"
	BreakBleed = "break-bleed"
)

type BleedState struct {
	DamagePercentage float64
	DamageValue      float64
}

type BreakBleedState struct {
	BaseDamageValue float64
}

func init() {
	// common bleed
	modifier.Register(Bleed, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		CanDispel:         true,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_BLEED,
		},
		Listeners: modifier.Listeners{
			OnPhase1: bleedPhase1,
		},
	})

	// break bleed
	modifier.Register(BreakBleed, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		MaxCount:          1,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		CanDispel:         true,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT,
			model.BehaviorFlag_STAT_DOT_BLEED,
		},
		Listeners: modifier.Listeners{
			OnPhase1: breakBleedPhase1,
		},
	})
}

func bleedPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*BleedState)

	if !ok {
		panic("incorrect state used for bleed modifier")
	}

	// perform bleed damage
	mod.Engine().Attack(info.Attack{
		Key:        Bleed,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.DamagePercentage,
		},
		DamageValue: state.DamageValue,
		UseSnapshot: true,
	})
}

func breakBleedPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*BreakBleedState)

	if !ok {
		panic("incorrect state used for break bleed modifier")
	}

	// perform break bleed damage
	mod.Engine().Attack(info.Attack{
		Key:        BreakBleed,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: 0,
		},
		DamageValue:  state.BaseDamageValue,
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}

// Custom event trigger for bleed dots
func (b BleedState) TriggerDot(mod info.Modifier, ratio float64, engine engine.Engine, target key.TargetID) {
	// perform bleed damage
	engine.Attack(info.Attack{
		Key:        Bleed,
		Source:     mod.Source,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: b.DamagePercentage * ratio,
		},
		DamageValue: b.DamageValue,
		UseSnapshot: true,
	})
}

// Ditto, but for break dots
func (b BreakBleedState) TriggerDot(mod info.Modifier, ratio float64, engine engine.Engine, target key.TargetID) {
	// perform break bleed damage
	engine.Attack(info.Attack{
		Key:        BreakBleed,
		Source:     mod.Source,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_PHYSICAL,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: 0,
		},
		DamageValue:  b.BaseDamageValue * ratio,
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}
