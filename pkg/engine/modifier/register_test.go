package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func TestRegisterDefaults(t *testing.T) {
	key := key.Modifier("TestModifierDefaults")
	Register(key, Config{})

	actual := modifierCatalog[key]

	if actual.Duration != -1 {
		t.Errorf("modifier Duration %v does not match expected -1", actual.Duration)
	}

	if actual.Count != 1 {
		t.Errorf("modifier Count %v does not match expected 1", actual.Count)
	}

	if actual.MaxCount != 1 {
		t.Errorf("modifier Count %v does not match expected 1", actual.MaxCount)
	}
}

func TestDuplicateRegistration(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("registration did not panic")
		}
	}()

	key := key.Modifier("TestModifierDuplicate")
	Register(key, Config{})
	Register(key, Config{})
}

func TestRegisterWithListeners(t *testing.T) {
	key := key.Modifier("TestModifierListeners")
	Register(key, Config{
		Listeners: Listeners{
			OnAdd: func(engine engine.Engine, modifier *info.ModifierInstance) {
				modifier.Params["Called"] = 1
			},
		},
	})

	mod := &info.ModifierInstance{
		Params: make(map[string]float64),
	}
	modifierCatalog[key].Listeners.OnAdd(nil, mod)

	if mod.Params["Called"] != 1 {
		t.Errorf("OnAdd registered logic was never called")
	}
}

func TestConfigHasFlag(t *testing.T) {
	key := key.Modifier("TestModifierHasFlag")
	Register(key, Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL},
	})

	config := modifierCatalog[key]
	if !config.HasFlag(model.BehaviorFlag_STAT_CTRL) {
		t.Errorf("config missing flag %v", model.BehaviorFlag_STAT_CTRL)
	}

	if config.HasFlag(model.BehaviorFlag_INVALID_FLAG) {
		t.Errorf("config has flag %v", model.BehaviorFlag_INVALID_FLAG)
	}
}
