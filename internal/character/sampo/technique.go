package sampo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const Tehnique key.Reason = "sampo-technique"

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, trg := range c.engine.Enemies() {
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    Tehnique,
			Target: trg,
			Source: c.id,
			Amount: 0.25,
		})
	}
}
