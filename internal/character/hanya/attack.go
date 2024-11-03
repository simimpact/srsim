package hanya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal key.Attack = "hanya-normal"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:     Normal,
		Source:  c.id,
		Targets: []key.TargetID{target},
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   20,
	})

	state.EndAttack()
}
