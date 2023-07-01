package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (sim *Simulation) subscribe() {
	sim.Event.TargetDeath.Subscribe(sim.onDeath)
}

func (sim *Simulation) onDeath(e event.TargetDeathEvent) {
	// remove this target from active arrays (these arrays represent order in battle map)
	switch sim.Targets[e.Target] {
	case info.ClassCharacter:
		sim.characters = remove(sim.characters, e.Target)
	case info.ClassEnemy:
		sim.enemies = remove(sim.enemies, e.Target)
	case info.ClassNeutral:
		sim.neutrals = remove(sim.neutrals, e.Target)
	}

	// remove this target from the turn order
	sim.Turn.RemoveTarget(e.Target)
}

func remove(arr []key.TargetID, id key.TargetID) []key.TargetID {
	for i, t := range arr {
		if id == t {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
