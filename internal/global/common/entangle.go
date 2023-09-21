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

type EntanglementState struct {
	DelayRatio       float64
	DamagePercentage float64
	DamageValue      float64
	HitsTakenCount   float64
}

type BreakEntanglementState struct {
	HitsTakenCount            float64
	TargetMaxStanceMultiplier float64
}

func init() {
	/// entanglement
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
			OnAdd:                entanglementAdd,
			OnPhase1:             entanglementPhase1,
			OnAfterBeingAttacked: entanglementAfterAttack,
		},
	})

	// break entanglement
	modifier.Register(BreakEntanglement, modifier.Config{
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
			OnAdd:                breakEntanglementAdd,
			OnPhase1:             breakEntanglementPhase1,
			OnAfterBeingAttacked: breakEntanglementAfterAttack,
		},
	})
}

func entanglementAdd(mod *modifier.Instance) {
	state, ok := mod.State().(*EntanglementState)

	if !ok {
		panic("incorrect state used for entanglement modifier")
	}

	mod.Engine().ModifyGaugeNormalized(info.ModifyAttribute{
		Key:    Entanglement,
		Target: mod.Owner(),
		Source: mod.Source(),
		Amount: state.DelayRatio,
	})
}

func entanglementAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	state, ok := mod.State().(*EntanglementState)

	if !ok {
		panic("incorrect state used for entanglement modifier")
	}

	// increase count by 1 for each attack within this state
	state.HitsTakenCount += 1
	if state.HitsTakenCount > 4 {
		state.HitsTakenCount = 4
	}
}

func entanglementPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*EntanglementState)

	if !ok {
		panic("incorrect state used for entanglement modifier")
	}

	// perform entanglement damage
	mod.Engine().Attack(info.Attack{
		Key:        Entanglement,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: (1 + state.HitsTakenCount) * state.DamagePercentage,
		},
		DamageValue: state.DamageValue,
		UseSnapshot: true,
	})
}

func breakEntanglementAdd(mod *modifier.Instance) {
	mod.Engine().ModifyGaugeNormalized(info.ModifyAttribute{
		Key:    BreakEntanglement,
		Target: mod.Owner(),
		Source: mod.Source(),
		Amount: 0.2 * (1 + mod.Engine().Stats(mod.Source()).BreakEffect()),
	})
}

func breakEntanglementAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	state, ok := mod.State().(*BreakEntanglementState)

	if !ok {
		panic("incorrect state used for entanglement modifier")
	}

	// increase count by 1 for each attack within this state
	state.HitsTakenCount += 1
	if state.HitsTakenCount > 4 {
		state.HitsTakenCount = 4
	}
}

func breakEntanglementPhase1(mod *modifier.Instance) {
	state, ok := mod.State().(*BreakEntanglementState)

	if !ok {
		panic("incorrect state used for entanglement modifier")
	}

	// perform entanglement damage
	mod.Engine().Attack(info.Attack{
		Key:        BreakEntanglement,
		Source:     mod.Source(),
		Targets:    []key.TargetID{mod.Owner()},
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_BREAK_DAMAGE: state.HitsTakenCount * state.TargetMaxStanceMultiplier * 0.6,
		},
		AsPureDamage: true,
		UseSnapshot:  true,
	})
}
