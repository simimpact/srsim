package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1      key.Modifier = "qingque-e1"
	E2      key.Reason   = "qingque-e2"
	Autarky key.Modifier = "qingque-e4"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.AttackType != model.AttackType_ULT {
					return
				}
				e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.1)
			},
		},
	})
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    E2,
			Target: c.id,
			Source: c.id,
			Amount: 1,
		})
	}
}
func (c *char) initEidolons() {
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E1,
			Source: c.id,
		})
	}
}
