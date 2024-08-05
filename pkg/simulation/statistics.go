package simulation

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

func (sim *Simulation) initStatCollection() {
	sim.res = &model.IterationResult{
		TotalDamageDealt: 0,
		TotalDamageTaken: 0,
		TotalAv:          0,
	}

	// collect total damage
	sim.Event.HitEnd.Subscribe(func(event event.HitEnd) {
		trg, ok := sim.Targets[event.Defender]
		//TODO: is this check actually necessary? persumably event shouldn't give us an invalid target?
		if !ok {
			return
		}
		switch trg {
		case info.ClassCharacter:
			sim.res.TotalDamageTaken += event.TotalDamage
		case info.ClassEnemy:
			sim.res.TotalDamageDealt += event.TotalDamage
		}
	})
}
