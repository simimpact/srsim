package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/event"
)

// A2:
// 	Bug's duration is extended for 1 turn(s). Every time an enemy is inflicted with Weakness Break, Silver Wolf has a 65% base chance of implanting a random Bug in the enemy.
// A4:
//	The duration of the Weakness implanted by Silver Wolf's Skill increases by 1 turn(s).
// A6:
//	If there are 3 or more debuff(s) affecting the enemy when the Skill is used, then the Skill decreases the enemy's All-Type RES by an additional 3%.

func (c *char) initTraces() {
	// A2
	if c.info.Traces["1006101"] {
		c.engine.Events().StanceBreak.Subscribe(func(e event.StanceBreakEvent) {
			mod := newRandomBug(c.engine, e.Target, c.id)
			mod.Chance = 0.65
			c.engine.AddModifier(e.Target, mod)
		})
	}
}
