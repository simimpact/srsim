package simulation

import (
	"math"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

func (sim *Simulation) initStatCollection() {
	cycleLimit := sim.cfg.Settings.CycleLimit
	if cycleLimit <= 0 {
		cycleLimit = 10
	}
	sim.res = &model.IterationResult{
		TotalDamageDealt:             0,
		TotalDamageTaken:             0,
		TotalAv:                      0,
		CumulativeDamageDealtByCycle: make([]float64, 1, cycleLimit),
		CumulativeDamageTakenByCycle: make([]float64, 1, cycleLimit),
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
		// split up total damage by cycle buckets
		cycle := int(math.Ceil(sim.TotalAV/100)) - 1
		if cycle < 0 {
			cycle = 0
		}
		// TODO: this code could be more performance?
		for len(sim.res.CumulativeDamageDealtByCycle) <= cycle {
			last := len(sim.res.CumulativeDamageDealtByCycle) - 1
			sim.res.CumulativeDamageDealtByCycle = append(sim.res.CumulativeDamageDealtByCycle, sim.res.CumulativeDamageDealtByCycle[last])
		}
		for len(sim.res.CumulativeDamageTakenByCycle) <= cycle {
			last := len(sim.res.CumulativeDamageTakenByCycle) - 1
			sim.res.CumulativeDamageTakenByCycle = append(sim.res.CumulativeDamageTakenByCycle, sim.res.CumulativeDamageTakenByCycle[last])
		}
		sim.res.CumulativeDamageDealtByCycle[cycle] = sim.res.TotalDamageDealt
		sim.res.CumulativeDamageTakenByCycle[cycle] = sim.res.TotalDamageTaken
	})
}
