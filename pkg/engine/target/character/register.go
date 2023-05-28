package character

import (
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var (
	mu               sync.Mutex
	characterCatalog = make(map[key.Character]Config)
)

func Register(key key.Character, character Config) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := characterCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}
	characterCatalog[key] = character
}

type Config struct {
	Create     func(engine engine.Engine, id key.TargetID, info info.Character) CharInstance
	Promotions []PromotionData
	Rarity     int
	Element    model.DamageType
	Path       model.Path
	MaxEnergy  float64
	// TODO:
	// 	- list of all valid traces?
	//	- ability metadata
	//	- body
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

func (c Config) fromPromotionData(asc int, level int) (info.PropMap, int) {
	if asc < 0 {
		asc = 0
	}
	if asc >= len(c.Promotions) {
		asc = len(c.Promotions) - 1
	}
	data := c.Promotions[asc]

	out := info.NewPropMap()
	out.Modify(model.Property_ATK_BASE, data.ATKBase+data.ATKAdd*float64(level-1))
	out.Modify(model.Property_DEF_BASE, data.DEFBase+data.DEFAdd*float64(level-1))
	out.Modify(model.Property_HP_BASE, data.HPBase+data.HPAdd*float64(level-1))
	out.Modify(model.Property_SPD_BASE, data.SPD)
	out.Modify(model.Property_CRIT_CHANCE, data.CritChance)
	out.Modify(model.Property_CRIT_DMG, data.CritDMG)
	out.Modify(model.Property_AGGRO_BASE, data.Aggro)
	return out, data.MaxLevel
}
