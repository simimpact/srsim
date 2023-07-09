package modifier_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/tests/mock"
	"github.com/stretchr/testify/assert"
)

func NewTestManager(t *testing.T) (*modifier.Manager, *gomock.Controller) {
	mockCtrl := gomock.NewController(t)
	engine := mock.NewMockEngineWithEvents(mockCtrl)
	engine.EXPECT().IsValid(gomock.Any()).Return(true).AnyTimes()
	manager := modifier.NewManager(engine)
	engine.EXPECT().
		Stats(gomock.Any()).
		DoAndReturn(func(target key.TargetID) *info.Stats {
			attr := new(info.Attributes)
			*attr = info.DefaultAttribute()
			mods := manager.EvalModifiers(target)
			return info.NewStats(target, attr, mods)
		}).
		AnyTimes()
	return manager, mockCtrl
}

func TestOnPropertyChangeBuff(t *testing.T) {
	// 1. add permanent modifier with conditional buff if DEF% >= 10%
	// 3. EvalModifiers to show not applied
	// 4. add temporary modifier that gives +0.15 DEF%
	// 5. show state before and after modifier expires
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	conditionalMod := key.Modifier("TestOnPropertyChangeBuffMod1")
	otherMod := key.Modifier("TestOnPropertyChangeBuffMod2")
	target := key.TargetID(1)
	var mods info.ModifierState
	var expectedProps info.PropMap

	modifier.Register(conditionalMod, modifier.Config{
		Listeners: modifier.Listeners{
			OnPropertyChange: func(mod *modifier.Instance) {
				stats := mod.Engine().Stats(mod.Owner())
				if stats.GetProperty(prop.DEFPercent) >= 0.1 {
					mod.SetProperty(prop.AllDamagePercent, 0.1)
				} else {
					mod.SetProperty(prop.AllDamagePercent, 0.0)
				}
			},
		},
	})

	manager.AddModifier(target, info.Modifier{
		Name:   conditionalMod,
		Source: target,
	})

	mods = manager.EvalModifiers(target)
	assert.Empty(t, mods.Props, "conditional mod was incorrectly applied")

	manager.AddModifier(target, info.Modifier{
		Name:            otherMod,
		Source:          target,
		Duration:        1,
		TickImmediately: true,
		Stats:           info.PropMap{prop.DEFPercent: 0.15},
	})

	mods = manager.EvalModifiers(target)
	expectedProps = info.PropMap{
		prop.DEFPercent:       0.15,
		prop.AllDamagePercent: 0.1,
	}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, info.ActionEnd)
	manager.Tick(target, info.ModifierPhase2)
	mods = manager.EvalModifiers(target)
	assert.Empty(t, mods.Props, "all modifiers were not removed")
}

func TestReplaceStacking(t *testing.T) {
	// 1. Register modifier w/ max stacks of 5
	// 2. Add 1 stack
	// 3. tick forward time
	// 4. add 2 independent stacks (check that duration resets)
	// 5. tick forward time
	// 6. add +3 stacks
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	mod := key.Modifier("TestReplaceStacking")
	target := key.TargetID(1)
	var mods info.ModifierState
	var expectedProps info.PropMap

	modifier.Register(mod, modifier.Config{
		MaxCount:          5,
		CountAddWhenStack: 1,
		TickMoment:        modifier.ModifierPhase1End,
		Stacking:          modifier.Replace,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.AddProperty(prop.CritChance, 0.05*mod.Count())
			},
		},
	})

	manager.AddModifier(target, info.Modifier{
		Name:     mod,
		Source:   target,
		Duration: 2,
	})

	mods = manager.EvalModifiers(target)
	expectedProps = info.PropMap{prop.CritChance: 0.05}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, info.TurnStart)
	manager.Tick(target, info.ModifierPhase1)

	manager.AddModifier(target, info.Modifier{
		Name:     mod,
		Source:   target,
		Duration: 2,
	})
	manager.AddModifier(target, info.Modifier{
		Name:     mod,
		Source:   target,
		Duration: 2,
	})

	mods = manager.EvalModifiers(target)
	expectedProps = info.PropMap{prop.CritChance: 0.05 * float64(int(3))}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, info.ModifierPhase1)
	manager.Tick(target, info.TurnStart)

	manager.AddModifier(target, info.Modifier{
		Name:     mod,
		Source:   target,
		Duration: 2,
		Count:    3,
	})

	mods = manager.EvalModifiers(target)
	expectedProps = info.PropMap{prop.CritChance: 0.05 * float64(int(5))}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, info.ModifierPhase1)
	manager.Tick(target, info.TurnStart)
	manager.Tick(target, info.ModifierPhase1)
	manager.Tick(target, info.TurnStart)
	manager.Tick(target, info.ModifierPhase1)

	mods = manager.EvalModifiers(target)
	assert.Empty(t, mods.Props)
	assert.Empty(t, mods.Modifiers)
}

func TestReplaceStackingBySource(t *testing.T) {
	// 1. add mod from source A
	// 2. add mod from source B
	// 3. verify that you have 2 instances of mod
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	mod := key.Modifier("TestReplaceStackingBySource")
	srcA := key.TargetID(1)
	srcB := key.TargetID(2)
	target := key.TargetID(3)

	modifier.Register(mod, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.AddProperty(prop.QuantumPEN, 0.1)
			},
		},
	})

	manager.AddModifier(target, info.Modifier{
		Name:     mod,
		Source:   srcA,
		Duration: 2,
	})

	manager.AddModifier(target, info.Modifier{
		Name:     mod,
		Source:   srcB,
		Duration: 2,
	})

	mods := manager.EvalModifiers(target)
	expectedProps := info.PropMap{prop.QuantumPEN: 0.2}
	assert.Equal(t, expectedProps, mods.Props)
}

func TestTickImmediatelyBeforeAction(t *testing.T) {
	// 1. Add Mod w/ tick immediately (1 turn duration)
	// 2. tick action
	// 3. tick phase 2
	// 4. verify mod gone
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	var mods info.ModifierState
	mod := key.Modifier("TestTickImmediatelyBeforeAction")
	target := key.TargetID(3)

	manager.Tick(target, info.ModifierPhase1)

	manager.AddModifier(target, info.Modifier{
		Name:            mod,
		Source:          target,
		Duration:        1,
		TickImmediately: true,
		Stats:           info.PropMap{prop.ATKPercent: 0.1},
	})

	mods = manager.EvalModifiers(target)
	expectedProps := info.PropMap{prop.ATKPercent: 0.1}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, info.ActionEnd)
	manager.Tick(target, info.ModifierPhase2)

	mods = manager.EvalModifiers(target)
	assert.Empty(t, mods.Props)
}

func TestTickImmediatelyAfterAction(t *testing.T) {
	// 1. tick action
	// 2. Add Mod w/ tick immediately (1 turn duration)
	// 3. tick phase 2
	// 4. verify mod still there
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	var mods info.ModifierState
	mod := key.Modifier("TestTickImmediatelyBeforeAction")
	target := key.TargetID(3)

	manager.Tick(target, info.ModifierPhase1)
	manager.Tick(target, info.ActionEnd)

	manager.AddModifier(target, info.Modifier{
		Name:            mod,
		Source:          target,
		Duration:        1,
		TickImmediately: true,
		Stats:           info.PropMap{prop.ATKPercent: 0.1},
	})

	mods = manager.EvalModifiers(target)
	expectedProps := info.PropMap{prop.ATKPercent: 0.1}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, info.ModifierPhase2)

	mods = manager.EvalModifiers(target)
	assert.Equal(t, expectedProps, mods.Props)
}

func TestModifierRemovesInListener(t *testing.T) {
	// 1. add modifier a
	// 2. add modifier b
	// 3. add modiifer c
	// 4. a, b, and c listen for the same event
	// 5. a removes b
	// 6. assert that a and c are called
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	modA := key.Modifier("TestModifierRemovesInListener-A")
	modB := key.Modifier("TestModifierRemovesInListener-B")
	modC := key.Modifier("TestModifierRemovesInListener-C")
	target := key.TargetID(1)

	calls := make(map[key.Modifier]int, 3)

	listener := func(mod *modifier.Instance) {
		calls[mod.Name()] += 1
		if mod.Name() == modA {
			manager.RemoveModifier(target, modB)
		}
	}

	// register mods w/ listener
	modifier.Register(modA, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: listener,
		},
	})
	modifier.Register(modB, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: listener,
		},
	})
	modifier.Register(modC, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: listener,
		},
	})

	// add A
	manager.AddModifier(target, info.Modifier{
		Name:   modA,
		Source: target,
	})

	// add B
	manager.AddModifier(target, info.Modifier{
		Name:   modB,
		Source: target,
	})

	// add C
	manager.AddModifier(target, info.Modifier{
		Name:   modC,
		Source: target,
	})

	manager.Tick(target, info.ModifierPhase1)

	expectedCalls := map[key.Modifier]int{
		modA: 1,
		modB: 1,
		modC: 1,
	}
	expectedRemaining := []key.Modifier{modA, modC}
	state := manager.EvalModifiers(target)

	assert.Equal(t, expectedCalls, calls)
	assert.ElementsMatch(t, expectedRemaining, state.Modifiers)
}
