package neutraltarget

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Config struct {
	Create  func(engine engine.Engine, id key.TargetID, info info.NeutralTarget) info.NeutralTargetInstance
	Attack  Attack
	Speed   float64
	Element model.DamageType
}

type Attack struct {
	TargetType model.TargetType
}
