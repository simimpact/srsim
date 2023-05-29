package relic

import (
	"fmt"
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
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
	relicCatalog[key] = relic
}

type Config struct {
	CreateSet func(engine engine.Engine, owner key.TargetID, count int)
	// TODO: RelicType? (planar vs cavern)
}

func Get(key key.Relic) (Config, error) {
	if config, ok := relicCatalog[key]; ok {
		return config, nil
	}
	return Config{}, fmt.Errorf("invalid relic: %v", key)
}
