package sampo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, trg := range c.engine.Enemies() {
		c.engine.ModifyGaugeNormalized(trg, 0.25)
	}
}
