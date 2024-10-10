package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Technique key.Modifier = "trailblazerimaginary-technique"

func init() {
	modifier.Register(Technique, modifier.Config{
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, target := range c.engine.Characters() {
		c.engine.AddModifier(target, info.Modifier{
			Name:   Technique,
			Source: c.id,
			Stats: info.PropMap{
				prop.BreakEffect: 0.3,
			},
		})
	}
}
