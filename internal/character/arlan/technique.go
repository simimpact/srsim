package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Technique key.Attack = "arlan-technique"

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Technique,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.8,
		},
		StanceDamage: 60.0,
		EnergyGain:   0.0,
	})
}
