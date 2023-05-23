package modifier

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestDuplicateRegistration(t *testing.T) {
	key := key.Modifier("TestModifierDuplicate")
	Register(key, Config{})
	assert.Panics(t, func() { Register(key, Config{}) })
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
	assert.Equal(t, 1.0, mod.Params["Called"])
}

func TestConfigHasFlag(t *testing.T) {
	key := key.Modifier("TestModifierHasFlag")
	Register(key, Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL},
	})

	config := modifierCatalog[key]
	assert.True(t, config.HasFlag(model.BehaviorFlag_STAT_CTRL), "config missing STAT_CTRL flag")
	assert.False(t, config.HasFlag(model.BehaviorFlag_INVALID_FLAG), "config has INVALID_FLAG")
}
