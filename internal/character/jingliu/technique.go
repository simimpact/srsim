package jingliu

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Tech key.Reason = "jingliu-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, Target := range c.engine.Enemies() {
		c.engine.AddModifier(Target, info.Modifier{
			Name:   common.Freeze,
			Source: c.id,
			State: &common.FreezeState{
				DamagePercentage: 0.8,
				DamageValue:      0,
			},
			Chance:   1,
			Duration: 1,
		})
	}
	c.gainSyzygy()
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Tech,
		Target: c.id,
		Source: c.id,
		Amount: 15,
	})
}
