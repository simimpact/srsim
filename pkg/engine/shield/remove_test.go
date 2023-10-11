package shield

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/stretchr/testify/assert"

	"github.com/simimpact/srsim/tests/mock"
)

func TestRemoveShieldWhenHasShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	otherShield := &Instance{name: "Shield", hp: 0.0}
	shieldToBeRemoved := &Instance{name: "ShieldToBeRemoved", hp: 0.0}
	manager.targets[target] = append(manager.targets[target], otherShield, shieldToBeRemoved)

	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Equal(t, key.Shield("ShieldToBeRemoved"), event.ID)
		assert.Equal(t, 1, len(manager.targets[target]))
		assert.Equal(t, otherShield.name, manager.targets[target][0].name)
	})

	manager.RemoveShield(shieldToBeRemoved.name, target)
}

func TestRemoveShieldWhenNoShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Fail(t, "This emit should never occur, ensure that the early exit condition for shield not found is still in the function")
	})

	manager.RemoveShield(key.Shield("Shield"), target)
}
