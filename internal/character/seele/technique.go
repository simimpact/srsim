package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Technique key.Attack = "seele-technique"
)

// After using her Technique, Seele gains Stealth for 20 second(s).
// While Stealth is active, Seele cannot be detected by enemies.
// And when entering battle by attacking enemies, Seele will immediately enter the buffed state.

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	resPenAmt := 0.0
	if c.info.Traces["102"] {
		resPenAmt = 0.2
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   BuffedState,
		Source: c.id,
		Stats: info.PropMap{
			prop.AllDamagePercent: talent[c.info.TalentLevelIndex()],
			prop.QuantumPEN:       resPenAmt,
		},
	})
}
