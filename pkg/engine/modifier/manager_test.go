package modifier_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/mock"
	"github.com/simimpact/srsim/pkg/model"
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
			attr := info.Attributes{}
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
			OnPropertyChange: func(mod *modifier.ModifierInstance) {
				stats := mod.Engine().Stats(mod.Owner())
				if stats.GetProperty(model.Property_DEF_PERCENT) >= 0.1 {
					mod.SetProperty(model.Property_ALL_DMG_PERCENT, 0.1)
				} else {
					mod.SetProperty(model.Property_ALL_DMG_PERCENT, 0.0)
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
		Stats:           info.PropMap{model.Property_DEF_PERCENT: 0.15},
	})

	mods = manager.EvalModifiers(target)
	expectedProps = info.PropMap{
		model.Property_DEF_PERCENT:     0.15,
		model.Property_ALL_DMG_PERCENT: 0.1,
	}
	assert.Equal(t, expectedProps, mods.Props)

	manager.Tick(target, modifier.ModifierPhase2End)
	mods = manager.EvalModifiers(target)
	assert.Empty(t, mods.Props, "all modifiers were not removed")
}
