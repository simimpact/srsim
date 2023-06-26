package simulation

import "github.com/simimpact/srsim/pkg/engine/info"

func (sim *Simulation) Attack(atk info.Attack) {
	for _, t := range atk.Targets {
		sim.ActionTargets[t] = true
	}
	sim.Combat.Attack(atk)
}

func (sim *Simulation) Heal(heal info.Heal) {
	for _, t := range heal.Targets {
		sim.ActionTargets[t] = true
	}
	sim.Combat.Heal(heal)
}

func (sim *Simulation) EndAttack() {
	sim.Combat.EndAttack()
}
