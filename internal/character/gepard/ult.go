package gepard

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult       key.Modifier = "gepard-ult"
	UltShield key.Shield   = "gepard-ult-shield"
)

type ultState struct {
	shieldPerc float64
	shieldFlat float64
}

func init() {
	modifier.Register(Ult, modifier.Config{
		Duration:   3,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				mod.Engine().AddShield(UltShield, info.Shield{
					Source:      mod.Source(),
					Target:      mod.Owner(),
					BaseShield:  info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: mod.State().(ultState).shieldPerc},
					ShieldValue: mod.State().(ultState).shieldFlat,
				})
			},
			OnRemove: func(mod *modifier.ModifierInstance) {
				mod.Engine().RemoveShield(UltShield, mod.Owner())
			},
		},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:   Ult,
			Source: c.id,
			State: ultState{
				shieldPerc: ultShieldPerc[c.info.UltLevelIndex()],
				shieldFlat: ultShieldFlat[c.info.UltLevelIndex()],
			},
			TickImmediately: true,
		})
	}

	c.engine.ModifyEnergy(c.id, 5.0)
}
