package march7th

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	freezeTarget := c.engine.Retarget(info.Retarget{
		Targets: c.engine.Enemies(),
		Filter: func(target key.TargetID) bool {
			return c.engine.Stats(target).GetDebuffRES(model.BehaviorFlag_STAT_CTRL_FROZEN, model.BehaviorFlag_STAT_CTRL) <= 0.8
		},
		Max: 1,
	})

	c.engine.AddModifier(freezeTarget[0], info.Modifier{
		Name:   common.Freeze,
		Source: c.id,
		Chance: 1,
		State: &common.FreezeState{
			DamagePercentage: 0.5,
			DamageValue:      0,
		},
		Duration: 1,
	})
}
