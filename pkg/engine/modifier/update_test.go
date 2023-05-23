package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func TestExtendDuration(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestExtendDuration")

	mod1 := &info.ModifierInstance{
		Name: name,
	}
	mod2 := &info.ModifierInstance{
		Name:     key.Modifier("Other"),
		Duration: 2,
	}
	mod3 := &info.ModifierInstance{
		Name:     name,
		Source:   key.TargetID(2),
		Duration: 1,
	}
	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3)

	callCount := 0
	Register(name, Config{
		Listeners: Listeners{
			OnExtendDuration: func(engine engine.Engine, modifier *info.ModifierInstance) {
				callCount += 1
			},
		},
	})

	manager.engine.Events().ModifierExtended.Subscribe(func(event event.ModifierExtendedEvent) {
		if event.Modifier != mod1 && event.Modifier != mod3 {
			t.Errorf("unknown modifier was extended: %v", event.Modifier)
		}

		if event.Operation != "ExtendDuration" {
			t.Errorf("unknown operation: %v", event)
		}

		if event.Modifier == mod1 {
			if event.OldValue != 0 || event.NewValue != 5 {
				t.Errorf("event old and new values do not match the expected 0 and 5: %v", event)
			}
		}

		if event.Modifier == mod3 {
			if event.OldValue != 1 || event.NewValue != 6 {
				t.Errorf("event old and new values do not match the expected 1 and 6: %v", event)
			}
		}
	})

	manager.ExtendDuration(target, name, 5)

	if callCount != 2 {
		t.Errorf("OnExtendedDuration listener was not called twice: %v", callCount)
	}
}

func TestExtendCount(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestExtendCount")

	mod1 := &info.ModifierInstance{
		Name:     name,
		MaxCount: 3,
	}
	mod2 := &info.ModifierInstance{
		Name:  key.Modifier("Other"),
		Count: 2,
	}
	mod3 := &info.ModifierInstance{
		Name:   name,
		Source: key.TargetID(2),
		Count:  1,
	}
	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3)

	callCount := 0
	Register(name, Config{
		Listeners: Listeners{
			OnExtendCount: func(engine engine.Engine, modifier *info.ModifierInstance) {
				callCount += 1
			},
		},
	})

	manager.engine.Events().ModifierExtended.Subscribe(func(event event.ModifierExtendedEvent) {
		if event.Modifier != mod1 && event.Modifier != mod3 {
			t.Errorf("unknown modifier was extended: %v", event.Modifier)
		}

		if event.Operation != "ExtendCount" {
			t.Errorf("unknown operation: %v", event)
		}

		if event.Modifier == mod1 {
			if event.OldValue != 0 || event.NewValue != 3 {
				t.Errorf("event old and new values do not match the expected 0 and 3: %v", event)
			}
		}

		if event.Modifier == mod3 {
			if event.OldValue != 1 || event.NewValue != 6 {
				t.Errorf("event old and new values do not match the expected 1 and 6: %v", event)
			}
		}
	})

	manager.ExtendCount(target, name, 5)

	if callCount != 2 {
		t.Errorf("OnExtendedDuration listener was not called twice: %v", callCount)
	}
}
