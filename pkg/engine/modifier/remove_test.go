package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
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

	other := &info.ModifierInstance{
		Name: key.Modifier("Other"),
	}
	manager.targets[target] = append(manager.targets[target], other)

	manager.RemoveModifier(target, mod)

	if manager.targets[target][0] != other {
		t.Errorf("RemoveModifier removed the incorrect modifier (was expecting no removal)")
	}
}

func TestRemoveModifierFromSourceNoOp(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	mod := key.Modifier("Test")

	other := &info.ModifierInstance{
		Name:   key.Modifier("Test"),
		Source: key.TargetID(2),
	}
	manager.targets[target] = append(manager.targets[target], other)
	manager.RemoveModifierFromSource(target, target, mod)

	if manager.targets[target][0] != other {
		t.Errorf("RemoveModifier removed the incorrect modifier (was expecting no removal)")
	}
}

func TestRemoveModifier(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	modsToRemove := key.Modifier("ToRemove")

	mod1 := &info.ModifierInstance{
		Name: key.Modifier("Other"),
	}
	mod2 := &info.ModifierInstance{
		Name:   modsToRemove,
		Source: target,
	}
	mod3 := &info.ModifierInstance{
		Name:   modsToRemove,
		Source: key.TargetID(3),
	}
	manager.targets[target] = append(manager.targets[target], mod3, mod1, mod2)
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
		if event.Modifier != mod2 && event.Modifier != mod3 {
			t.Errorf("RemoveModifier removed an unexpected modifier: %v", event.Modifier)
		}
	})

	manager.RemoveModifier(target, modsToRemove)

	if len(manager.targets[target]) != 1 {
		t.Errorf("RemoveModifier did not remove all modifier instances: %v", manager.targets[target])
	}
	if manager.targets[target][0] != mod1 {
		t.Errorf("RemoveModifier removed the incorrect modifier (was expecting no removal)")
	}
}

func TestRemoveModifierFromSource(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	modsToRemove := key.Modifier("ToRemove")

	mod1 := &info.ModifierInstance{
		Name:   modsToRemove,
		Source: key.TargetID(2),
	}
	mod2 := &info.ModifierInstance{
		Name:   modsToRemove,
		Source: target,
	}
	mod3 := &info.ModifierInstance{
		Name:   modsToRemove,
		Source: key.TargetID(3),
	}
	manager.targets[target] = append(manager.targets[target], mod3, mod2, mod1)
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
		if event.Modifier != mod2 {
			t.Errorf("RemoveModifier removed an unexpected modifier: %v", event.Modifier)
		}
	})

	manager.RemoveModifierFromSource(target, target, modsToRemove)

	if len(manager.targets[target]) != 2 {
		t.Errorf("did not remove the correct modifier instances: %v", manager.targets[target])
	}
	if manager.targets[target][0] != mod3 {
		t.Errorf("unknown mod at index 0: expected %v, actual %v", mod3, manager.targets[target][0])
	}
	if manager.targets[target][1] != mod1 {
		t.Errorf("unknown mod at index 1: expected %v, actual %v", mod1, manager.targets[target][1])
	}
}

func TestRemoveModifierWithOnRemoveListener(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestRemoveModifierWithListener")

	Register(name, Config{
		Listeners: Listeners{
			OnRemove: func(engine engine.Engine, modifier *info.ModifierInstance) {
				modifier.Params["OnRemoveCalled"] = 1.0
			},
		},
	})

	// Note: params map is assumed to be made by the Add call
	mod := &info.ModifierInstance{
		Name:   name,
		Params: make(map[string]float64),
	}

	manager.targets[target] = append(manager.targets[target], mod)
	manager.engine.Events().ModifierRemoved.Subscribe(func(event event.ModifierRemovedEvent) {
		if event.Modifier != mod {
			t.Errorf("RemoveModifier removed an unexpected modifier: %v", event.Modifier)
		}

		if event.Modifier.Params["OnRemoveCalled"] != 1.0 {
			t.Errorf("OnRemove was not called")
		}
	})

	manager.RemoveModifier(target, name)

	if len(manager.targets[target]) != 0 {
		t.Errorf("RemoveModifier did not remove all modifier instances: %v", manager.targets[target])
	}
}
