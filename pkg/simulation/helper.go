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

func (s *simulation) createSnapshot() snapshot {
	charStats := make([]*info.Stats, len(s.characters))
	for i, t := range s.characters {
		charStats[i] = s.attr.Stats(t)
	}
	enemyStats := make([]*info.Stats, len(s.enemies))
	for i, t := range s.enemies {
		enemyStats[i] = s.attr.Stats(t)
	}
	neutralStats := make([]*info.Stats, len(s.neutrals))
	for i, t := range s.neutrals {
		neutralStats[i] = s.attr.Stats(t)
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
func (s *simulation) deathCheck(targets []key.TargetID) {
	for _, target := range targets {
		if s.attr.HPRatio(target) <= 0 {
			s.event.TargetDeath.Emit(event.TargetDeathEvent{
				Target: target,
				Killer: target,
			})
		}
	}
}
