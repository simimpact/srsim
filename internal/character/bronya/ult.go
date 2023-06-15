package bronya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult key.Modifier = "bronya-ult"
)

func init() {
	modifier.Register(Ult, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   2,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	bronyaCDmg := c.engine.Stats(c.id).CritDamage()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   Ult,
			Source: c.id,
			Stats: info.PropMap{prop.ATKPercent: ultAtkPerc[c.info.AbilityLevel.Ult-1],
				prop.CritDMG: ultCDmgDefault[c.info.AbilityLevel.Ult-1] +
					ultCDmgBronya[c.info.AbilityLevel.Ult-1]*bronyaCDmg},
			TickImmediately: true,
		})
	}

	c.engine.ModifyEnergy(c.id, 5.0)
}
