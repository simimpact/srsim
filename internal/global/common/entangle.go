package common

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Entanglement      = "common-entanglement"
	BreakEntanglement = "break-entanglement"
)

type EntangleState struct {
	DelayRatio       float64
	DamagePercentage float64
	DamageValue      float64
	count            float64
}

func init() {
	modifier.Register(Entanglement, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		TickMoment: modifier.ModifierPhase1End,
		Duration:   1,
		Count:      1,
		StatusType: model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_DISABLE_ACTION,
			model.BehaviorFlag_STAT_ENTANGLE,
			model.BehaviorFlag_STAT_CTRL,
		},
		Listeners: modifier.Listeners{
			OnAdd:                entangleAdd,
			OnPhase1:             entanglePhase1,
			OnAfterBeingAttacked: entangleAfterAttack,
		},
	})

	// TODO: Break Entanglement
}

func entangleAdd(mod *modifier.Instance) {
	state, ok := mod.State().(EntangleState)
	if !ok {
		panic("incorrect state used for Entanglement modifier")
	}

	mod.Engine().ModifyGaugeNormalized(info.ModifyAttribute{
		Key:    Entanglement,
		Target: mod.Owner(),
		Source: mod.Source(),
		Amount: state.DelayRatio,
	})
}

func entangleAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(EntangleState)

	// increase count by 1 for each attack within this state
	state.count += 1
	if state.count > 4 {
		state.count = 4
	}
}

func entanglePhase1(mod *modifier.Instance) {
	state := mod.State().(EntangleState)

	// perform quantum damage
	mod.Engine().Attack(info.Attack{
		Key:        Entanglement,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: (1 + state.count) * state.DamagePercentage,
		},
		DamageValue: state.DamageValue,
		UseSnapshot: true,
	})
}
