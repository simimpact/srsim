package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Simulation) subscribe() {
	s.Event.TargetDeath.Subscribe(s.onDeath)
}

func (s *Simulation) onDeath(e event.TargetDeathEvent) {
	// remove this target from active arrays (these arrays represent order in battle map)
	switch s.Targets[e.Target] {
	case info.ClassCharacter:
		s.characters = remove(s.characters, e.Target)
	case info.ClassEnemy:
		s.enemies = remove(s.enemies, e.Target)
	case info.ClassNeutral:
		s.neutrals = remove(s.neutrals, e.Target)
	}

	// remove this target from the turn order
	s.Turn.RemoveTarget(e.Target)
}

func remove(arr []key.TargetID, id key.TargetID) []key.TargetID {
	for i, t := range arr {
		if id == t {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
