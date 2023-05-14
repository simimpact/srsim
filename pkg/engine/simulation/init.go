package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/turn"
	"github.com/simimpact/srsim/pkg/key"
)

func (s *Simulation) setupServices() {
	someFakeTargetID := []key.TargetID{2,3,4}
	s.turnManager = turn.New(someFakeTargetID)
}