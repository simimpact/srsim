package enemy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type Config struct {
	Create func(engine engine.Engine, id key.TargetID) info.EnemyInstance
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
