package hanya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	victim := c.engine.Retarget(info.Retarget{
		Targets: c.engine.Enemies(),
		Filter: func(target key.TargetID) bool {
			return c.engine.HPRatio(target) > 0
		},
		Max:          1,
		IncludeLimbo: false,
	})[0]

	if c.engine.HPRatio(victim) > 0 {
		c.engine.AddModifier(victim, info.Modifier{
			Name:   Burden,
			Source: c.id,
			State: BurdenState{
				atkCount:          0,
				triggersRemaining: 2,
			},
		})
	}
}
