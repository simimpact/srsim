package hook

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4 key.Modifier = "hook-a4"
)

func (c *char) initTraces() {
	// A4
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4,
			Source: c.id,
			DebuffRES: info.DebuffRESMap{
				model.BehaviorFlag_STAT_CTRL: 0.35,
			},
		})
	}
}
