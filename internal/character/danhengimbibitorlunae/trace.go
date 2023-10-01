package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4            = "danhengimbibitorlunae-a4"
	A6            = "danhengimbibitorlunae-a6"
	A2 key.Reason = "danhengimbibitorlunae-a4"
)

func (c *char) initTraces() {
	modifier.Register(A4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(A6, modifier.Config{
		StatusType: model.StatusType_UNKNOWN_STATUS,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: A6OnHit,
		},
		CanModifySnapshot: true,
	})
	if c.info.Traces["101"] {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    A2,
			Target: c.id,
			Source: c.id,
			Amount: 15,
		})
	}
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:      A4,
			Source:    c.id,
			DebuffRES: info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.35},
		})
	}
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}
func A6OnHit(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.Defender.IsWeakTo(model.DamageType_IMAGINARY) {
		e.Hit.Attacker.AddProperty(A6, prop.CritDMG, 0.24)
	}
}
