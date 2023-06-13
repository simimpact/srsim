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

type ultState struct {
	bonusAtkPerc     float64
	bonusCDmgDefault float64
	bonusCDmgBronya  float64
}

func init() {
	modifier.Register(Ult, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.SetProperty(prop.ATKPercent, mod.State().(ultState).bonusAtkPerc)

				cDmg := mod.State().(ultState).bonusCDmgDefault + mod.State().(ultState).bonusCDmgBronya
				mod.SetProperty(prop.CritDMG, cDmg)
			},
		},
		Duration: 2,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   Ult,
			Source: c.id,
			State: ultState{
				bonusAtkPerc:     ultAtkPerc[c.info.AbilityLevel.Ult],
				bonusCDmgDefault: ultCDmgDefault[c.info.AbilityLevel.Ult],
				bonusCDmgBronya:  ultCDmgBronya[c.info.AbilityLevel.Ult] * c.engine.Stats(c.id).CritDamage(),
			},
			TickImmediately: true,
		})
	}
}
