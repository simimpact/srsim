package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type snapshot struct {
	characters []*info.Stats
	enemies    []*info.Stats
	neutrals   []*info.Stats
}

func (sim *Simulation) createSnapshot() snapshot {
	charStats := make([]*info.Stats, len(sim.characters))
	for i, t := range sim.characters {
		charStats[i] = sim.Attr.Stats(t)
	}
	enemyStats := make([]*info.Stats, len(sim.enemies))
	for i, t := range sim.enemies {
		enemyStats[i] = sim.Attr.Stats(t)
	}
	neutralStats := make([]*info.Stats, len(sim.neutrals))
	for i, t := range sim.neutrals {
		neutralStats[i] = sim.Attr.Stats(t)
	}
	return snapshot{
		characters: charStats,
		enemies:    enemyStats,
		neutrals:   neutralStats,
	}
}

// This is for handling a special case where a target should be dead but is not.
// atm the only scenario where this can occur is if a target was scheduled to be revived but then
// the reviver died/was incapacitated and the revive was cancelled.
//
// Due to the current sim control flow, the information about who killed this target has been lost.
// Emitting event that this death was a suicide, which is not the ideal behavior
func (sim *Simulation) deathCheck(targets []key.TargetID) {
	for _, target := range targets {
		if sim.Attr.HPRatio(target) <= 0 {
			sim.Event.TargetDeath.Emit(event.TargetDeath{
				Target: target,
				Killer: target,
			})
		}
	}
}
