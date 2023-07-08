package clara

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Technique key.Modifier = "clara-technique-aggro"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     Technique,
		Source:   c.id,
		Stats:    info.PropMap{prop.AggroPercent: 5},
		Duration: 2,
	})
}
