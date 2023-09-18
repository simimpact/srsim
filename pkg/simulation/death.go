package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

// attempt to kill off any dead targets and remove them from the battlefield
func (sim *Simulation) deathCheck(killLimbo bool) {
	toKill := make([]key.TargetID, 0, 10)

	charIdx := 0
	for _, target := range sim.characters {
		if sim.kill(target, killLimbo) {
			toKill = append(toKill, target)
		} else {
			sim.characters[charIdx] = target
			charIdx++
		}
	}
	sim.characters = sim.characters[:charIdx]

	enemyIdx := 0
	for _, target := range sim.enemies {
		if sim.kill(target, killLimbo) {
			toKill = append(toKill, target)
		} else {
			sim.enemies[enemyIdx] = target
			enemyIdx++
		}
	}
	sim.enemies = sim.enemies[:enemyIdx]

	// TODO: RemoveTarget -> RemoveTargets for better efficiency
	for _, target := range toKill {
		// remove this target from the turn order
		sim.Turn.RemoveTarget(target)
		sim.deathEvent(target)
	}
}

func (sim *Simulation) kill(target key.TargetID, killLimbo bool) bool {
	switch sim.Attr.State(target) {
	case info.Dead:
		return true
	case info.Alive:
		return false
	case info.Limbo:
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
