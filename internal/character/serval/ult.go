package serval

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const Ult key.Attack = "serval-ult"
const shock key.Modifier = "common-shock"

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Ult,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		DamageType: model.DamageType_THUNDER,
		AttackType: model.AttackType_ULT,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
		},
		StanceDamage: 60.0,
		EnergyGain:   5,
	})

	for _, trg := range c.engine.Enemies() {
		// todo investigate shock lol
		c.engine.ExtendModifierDuration(trg, shock, 2)
	}
}
