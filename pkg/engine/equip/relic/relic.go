package relic

import (
	"fmt"
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

var (
	mu           sync.Mutex
	relicCatalog = make(map[key.Relic]Config)
)

func Register(key key.Relic, relic Config) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := relicCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}

	// Ensure Stats is never nil
	for i := 0; i < len(relic.Effects); i++ {
		if relic.Effects[i].Stats == nil {
			relic.Effects[i].Stats = info.NewPropMap()
		}
	}

	relicCatalog[key] = relic
}

func Get(key key.Relic) (Config, error) {
	if config, ok := relicCatalog[key]; ok {
		return config, nil
	}
	return Config{}, fmt.Errorf("invalid relic: %v", key)
}

type Config struct {
	Effects []SetEffect
	// TODO: RelicType? (planar vs cavern)
}

type CreateEffectFunc func(engine engine.Engine, owner key.TargetID)

type SetEffect struct {
	MinCount     int
	Stats        info.PropMap
	CreateEffect CreateEffectFunc
}
