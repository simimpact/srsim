package common

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Freeze      key.Modifier = "common_freeze"
	BreakFreeze key.Modifier = "break_freeze"
)

type FreezeState struct {
	DamagePercentage float64
	DamageValue      float64
}

func init() {
	modifier.Register(Freeze, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_DISABLE_ACTION,
			model.BehaviorFlag_STAT_CTRL,
			model.BehaviorFlag_STAT_CTRL_FROZEN,
		},
		Listeners: modifier.Listeners{
			OnPhase1: freezePhase1,
		},
	})
}

func freezePhase1(mod *modifier.Instance) {
	state, ok := mod.State().(FreezeState)
	if !ok {
		panic("incorrect state used for freeze modifier")
	}

	// deal frozen damage
	mod.Engine().Attack(info.Attack{
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_ICE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.DamagePercentage,
		},
		DamageValue: state.DamageValue,
		UseSnapshot: true,
	})

	// if frozen is getting removed this turn, set their next turn to half-cost for the "thaw" effect
	if mod.Duration() <= 1 {
		mod.Engine().ModifyCurrentGaugeCost(0.5)
	}
}
