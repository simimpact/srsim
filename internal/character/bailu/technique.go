package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// After Technique is used, at the start of the next battle,
// all allies are granted Invigoration for 2 turn(s).

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, char := range c.engine.Characters() {
		c.addInvigoration(char, 2)
	}
}
