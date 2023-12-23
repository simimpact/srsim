package welt

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// After using Welt's Technique, create a dimension that lasts for 15 second(s).
// Enemies in this dimension have their Movement SPD reduced by 50%.
// After entering battle with enemies in the dimension,
// there is a 100% base chance to Imprison the enemies for 1 turn.
// Imprisoned enemies have their actions delayed by 20% and SPD reduced by 10%.
// Only 1 dimension created by allies can exist at the same time.

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	// global AOE imprison : use common.Imprisonment again
	for _, target := range c.engine.Enemies() {
		c.engine.AddModifier(target, info.Modifier{
			Name:   common.Imprisonment,
			Source: c.id,
			State: common.ImprisonState{
				DelayRatio:     0.2,
				SpeedDownRatio: 0.1,
			},
			Chance:          1,
			Duration:        1,
			TickImmediately: true,
		})
	}
}
