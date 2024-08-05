package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	technique = "himeko-technique"
)

func init() {
	modifier.Register(technique, modifier.Config{
		Duration: 2,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, e := range c.engine.Enemies() {
		c.engine.AddModifier(e, info.Modifier{
			Name:     technique,
			Source:   c.id,
			Chance:   1,
			Duration: 2,
			Stats: info.PropMap{
				prop.FireDamageTaken: 0.1,
			},
		})
	}
}
