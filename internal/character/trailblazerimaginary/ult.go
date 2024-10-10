package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const Ult key.Reason = "trailblazerimaginary-ult"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   BackupDancerCountdown,
		Source: c.id,
	})
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Ult,
		Source: c.id,
		Target: c.id,
		Amount: 5,
	})
}
