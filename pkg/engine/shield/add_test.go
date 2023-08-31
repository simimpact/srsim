package shield

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/stretchr/testify/assert"

	"github.com/simimpact/srsim/tests/mock"
)

// Unit Tests for AddShield()
func TestShieldHealthByPositiveValues(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)
	sourceStats := mock.NewEmptyStats(source)
	sourceStats.AddProperty("tst", prop.ATKBase, 100.0)
	sourceStats.AddProperty("tst", prop.DEFBase, 80.0)
	sourceStats.AddProperty("tst", prop.HPBase, 60.0)
	shield := &Instance{name: "SourceShield", hp: 20.0}
	manager.targets[source] = append(manager.targets[source], shield)
	attr.EXPECT().Stats(gomock.Eq(source)).Return(sourceStats).Times(5)

	target := key.TargetID(2)
	targetStats := mock.NewEmptyStats(target)
	targetStats.AddProperty("tst", prop.HPBase, 40.0)
	attr.EXPECT().Stats(gomock.Eq(target)).Return(targetStats).Times(5)

	type shieldConfig struct {
		ID            key.Shield
		Info          info.Shield
		ExpectedValue float64
	}

	shieldConfigs := []shieldConfig{
		{
			ID: key.Shield("ShieldBySourceATK"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_ATK: 0.5},
			},
			ExpectedValue: 50.0,
		},
		{
			ID: key.Shield("ShieldBySourceDEF"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.5},
			},
			ExpectedValue: 40.0,
		},
		{
			ID: key.Shield("ShieldBySourceHP"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP: 0.5},
			},
			ExpectedValue: 30.0,
		},
		{
			ID: key.Shield("ShieldByTargetHP"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_TARGET_MAX_HP: 0.5},
			},
			ExpectedValue: 20.0,
		},
		{
			ID: key.Shield("ShieldBySourceShield"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD: 0.5},
			},
			ExpectedValue: 10.0,
		},
	}

	i := 0
	manager.event.ShieldAdded.Subscribe(func(event event.ShieldAdded) {
		assert.Equal(t, shieldConfigs[i].ExpectedValue, event.ShieldHealth)
		i++
	})

	for _, config := range shieldConfigs {
		shieldID := config.ID
		shieldInfo := config.Info
		manager.AddShield(shieldID, shieldInfo)
	}
}

func TestShieldHealthByNegativeValues(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)
	sourceStats := mock.NewEmptyStats(source)
	sourceStats.AddProperty("tst", prop.ATKBase, -100.0)
	sourceStats.AddProperty("tst", prop.DEFBase, -80.0)
	sourceStats.AddProperty("tst", prop.HPBase, -60.0)
	shield := &Instance{name: "SourceShield", hp: -20.0}
	manager.targets[source] = append(manager.targets[source], shield)
	attr.EXPECT().Stats(gomock.Eq(source)).Return(sourceStats).Times(5)

	target := key.TargetID(2)
	targetStats := mock.NewEmptyStats(target)
	targetStats.AddProperty("tst", prop.HPBase, -40.0)
	attr.EXPECT().Stats(gomock.Eq(target)).Return(targetStats).Times(5)

	type shieldConfig struct {
		ID            key.Shield
		Info          info.Shield
		ExpectedValue float64
	}

	shieldConfigs := []shieldConfig{
		{
			ID: key.Shield("ShieldBySourceATK"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_ATK: 0.5},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldBySourceDEF"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.5},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldBySourceHP"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP: 0.5},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldByTargetHP"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_TARGET_MAX_HP: 0.5},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldBySourceShield"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD: 0.5},
			},
			ExpectedValue: 0.0,
		},
	}

	i := 0
	manager.event.ShieldAdded.Subscribe(func(event event.ShieldAdded) {
		assert.Equal(t, shieldConfigs[i].ExpectedValue, event.ShieldHealth)
		i++
	})

	for _, config := range shieldConfigs {
		shieldID := config.ID
		shieldInfo := config.Info
		manager.AddShield(shieldID, shieldInfo)
	}
}

func TestShieldHealthBy0Values(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)
	sourceStats := mock.NewEmptyStats(source)
	sourceStats.AddProperty("tst", prop.ATKBase, 0.0)
	sourceStats.AddProperty("tst", prop.DEFBase, 0.0)
	sourceStats.AddProperty("tst", prop.HPBase, 0.0)
	shield := &Instance{name: "SourceShield", hp: 0.0}
	manager.targets[source] = append(manager.targets[source], shield)
	attr.EXPECT().Stats(gomock.Eq(source)).Return(sourceStats).Times(5)

	target := key.TargetID(2)
	targetStats := mock.NewEmptyStats(target)
	targetStats.AddProperty("tst", prop.HPBase, 0.0)
	attr.EXPECT().Stats(gomock.Eq(target)).Return(targetStats).Times(5)

	type shieldConfig struct {
		ID            key.Shield
		Info          info.Shield
		ExpectedValue float64
	}

	shieldConfigs := []shieldConfig{
		{
			ID: key.Shield("ShieldBySourceATK"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_ATK: 0.5},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldBySourceDEF"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.4},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldBySourceHP"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP: 0.3},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldByTargetHP"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_TARGET_MAX_HP: 0.2},
			},
			ExpectedValue: 0.0,
		},
		{
			ID: key.Shield("ShieldBySourceShield"),
			Info: info.Shield{
				Source:     source,
				Target:     target,
				BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD: 0.1},
			},
			ExpectedValue: 0.0,
		},
	}

	i := 0
	manager.event.ShieldAdded.Subscribe(func(event event.ShieldAdded) {
		assert.Equal(t, shieldConfigs[i].ExpectedValue, event.ShieldHealth)
		i++
	})

	for _, config := range shieldConfigs {
		shieldID := config.ID
		shieldInfo := config.Info
		manager.AddShield(shieldID, shieldInfo)
	}
}

func TestShieldHealthByNoSourceShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)
	sourceStats := mock.NewEmptyStats(source)
	attr.EXPECT().Stats(gomock.Eq(source)).Return(sourceStats).Times(1)

	target := key.TargetID(2)
	targetStats := mock.NewEmptyStats(target)
	attr.EXPECT().Stats(gomock.Eq(target)).Return(targetStats).Times(1)

	shieldID := key.Shield("MockShield")
	shieldInfo := info.Shield{
		Source:     source,
		Target:     target,
		BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD: 0.5},
	}

	manager.event.ShieldAdded.Subscribe(func(event event.ShieldAdded) {
		assert.Equal(t, 0.0, event.ShieldHealth)
	})
	manager.AddShield(shieldID, shieldInfo)
}

// Unit Tests for CheckMatching()
func TestCheckMatchingWhenMatch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)

	target := key.TargetID(2)
	shield := &Instance{name: "MockShield", hp: 10.0}
	manager.targets[target] = append(manager.targets[target], shield)

	shieldID := key.Shield("MockShield")
	shieldInfo := info.Shield{
		Source:     source,
		Target:     target,
		BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.5},
	}

	isTrue, _ := manager.CheckMatching(shieldID, shieldInfo)
	assert.Equal(t, true, isTrue)
}

func TestCheckMatchingWhenNoMatch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)

	target := key.TargetID(2)
	shield := &Instance{name: "NewShield", hp: 10.0}
	manager.targets[target] = append(manager.targets[target], shield)

	shieldID := key.Shield("MockShield")
	shieldInfo := info.Shield{
		Source:     source,
		Target:     target,
		BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.5},
	}

	isTrue, _ := manager.CheckMatching(shieldID, shieldInfo)
	assert.Equal(t, false, isTrue)
}

func TestCheckMatchingWhenNoExistingShield(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	events := &event.System{}
	attr := mock.NewMockAttribute(mockCtrl)
	manager := New(events, attr)

	source := key.TargetID(1)

	target := key.TargetID(2)

	shieldID := key.Shield("MockShield")
	shieldInfo := info.Shield{
		Source:     source,
		Target:     target,
		BaseShield: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.5},
	}

	isTrue, _ := manager.CheckMatching(shieldID, shieldInfo)
	assert.Equal(t, false, isTrue)
}
