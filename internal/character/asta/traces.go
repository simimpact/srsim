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
	modifier.Register(a4, modifier.Config{
		Listeners: modifier.Listeners{
			OnRemove: a4OnDestroy,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["102"] {
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
}

func a4OnDestroy(mod *modifier.Instance) {
	for _, ally := range mod.Engine().Characters() {
		mod.Engine().RemoveModifier(ally, a4)
	}
}
