package yanqing

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TechEffect                  = "yanqing-tech"
	TechEffectReason key.Reason = "yanqing-tech"
)

func init() {
	modifier.Register(TechEffect, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   2,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: TechOnHit,
		},
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   TechEffect,
		Source: c.id,
	})
}

func TechOnHit(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HPRatio(e.Defender) > 0.5 {
		e.Hit.Attacker.AddProperty(TechEffectReason, prop.AllDamagePercent, 0.3)
	}
}
