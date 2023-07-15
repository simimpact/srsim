package lightcone

import (
	"fmt"
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var (
	mu               sync.Mutex
	lightConeCatalog = make(map[key.LightCone]Config)
)

func Register(key key.LightCone, lightCone Config) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := lightConeCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}

	if lightCone.CreatePassive == nil {
		panic("LightCone CreatePassive must be defined: " + key)
	}
	if len(lightCone.Promotions) == 0 {
		panic("LightCone promotions must be defined: " + key)
	}
	if lightCone.Path == model.Path_INVALID_PATH {
		panic("LightCone path must be defined: " + key)
	}

	lightConeCatalog[key] = lightCone
}

type Config struct {
	CreatePassive func(engine engine.Engine, owner key.TargetID, lc info.LightCone)
	Promotions    []PromotionData
	Rarity        int
	Path          model.Path
}

type PromotionData struct {
	MaxLevel int
	ATKBase  float64
	ATKAdd   float64
	HPBase   float64
	HPAdd    float64
	DEFBase  float64
	DEFAdd   float64
	// TODO: add a stat map to each promotion level to simplify implementations?
}

func Get(key key.LightCone) (Config, error) {
	if config, ok := lightConeCatalog[key]; ok {
		return config, nil
	}
	return Config{}, fmt.Errorf("invalid lightcone: %v", key)
}

func (c Config) Ascension(maxLvl, lvl int) int {
	if maxLvl <= 0 {
		maxLvl = lvl
	}

	for i, promo := range c.Promotions {
		if promo.MaxLevel >= maxLvl {
			return i
		}
	}
	return len(c.Promotions) - 1
}

func AddBaseStats(stats info.PropMap, data PromotionData, level int) {
	stats.Modify(prop.ATKBase, data.ATKBase+data.ATKAdd*float64(level-1))
	stats.Modify(prop.DEFBase, data.DEFBase+data.DEFAdd*float64(level-1))
	stats.Modify(prop.HPBase, data.HPBase+data.HPAdd*float64(level-1))
}
