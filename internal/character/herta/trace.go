package herta

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	a4 = "herta-a4"
	a6 = "herta-a6"
)

func init() {
	modifier.Register(a4, modifier.Config{})

	modifier.Register(a6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingHit: a6BeforeBeingHit,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   a4,
			Source: c.id,
			DebuffRES: info.DebuffRESMap{
				model.BehaviorFlag_STAT_CTRL: 0.35,
			},
		},
		)
	}
}

func a6BeforeBeingHit(mod *modifier.Instance, e event.HitStart) {
	e.Hit.Attacker.AddProperty(a6, prop.AllDamagePercent, 0.2)
}
