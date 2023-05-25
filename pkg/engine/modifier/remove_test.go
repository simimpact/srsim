package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/stretchr/testify/assert"
)

func TestRemoveModifierUnknownTarget(t *testing.T) {
	target := key.TargetID(1)
	mod := key.Modifier("Test")

	manager := Manager{
		targets: make(map[key.TargetID]activeModifiers),
	}

	manager.RemoveModifier(target, mod)
}

func TestRemoveModifierNoOp(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	mod := key.Modifier("Test")

	other := &ModifierInstance{
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

	other := &ModifierInstance{
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

	mod1 := &ModifierInstance{
		name: key.Modifier("Other"),
	}
	mod2 := &ModifierInstance{
		name:   modsToRemove,
		source: target,
	}
	mod3 := &ModifierInstance{
		name:   modsToRemove,
		source: key.TargetID(3),
	}
	manager.targets[target] = append(manager.targets[target], mod3, mod1, mod2)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
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

	mod1 := &ModifierInstance{
		name:   modsToRemove,
		source: key.TargetID(2),
	}
	mod2 := &ModifierInstance{
		name:   modsToRemove,
		source: target,
	}
	mod3 := &ModifierInstance{
		name:   modsToRemove,
		source: key.TargetID(3),
	}
	manager.targets[target] = append(manager.targets[target], mod3, mod2, mod1)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
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
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestRemoveModifierWithListener")

	mod := &ModifierInstance{
		name:   name,
		params: make(map[string]float64),
		listeners: Listeners{
			OnRemove: func(modifier *ModifierInstance) {
				modifier.Params()["OnRemoveCalled"] = 1.0
			},
		},
	}

	manager.targets[target] = append(manager.targets[target], mod)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
		assert.Equal(t, name, event.Modifier.Name)
		assert.Contains(t, event.Modifier.Params, "OnRemoveCalled")
		called += 1
	})

	manager.RemoveModifier(target, name)
	assert.Empty(t, manager.targets[target])
	assert.Equal(t, 1, called)
}

func TestRemoveModifierSelf(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestRemoveModifierSelf")

	mod := &ModifierInstance{
		name:   name,
		owner:  target,
		params: make(map[string]float64),
		listeners: Listeners{
			OnRemove: func(modifier *ModifierInstance) {
				modifier.Params()["OnRemoveCalled"] = 1.0
			},
		},
		manager: manager,
	}

	manager.targets[target] = append(manager.targets[target], mod)

	called := 0
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
		assert.Equal(t, name, event.Modifier.Name)
		assert.Contains(t, event.Modifier.Params, "OnRemoveCalled")
		called += 1
	})

	mod.RemoveSelf()
	assert.Empty(t, manager.targets[target])
	assert.Equal(t, 1, called)
}
