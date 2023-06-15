package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_QUANTUM,
		AttackType: model.AttackType_MAZE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.8,
		},
		StanceDamage: 0,
		EnergyGain:   0.0,
	})

	for _, trg := range c.engine.Enemies() {
		c.engine.ModifyStance(trg, c.id, 60)
	}
}
