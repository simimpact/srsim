package lightcone

import (
	"fmt"
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
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

func (c Config) Ascension(maxLvl int) int {
	for i, promo := range c.Promotions {
		if promo.MaxLevel >= maxLvl {
			return i
		}
	}
	return len(c.Promotions) - 1
}

func AddBaseStats(stats info.PropMap, data PromotionData, level int) {
	stats.Modify(model.Property_ATK_BASE, data.ATKBase+data.ATKAdd*float64(level-1))
	stats.Modify(model.Property_DEF_BASE, data.DEFBase+data.DEFAdd*float64(level-1))
	stats.Modify(model.Property_HP_BASE, data.HPBase+data.HPAdd*float64(level-1))
}
