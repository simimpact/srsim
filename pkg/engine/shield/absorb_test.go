package shield

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/stretchr/testify/assert"

	"github.com/simimpact/srsim/tests/mock"
)

func TestAbsorbDamageWhenNoShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	damageInitial := 100.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Fail(t, "There should be an early exit condition for 0 damage and/or no shield, event should not emit")
	})

	damage := manager.AbsorbDamage(target, damageInitial)
	assert.Equal(t, damageInitial, damage)
}

func TestAbsorbDamageWhen0Damage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 20.0}
	manager.targets[target] = append(manager.targets[target], shield)

	damageInitial := 0.

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Fail(t, "There should be an early exit condition for 0 damage and/or no shield, event should not emit")
	})

	damage := manager.AbsorbDamage(target, damageInitial)
	assert.Equal(t, damageInitial, damage)
}

func TestAbsorbDamageWhenNegativeDamage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 20.0}
	manager.targets[target] = append(manager.targets[target], shield)

	damageInitial := -10.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Fail(t, "There should be an early exit condition for 0 damage and/or no shield, event should not emit")
	})

	damage := manager.AbsorbDamage(target, damageInitial)
	assert.Equal(t, damageInitial, damage)
}

func TestAbsorbDamageWhen0DamageAndNoShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	damageInitial := 0.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Fail(t, "There should be an early exit condition for 0 damage and/or no shield, event should not emit")
	})

	damage := manager.AbsorbDamage(target, damageInitial)
	assert.Equal(t, damageInitial, damage)
}

func TestAbsorbDamageWhenNegativeDamageAndNoShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)

	damageInitial := -10.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Fail(t, "There should be an early exit condition for 0 damage and/or no shield, event should not emit")
	})

	damage := manager.AbsorbDamage(target, damageInitial)
	assert.Equal(t, damageInitial, damage)
}

func TestAbsorbDamageWithSingleShieldWhenDamageLessThanShieldHealth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 100.0}
	manager.targets[target] = append(manager.targets[target], shield)

	damageInitial := 10.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Equal(t, target, event.Target)
		assert.Equal(t, shield.name, event.ID)
		assert.Equal(t, 90.0, event.NewHP)
		assert.Equal(t, 100.0, event.OldHP)
		assert.Equal(t, damageInitial, event.DamageIn)
		assert.Equal(t, 0.0, event.DamageOut)
	})
	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Fail(t, "Damage was less than shield health, therefore shield should not be removed")
	})

	damage := manager.AbsorbDamage(target, damageInitial)

	assert.Equal(t, 0.0, damage)
}

func TestAbsorbDamageWithSingleShieldWhenDamageGreaterThanShieldHealth(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shield := &Instance{name: "Shield", hp: 90.0}
	manager.targets[target] = append(manager.targets[target], shield)

	damageInitial := 110.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Equal(t, target, event.Target)
		assert.Equal(t, key.Shield(""), event.ID)
		assert.Equal(t, 0.0, event.NewHP)
		assert.Equal(t, 90.0, event.OldHP)
		assert.Equal(t, damageInitial, event.DamageIn)
		assert.Equal(t, 20.0, event.DamageOut)
	})
	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Equal(t, shield.name, event.ID)
		assert.Equal(t, 0, len(manager.targets[target]))
	})

	damage := manager.AbsorbDamage(target, damageInitial)

	assert.Equal(t, 20.0, damage)
}

func TestAbsorbDamageWithTwoShieldsOfHPBelowAndAboveIncomingDamage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	shieldremoved := &Instance{name: "ShieldRemoved", hp: 100.0}
	shieldkept := &Instance{name: "ShieldKept", hp: 120.0}
	manager.targets[target] = append(manager.targets[target], shieldremoved, shieldkept)

	damageInitial := 110.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Equal(t, target, event.Target)
		assert.Equal(t, shieldkept.name, event.ID)
		assert.Equal(t, 10.0, event.NewHP)
		assert.Equal(t, 120.0, event.OldHP)
		assert.Equal(t, damageInitial, event.DamageIn)
		assert.Equal(t, 0.0, event.DamageOut)
	})
	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Equal(t, shieldremoved.name, event.ID)
		assert.Equal(t, 1, len(manager.targets[target]))
	})

	damage := manager.AbsorbDamage(target, damageInitial)

	assert.Equal(t, 0.0, damage)
}

func TestAbsorbDamageWithTwoShieldsOfHPAboveIncomingDamage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	maxshield := &Instance{name: "MaxShield", hp: 130.0}
	shield := &Instance{name: "Shield", hp: 120.0}
	manager.targets[target] = append(manager.targets[target], maxshield, shield)

	damageInitial := 110.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Equal(t, target, event.Target)
		assert.Equal(t, maxshield.name, event.ID)
		assert.Equal(t, 20.0, event.NewHP)
		assert.Equal(t, 130.0, event.OldHP)
		assert.Equal(t, damageInitial, event.DamageIn)
		assert.Equal(t, 0.0, event.DamageOut)
		assert.Equal(t, 2, len(manager.targets[target]))
	})
	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Fail(t, "No shields hp were below the incoming damage so this should never occur")
	})

	damage := manager.AbsorbDamage(target, damageInitial)

	assert.Equal(t, 0.0, damage)
}

func TestAbsorbDamageWithTwoShieldsOfHPBelowIncomingDamage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	target := key.TargetID(1)
	maxshield := &Instance{name: "MaxShield", hp: 20.0}
	shield := &Instance{name: "Shield", hp: 10.0}
	manager.targets[target] = append(manager.targets[target], maxshield, shield)

	damageInitial := 110.0

	manager.event.ShieldChange.Subscribe(func(event event.ShieldChange) {
		assert.Equal(t, target, event.Target)
		assert.Equal(t, key.Shield(""), event.ID)
		assert.Equal(t, 0.0, event.NewHP)
		assert.Equal(t, 20.0, event.OldHP)
		assert.Equal(t, damageInitial, event.DamageIn)
		assert.Equal(t, 90.0, event.DamageOut)
	})
	manager.event.ShieldRemoved.Subscribe(func(event event.ShieldRemoved) {
		assert.Equal(t, 0, len(manager.targets[target]))
	})

	damage := manager.AbsorbDamage(target, damageInitial)

	assert.Equal(t, 90.0, damage)
}
