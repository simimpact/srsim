package hook

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E6 = "hook-e6"
)

func init() {
	modifier.Register(E6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: func(mod *modifier.Instance, e event.HitStart) {
				if mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_DOT_BURN) {
					e.Hit.Attacker.AddProperty(E6, prop.AllDamagePercent, 0.2)
				}
			},
		},
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}
