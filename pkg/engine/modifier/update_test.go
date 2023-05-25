package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/stretchr/testify/assert"
)

func TestExtendDuration(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestExtendDuration")

	callCount := 0
	listeners := Listeners{
		OnExtendDuration: func(modifier *ModifierInstance) {
			callCount += 1
		},
	}

	mod1 := &ModifierInstance{
		name:      name,
		listeners: listeners,
	}
	mod2 := &ModifierInstance{
		name:      key.Modifier("Other"),
		duration:  2,
		listeners: listeners,
	}
	mod3 := &ModifierInstance{
		name:      name,
		source:    key.TargetID(2),
		duration:  1,
		listeners: listeners,
	}
	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3)

	called := 0
	manager.engine.Events().ModifierExtendedDuration.Subscribe(func(event event.ModifierExtendedDurationEvent) {
		switch called {
		case 0:
			assert.Equal(t, 0, event.OldValue)
			assert.Equal(t, 5, event.NewValue)
		case 1:
			assert.Equal(t, 1, event.OldValue)
			assert.Equal(t, 6, event.NewValue)
		default:
			assert.Fail(t, "unexpected extension call")
		}
		called += 1
	})

	manager.ExtendDuration(target, name, 5)
	assert.Equal(t, 2, called)
	assert.Equal(t, 2, callCount)
}

func TestExtendCount(t *testing.T) {
	manager, mockCtrl := NewTestManagerWithEvents(t)
	defer mockCtrl.Finish()

	target := key.TargetID(1)
	name := key.Modifier("TestExtendCount")

	callCount := 0
	listeners := Listeners{
		OnExtendCount: func(modifier *ModifierInstance) {
			callCount += 1
		},
	}

	mod1 := &ModifierInstance{
		name:      name,
		maxCount:  3,
		listeners: listeners,
	}
	mod2 := &ModifierInstance{
		name:      key.Modifier("Other"),
		count:     2,
		listeners: listeners,
	}
	mod3 := &ModifierInstance{
		name:      name,
		source:    key.TargetID(2),
		count:     1,
		listeners: listeners,
	}
	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3)

	called := 0
	manager.engine.Events().ModifierExtendedCount.Subscribe(func(event event.ModifierExtendedCountEvent) {
		switch called {
		case 0:
			assert.Equal(t, 0.0, event.OldValue)
			assert.Equal(t, 3.0, event.NewValue)
		case 1:
			assert.Equal(t, 1.0, event.OldValue)
			assert.Equal(t, 6.0, event.NewValue)
		default:
			assert.Fail(t, "unexpected extension call")
		}
		called += 1
	})

	manager.ExtendCount(target, name, 5)
	assert.Equal(t, 2, called)
}
