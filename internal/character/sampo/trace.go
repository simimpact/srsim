package sampo

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A6 key.Modifier = "Sampo_A6"
)

func init() {
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll: A6OnBeforeBeingHitAll,
		},
	})
}

func (c *char) a4() {
	if c.info.Traces["102"] {
		c.engine.ModifyEnergy(c.id, 10)
	}
}

func A6OnBeforeBeingHitAll(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HasBehaviorFlag(e.Hit.Attacker.ID(), model.BehaviorFlag_STAT_DOT_POISON) {
		e.Hit.Attacker.AddProperty(prop.Fatigue, 0.15)
	}
}

func (c *char) initTraces() {
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}
