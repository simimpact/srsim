package gepard

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult key.Shield = "gepard-ult"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	for _, trg := range targets {
		c.engine.AddShield(Ult, info.Shield{
			Source:      c.id,
			Target:      trg,
			BaseShield:  info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: ultShieldPerc[c.info.UltLevelIndex()]},
			ShieldValue: ultShieldFlat[c.info.UltLevelIndex()],
		})
	}

	c.engine.ModifyEnergy(c.id, 5.0)
}
