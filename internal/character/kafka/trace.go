package kafka

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4 = "kafka-a4"
)

func (c *char) initTraces() {
	if c.info.Traces["102"] {
		c.engine.Events().TargetDeath.Subscribe(c.a4Listener)
	}
}

func (c *char) a4Listener(e event.TargetDeath) {
	if c.engine.HasBehaviorFlag(e.Target, model.BehaviorFlag_STAT_DOT_ELECTRIC) {
		if c.engine.IsEnemy(e.Target) {
			c.engine.ModifyEnergy(info.ModifyAttribute{
				Key:    A4,
				Target: c.id,
				Source: c.id,
				Amount: 5,
			})
		}
	}
}
