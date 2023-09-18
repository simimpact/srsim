package enemy

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Config struct {
	Create func(engine engine.Engine, id key.TargetID, info info.Enemy) info.EnemyInstance
	Curve  LevelCurve
	Rank   model.EnemyRank
	Base   BaseStats
}

type BaseStats struct {
	ATK        float64
	DEF        float64
	HP         float64
	SPD        float64
	Stance     float64
	EffectRES  float64
	CritChance float64
	CritDMG    float64
	MinFatigue float64
}

type LevelData struct {
	ATKScaling    float64
	DEFScaling    float64
	HPScaling     float64
	SPDScaling    float64
	StanceScaling float64
	EffectHitRate float64
	EffectRES     float64
}

type LevelCurve int

const (
	Curve1 LevelCurve = 1
	Curve2 LevelCurve = 2
)

func Curve(c LevelCurve) []LevelData {
	switch c {
	case Curve1:
		return LevelCurve1
	case Curve2:
		return LevelCurve2
	default:
		panic(fmt.Sprintf("unknown level curve %d", c))
	}
}
