package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestRemoveModifierUnknownTarget(t *testing.T) {
	target := key.TargetID(1)
	mod := key.Modifier("Test")

	manager := Manager{
		engine:    nil,
		targets:   make(map[key.TargetID]activeModifiers),
		turnCount: 0,
	}

	manager.RemoveModifier(target, mod)
}

func TestRemoveModifierNoOp(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	mod := key.Modifier("Test")

	other := &Instance{
		name: key.Modifier("Other"),
	}

	manager.targets[target] = append(manager.targets[target], other)
	manager.RemoveModifier(target, mod)
	assert.Equal(t, other, manager.targets[target][0])
}

func TestRemoveModifierFromSourceNoOp(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	mod := key.Modifier("Test")

	other := &Instance{
		name:   key.Modifier("Test"),
		source: key.TargetID(2),
	}

	manager.targets[target] = append(manager.targets[target], other)
	manager.RemoveModifierFromSource(target, target, mod)
	assert.Equal(t, other, manager.targets[target][0])
}

func TestRemoveModifier(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	modsToRemove := key.Modifier("ToRemove")

	mod1 := &Instance{
		name: key.Modifier("Other"),
	}
	mod2 := &Instance{
		name:   modsToRemove,
		source: target,
	}
	mod3 := &Instance{
		name:   modsToRemove,
		source: key.TargetID(3),
	}
	manager.targets[target] = append(manager.targets[target], mod3, mod1, mod2)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemoved) {
		assert.Equal(t, modsToRemove, event.Modifier.Name)
		called += 1
	})

	manager.RemoveModifier(target, modsToRemove)
	assert.Len(t, manager.targets[target], 1)
	assert.Equal(t, mod1, manager.targets[target][0])
	assert.Equal(t, 2, called)
}

func TestRemoveModifierFromSource(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	modsToRemove := key.Modifier("ToRemove")

	mod1 := &Instance{
		name:   modsToRemove,
		source: key.TargetID(2),
	}
	mod2 := &Instance{
		name:   modsToRemove,
		source: target,
	}
	mod3 := &Instance{
		name:   modsToRemove,
		source: key.TargetID(3),
	}
	manager.targets[target] = append(manager.targets[target], mod3, mod2, mod1)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemoved) {
		assert.Equal(t, modsToRemove, event.Modifier.Name)
		called += 1
	})

	manager.RemoveModifierFromSource(target, target, modsToRemove)
	assert.Len(t, manager.targets[target], 2)
	assert.Equal(t, mod3, manager.targets[target][0])
	assert.Equal(t, mod1, manager.targets[target][1])
	assert.Equal(t, 1, called)
}

func TestRemoveModifierWithOnRemoveListener(t *testing.T) {
	type state struct {
		OnRemoveCalled bool
	}

	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestRemoveModifierWithListener")

	mod := &Instance{
		name:  name,
		state: &state{OnRemoveCalled: false},
		listeners: Listeners{
			OnRemove: func(modifier *Instance) {
				state := modifier.State().(*state)
				state.OnRemoveCalled = true
			},
		},
	}

	manager.targets[target] = append(manager.targets[target], mod)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemoved) {
		state := event.Modifier.State.(*state)
		assert.Equal(t, name, event.Modifier.Name)
		assert.True(t, state.OnRemoveCalled)
		called += 1
	})

	manager.RemoveModifier(target, name)
	assert.Empty(t, manager.targets[target])
	assert.Equal(t, 1, called)
}

func TestRemoveModifierSelf(t *testing.T) {
	type state struct {
		OnRemoveCalled bool
	}

	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestRemoveModifierSelf")

	mod := &Instance{
		name:  name,
		owner: target,
		state: &state{OnRemoveCalled: false},
		listeners: Listeners{
			OnRemove: func(modifier *Instance) {
				state := modifier.State().(*state)
				state.OnRemoveCalled = true
			},
		},
		manager: manager,
	}

	manager.targets[target] = append(manager.targets[target], mod)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemoved) {
		state := event.Modifier.State.(*state)
		assert.Equal(t, name, event.Modifier.Name)
		assert.True(t, state.OnRemoveCalled)
		called += 1
	})

	mod.RemoveSelf()
	assert.Empty(t, manager.targets[target])
	assert.Equal(t, 1, called)
}

func TestDispelStatus(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	statusToDispel := model.StatusType_STATUS_BUFF // Example status type to dispel

	mod1 := &Instance{
		name:       key.Modifier("Mod1"),
		source:     target,
		statusType: statusToDispel,
		canDispel:  true,
	}
	mod2 := &Instance{
		name:       key.Modifier("Mod2"),
		source:     target,
		statusType: statusToDispel,
		canDispel:  false, // Not dispellable
	}
	mod3 := &Instance{
		name:       key.Modifier("Mod3"),
		source:     target,
		statusType: model.StatusType_STATUS_DEBUFF, // Different status type
		canDispel:  true,
	}

	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3)

	dispel := info.Dispel{
		Status: statusToDispel,
		Count:  1, // Dispel only one
		Order:  model.DispelOrder_LAST_ADDED,
	}

	called := 0
	manager.engine.Events().ModifierDispelled.Subscribe(func(event event.ModifierDispelled) {
		assert.Equal(t, key.Modifier("Mod1"), event.Modifier.Name) // Expect Mod1
		called++
	})

	manager.DispelStatus(target, dispel)

	// Verify that only mod1 was removed (last dispellable of the specified type)
	assert.Len(t, manager.targets[target], 2)
	assert.NotContains(t, manager.targets[target], mod1)
	assert.Contains(t, manager.targets[target], mod2)
	assert.Contains(t, manager.targets[target], mod3)

	assert.Equal(t, 1, called)
}
