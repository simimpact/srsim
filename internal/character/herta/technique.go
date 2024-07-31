package herta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	technique = "herta-technique"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   technique,
		Source: c.id,
		Stats: info.PropMap{
			prop.ATKPercent: 0.4,
		},
		Duration: 3,
	})
}
