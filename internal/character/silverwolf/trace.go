package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/event"
)

func (c *char) initTraces() {
	// A2
	// 	Bug's duration is extended for 1 turn(s). Every time an enemy is inflicted
	// 	with Weakness Break, Silver Wolf has a 65% base chance of implanting a
	// 	random Bug in the enemy.
	if c.info.Traces["101"] {
		c.engine.Events().StanceBreak.Subscribe(func(e event.StanceBreak) {
			mod := newRandomBug(c.engine, e.Target, c.id)
			mod.Chance = 0.65
			c.engine.AddModifier(e.Target, mod)
		})
	}
}
