package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// attempt to kill off any dead targets and remove them from the battlefield
func (sim *Simulation) deathCheck(killLimbo bool) {
	charIdx := 0
	for _, target := range sim.characters {
		if sim.kill(target, killLimbo) {
			sim.deathEvent(target)
		} else {
			sim.characters[charIdx] = target
			charIdx++
		}
	}
	sim.characters = sim.characters[:charIdx]

	enemyIdx := 0
	for _, target := range sim.enemies {
		if sim.kill(target, killLimbo) {
			sim.deathEvent(target)
		} else {
			sim.enemies[enemyIdx] = target
			enemyIdx++
		}
	}
	sim.enemies = sim.enemies[:enemyIdx]
}

func (sim *Simulation) kill(target key.TargetID, killLimbo bool) bool {
	state := sim.Attr.State(target)
	switch {
	case state == info.Alive:
		return false
	case state == info.Limbo:
		return killLimbo
	default:
		return true
	}
}

func (sim *Simulation) deathEvent(target key.TargetID) {
	sim.Event.TargetDeath.Emit(event.TargetDeath{
		Target: target,
		Killer: sim.Attr.LastAttacker(target),
	})
}
