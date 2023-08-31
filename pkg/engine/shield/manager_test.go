package shield

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/stretchr/testify/assert"

	"github.com/simimpact/srsim/tests/mock"
)

// Unit Tests for HasShield()
func TestHasShieldWhenShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 0.0}
	manager.targets[target] = append(manager.targets[target], shield)

	isTrue := manager.HasShield(target, shield.name)
	assert.Equal(t, true, isTrue)
}

func TestHasShieldWhenNoShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	isTrue := manager.HasShield(target, key.Shield("Shield"))
	assert.Equal(t, false, isTrue)
}

// Unit Tests for IsShielded()
func TestIsShieldedWhenShielded(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 0.0} // Current behavior still counts a shield with 0 hp on a target as shielded
	manager.targets[target] = append(manager.targets[target], shield)

	isTrue := manager.IsShielded(target)
	assert.Equal(t, true, isTrue)
}

func TestIsShieldedWhenNotShielded(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	isTrue := manager.IsShielded(target)
	assert.Equal(t, false, isTrue)
}

// Unit Tests for MaxShield()
func TestMaxShieldWhenPositiveShieldHP(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 100.0}
	manager.targets[target] = append(manager.targets[target], shield)

	hp := manager.MaxShield(target)
	assert.Equal(t, 100.0, hp)
}

func TestMaxShieldWhenNegativeShieldHP(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: -100.0} // With current implementation, this becomes 0.0
	manager.targets[target] = append(manager.targets[target], shield)

	hp := manager.MaxShield(target)
	assert.Equal(t, 0.0, hp)
}

func TestMaxShieldWhen0ShieldHP(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 0.0}
	manager.targets[target] = append(manager.targets[target], shield)

	hp := manager.MaxShield(target)
	assert.Equal(t, 0.0, hp)
}

func TestMaxShieldWhenNotShielded(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	hp := manager.MaxShield(target)
	assert.Equal(t, 0.0, hp)
}
