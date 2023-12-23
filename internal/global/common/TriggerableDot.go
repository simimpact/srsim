package common

import "github.com/simimpact/srsim/pkg/engine/modifier"

type TriggerableDot interface {
	//Trigger the corresponding DoT effect to deal damage
	TriggerDot(mod *modifier.Instance, ratio float64)
}
