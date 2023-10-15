package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Technique key.Attack = "seele-technique"
)

// After using her Technique, Seele gains Stealth for 20 second(s).
// While Stealth is active, Seele cannot be detected by enemies.
// And when entering battle by attacking enemies, Seele will immediately enter the buffed state.

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	// enter buffed state
	c.enterBuffedState()
}
