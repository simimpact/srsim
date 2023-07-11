package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	//If eidolon 2 or above, will attempt to apply HOT.
	c.e2(targets)

	c.engine.Heal(info.Heal{
		Targets: targets,
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: ult[c.info.UltLevelIndex()],
		},
		HealValue: ultFlat[c.info.UltLevelIndex()],
	})

	c.engine.ModifyEnergy(c.id, 5)

}
