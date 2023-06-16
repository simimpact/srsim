package sampo

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) E1TargetDeathListener(e event.TargetDeathEvent) {
	if c.engine.HPRatio(c.id) <= 0 {
		return
	}

	if !c.engine.IsCharacter(e.Killer) {
		return
	}

	if !c.engine.IsEnemy(e.Target) {
		return
	}

	if !c.engine.HasBehaviorFlag(e.Target, model.BehaviorFlag_STAT_DOT_POISON) {
		return
	}

	duration := 1
	for _, target := range c.engine.Enemies() {
		AddWindShearTalent(c.info, c.engine, c.id, target, duration, 1)
	}

}
