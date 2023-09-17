package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4 = "danhengimbibitorlunae-a4"
	A6 = "danhengimbibitorlunae-a6"
)

func init() {
	modifier.Register(A4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(A6, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnBeforeHit: A6OnBeforeHit,
			OnAfterHit:  A6OnAfterHit,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.ModifyEnergy(info.ModifyAttribute{})
	}
	if c.info.Traces["102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4,
			Source: c.id,
			State:  info.PropMap{prop.EffectRES: 0.35},
		})
	}
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}
func A6OnBeforeHit(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.Defender.IsWeakTo(model.DamageType_IMAGINARY) {
		mod.SetProperty(prop.CritDMG, 0.24)
	}
}
func A6OnAfterHit(mod *modifier.Instance, e event.HitEnd) {
	mod.SetProperty(prop.CritDMG, 0)
}
