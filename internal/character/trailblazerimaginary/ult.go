package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const Ult key.Reason = "trailblazerimaginary-ult"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.ultLifeTime = 3
	for _, target := range c.engine.Characters() {
		c.engine.AddModifier(target, info.Modifier{
			Name:   BackupDancer,
			Source: c.id,
			Stats: info.PropMap{
				prop.BreakEffect: ultBreakEffect[c.info.UltLevelIndex()],
			},
		})
	}
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Ult,
		Source: c.id,
		Target: c.id,
		Amount: 5,
	})
}
