package attribute

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type attrTarget struct {
	attributes   *info.Attributes
	lastAttacker key.TargetID
	state        info.TargetState
}
