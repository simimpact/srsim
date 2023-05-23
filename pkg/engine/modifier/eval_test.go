package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestEvalNoMods(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()
	defer FailOnPanic(t)

	result := manager.EvalModifiers(key.TargetID(1))

	assert.Empty(t, result.Props, "Props was not empty")
	assert.Empty(t, result.DebuffRES, "DebuffRES was not empty")
	assert.Empty(t, result.Counts, "Counts was not empty")
	assert.Empty(t, result.Flags, "Flags was not empty")
	assert.Empty(t, result.Modifiers, "Modifiers was not empty")
}

func TestEvalWithMod(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()
	defer FailOnPanic(t)

	target := key.TargetID(1)
	name := key.Modifier("TestEvalWithMod")
	mod := &info.ModifierInstance{
		Name: name,
	}
	mod.Reset(1)
	mod.AddProperty(model.Property_ATK_FLAT, 0.1)
	mod.AddDebuffRES(model.BehaviorFlag_STAT_CTRL, 0.35)

	Register(name, Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL},
		StatusType:    model.StatusType_STATUS_BUFF,
	})

	manager.targets[target] = append(manager.targets[target], mod)

	result := manager.EvalModifiers(target)

	expectedProps := info.NewPropMap()
	expectedProps.Modify(model.Property_ATK_FLAT, 0.1)

	expectedDebuff := info.NewDebuffRESMap()
	expectedDebuff.Modify(model.BehaviorFlag_STAT_CTRL, 0.35)

	expectedCounts := make(map[model.StatusType]int)
	expectedCounts[model.StatusType_STATUS_BUFF] = 1

	assert.Equal(t, expectedProps, result.Props)
	assert.Equal(t, expectedDebuff, result.DebuffRES)
	assert.Equal(t, expectedCounts, result.Counts)
	assert.Equal(t, []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL}, result.Flags)
	assert.Equal(t, []key.Modifier{name}, result.Modifiers)
}

func TestEvalWithMultipleMods(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()
	defer FailOnPanic(t)

	target := key.TargetID(1)

	mod1Name := key.Modifier("TestEvalWithMultipleMods1")
	mod1 := &info.ModifierInstance{
		Name: mod1Name,
	}
	mod1.Reset(1)
	mod1.AddProperty(model.Property_FIRE_DMG_RES, 0.45)
	mod1.AddDebuffRES(model.BehaviorFlag_STAT_DOT_BURN, 1.0)

	Register(mod1Name, Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})

	mod2Name := key.Modifier("TestEvalWithMultipleMods2")
	mod2 := &info.ModifierInstance{
		Name: mod2Name,
	}
	mod2.Reset(1)
	mod2.AddProperty(model.Property_ALL_DMG_TAKEN, 0.1)
	mod2.AddDebuffRES(model.BehaviorFlag_STAT_CTRL, -0.05)

	Register(mod2Name, Config{
		StatusType:    model.StatusType_STATUS_DEBUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_STAT_CTRL_STUN},
	})

	mod3Name := key.Modifier("TestEvalWithMultipleMods3")
	mod3 := &info.ModifierInstance{
		Name: mod3Name,
	}
	mod3.Reset(1)

	Register(mod3Name, Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_STAT_ATTACH_WEAKNESS},
	})

	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3, mod3)

	result := manager.EvalModifiers(target)

	expectedProps := info.NewPropMap()
	expectedProps.Modify(model.Property_FIRE_DMG_RES, 0.45)
	expectedProps.Modify(model.Property_ALL_DMG_TAKEN, 0.1)

	expectedDebuff := info.NewDebuffRESMap()
	expectedDebuff.Modify(model.BehaviorFlag_STAT_CTRL, -0.05)
	expectedDebuff.Modify(model.BehaviorFlag_STAT_DOT_BURN, 1.0)

	expectedCounts := map[model.StatusType]int{
		model.StatusType_UNKNOWN_STATUS: 2,
		model.StatusType_STATUS_BUFF:    1,
		model.StatusType_STATUS_DEBUFF:  1,
	}

	expectedFlags := []model.BehaviorFlag{
		model.BehaviorFlag_STAT_CTRL,
		model.BehaviorFlag_STAT_CTRL_STUN,
		model.BehaviorFlag_STAT_ATTACH_WEAKNESS,
	}

	assert.Equal(t, expectedProps, result.Props)
	assert.Equal(t, expectedDebuff, result.DebuffRES)
	assert.Equal(t, expectedCounts, result.Counts)
	assert.ElementsMatch(t, expectedFlags, result.Flags)
	assert.ElementsMatch(t, []key.Modifier{mod1Name, mod2Name, mod3Name}, result.Modifiers)
}
