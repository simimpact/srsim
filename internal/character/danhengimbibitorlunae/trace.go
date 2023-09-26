package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4 key.Reason = "danhengimbibitorlunae-a4"
	A6 key.Reason = "danhengimbibitorlunae-a6"
	A2 key.Reason = "danhengimbibitorlunae-a4"
)

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    A2,
			Target: c.id,
			Source: c.id,
			Amount: 15,
		})
	}
	if c.info.Traces["102"] {
		c.engine.Stats(c.id).AddProperty(A4, prop.EffectRES, 0.35)
	}
	if c.info.Traces["103"] {
		c.engine.Events().HitStart.Subscribe(c.A6OnHit)
	}
}
func (c *char) A6OnHit(e event.HitStart) {
	if e.Hit.Attacker.ID() == c.id && e.Hit.Defender.IsWeakTo(model.DamageType_IMAGINARY) {
		e.Hit.Attacker.AddProperty(A6, prop.CritDMG, 0.24)
	}
}
