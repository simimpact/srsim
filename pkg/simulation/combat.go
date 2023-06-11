package simulation

import "github.com/simimpact/srsim/pkg/engine/info"

func (sim *simulation) Attack(atk info.Attack) {
	sim.combat.Attack(atk)
}

func (sim *simulation) Heal(heal info.Heal) {
	sim.combat.Heal(heal)
}

func (sim *simulation) EndAttack() {
	sim.combat.EndAttack()
}
