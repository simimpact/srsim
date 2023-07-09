package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type attrTarget struct {
	attributes   *info.Attributes
	lastAttacker key.TargetID
	state        TargetState
}

type TargetState int

const (
	Invalid TargetState = 0
	Dead    TargetState = 1
	Limbo   TargetState = 2
	Alive   TargetState = 3
)
