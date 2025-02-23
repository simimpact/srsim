package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	technique = "sparkle-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.ModifySP(info.ModifySP{
		Source: c.id,
		Key:    technique,
		Amount: 3,
	})
}
