package simulation

import "github.com/simimpact/srsim/pkg/engine/info"

func (sim *simulation) Attack(atk info.Attack) {
	for _, t := range atk.Targets {
		sim.actionTargets[t] = true
	}
	sim.combat.Attack(atk)
}

func (sim *simulation) Heal(heal info.Heal) {
	for _, t := range heal.Targets {
		sim.actionTargets[t] = true
	}
	sim.combat.Heal(heal)
}

func (sim *simulation) EndAttack() {
	sim.combat.EndAttack()
}
