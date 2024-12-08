package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const Technique = "luocha-technique"

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   AbyssFlower,
		Source: c.id,
		Count:  2,
	})
}
