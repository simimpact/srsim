package guinaifen

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4 = "guinaifen-a4"
	A6 = "guinaifen-a6"
)

func init() {
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: addA6,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["102"] {
		c.engine.Events().BattleStart.Subscribe(func(e event.BattleStart) {
			c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
				Key:    A4,
				Target: c.id,
				Source: c.id,
				Amount: -0.25,
			})
		})
	}

	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}

func addA6(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_DOT_BURN) {
		e.Hit.Attacker.AddProperty(A6, prop.AllDamagePercent, 0.2)
	}
}
