package common

import "github.com/simimpact/srsim/pkg/engine/modifier"

type TriggerableDot interface {
	TriggerDot(mod *modifier.Instance, ratio float64)
}
