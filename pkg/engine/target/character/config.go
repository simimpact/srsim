package character

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Config struct {
	Create     func(engine engine.Engine, id key.TargetID, info info.Character) CharInstance
	Promotions []PromotionData
	Rarity     int
	Element    model.DamageType
	Path       model.Path
	MaxEnergy  float64
	Traces     TraceMap
	// TODO:
	//	- ability metadata
	//	- body
}

type TraceMap map[string]Trace

type Trace struct {
	Stat      model.Property
	Amount    float64
	Ascension int
	Level     int
}

type PromotionData struct {
	MaxLevel   int
	ATKBase    float64
	ATKAdd     float64
	DEFBase    float64
	DEFAdd     float64
	HPBase     float64
	HPAdd      float64
	SPD        float64
	CritChance float64
	CritDMG    float64
	Aggro      float64
}

func (c Config) ascension(level int) int {
	for i, promo := range c.Promotions {
		if promo.MaxLevel >= level {
			return i
		}
	}
	return len(c.Promotions) - 1
}
