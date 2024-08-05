package common

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type TriggerableDot interface {
	// Trigger the corresponding DoT effect to deal damage
	TriggerDot(mod info.Modifier, ratio float64, engine engine.Engine, target key.TargetID)
}
