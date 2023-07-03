package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
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
	mod := &Instance{
		name:       name,
		statusType: model.StatusType_STATUS_BUFF,
		flags:      []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL},
		stats:      info.PropMap{prop.ATKFlat: 0.1},
		debuffRES:  info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.35},
	}

	manager.targets[target] = append(manager.targets[target], mod)

	result := manager.EvalModifiers(target)

	expectedProps := info.NewPropMap()
	expectedProps.Modify(prop.ATKFlat, 0.1)

	expectedDebuff := info.NewDebuffRESMap()
	expectedDebuff.Modify(model.BehaviorFlag_STAT_CTRL, 0.35)

	expectedCounts := make(map[model.StatusType]int)
	expectedCounts[model.StatusType_STATUS_BUFF] = 1

	assert.Equal(t, expectedProps, result.Props)
	assert.Equal(t, expectedDebuff, result.DebuffRES)
	assert.Equal(t, expectedCounts, result.Counts)
	assert.Equal(t, []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL}, result.Flags)
	assert.Equal(t, []key.Modifier{name}, result.Modifiers)
	assert.Empty(t, result.Weakness)
}

func TestEvalWithMultipleMods(t *testing.T) {
	manager, mockCtrl := NewTestManager(t)
	defer mockCtrl.Finish()
	defer FailOnPanic(t)

	target := key.TargetID(1)

	mod1Name := key.Modifier("TestEvalWithMultipleMods1")
	mod1 := &Instance{
		name:       mod1Name,
		statusType: model.StatusType_STATUS_BUFF,
		stats:      info.PropMap{prop.FireDamageRES: 0.45},
		debuffRES:  info.DebuffRESMap{model.BehaviorFlag_STAT_DOT_BURN: 1.0},
	}

	mod2Name := key.Modifier("TestEvalWithMultipleMods2")
	mod2 := &Instance{
		name:       mod2Name,
		statusType: model.StatusType_STATUS_DEBUFF,
		flags:      []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_STAT_CTRL_STUN},
		stats:      info.PropMap{prop.AllDamageTaken: 0.1},
		debuffRES:  info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: -0.05},
		weakness:   info.WeaknessMap{model.DamageType_ICE: true},
	}

	mod3Name := key.Modifier("TestEvalWithMultipleMods3")
	mod3 := &Instance{
		name:      mod3Name,
		flags:     []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_STAT_ATTACH_WEAKNESS},
		stats:     info.PropMap{},
		debuffRES: info.DebuffRESMap{},
		weakness:  info.WeaknessMap{model.DamageType_QUANTUM: true},
	}

	manager.targets[target] = append(manager.targets[target], mod1, mod2, mod3, mod3)

	result := manager.EvalModifiers(target)

	expectedProps := info.NewPropMap()
	expectedProps.Modify(prop.FireDamageRES, 0.45)
	expectedProps.Modify(prop.AllDamageTaken, 0.1)

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
	assert.Contains(t, result.Weakness, model.DamageType_ICE, model.DamageType_QUANTUM)
}
