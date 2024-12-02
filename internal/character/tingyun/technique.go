package tingyun

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const Technique key.Reason = "tingyun-technique"

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	// Technique can be used multiple times in succession to gain the Energy the same amount of times
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Technique,
		Target: c.id,
		Source: c.id,
		Amount: 50.0,
	})
}
