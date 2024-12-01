package gepard

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Technique       key.Modifier = "gepard-technique"
	TechniqueShield key.Shield   = "gepard-technique-shield"
)

func init() {
	modifier.Register(Technique, modifier.Config{
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.Engine().AddShield(TechniqueShield, info.Shield{
					Source:      mod.Source(),
					Target:      mod.Owner(),
					BaseShield:  info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.24},
					ShieldValue: 150,
				})
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(TechniqueShield, mod.Owner())
			},
		},
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	targets := c.engine.Characters()

	for _, trg := range targets {
		c.engine.AddModifier(trg, info.Modifier{
			Name:            Technique,
			Source:          c.id,
			TickImmediately: true,
		})
	}
}
