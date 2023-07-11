package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	A4 key.Modifier = "natasha-a4"
)

func (c *char) initTraces() {
	// A4
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4,
			Source: c.id,
			Stats:  info.PropMap{prop.HealBoost: 0.1},
		})
	}
}
