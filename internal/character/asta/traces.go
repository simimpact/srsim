package asta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	a4 = "asta-a4"
)

func init() {
	modifier.Register(a4, modifier.Config{})
}

func (c *char) initTraces() {
	for _, ally := range c.engine.Characters() {
		c.engine.AddModifier(ally, info.Modifier{
			Name:   a4,
			Source: c.id,
			Stats: info.PropMap{
				prop.FireDamagePercent: 0.18,
			},
		})
	}
}
