package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	e1 = "himeko-e1"
	e2 = "himeko-e2"
)

func init() {
	modifier.Register(e1, modifier.Config{
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(e2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: e2Listener,
		},
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   e2,
			Source: c.id,
		})
	}
}

func e2Listener(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HPRatio(e.Defender) <= 0.5 {
		e.Hit.Attacker.AddProperty(e2, prop.AllDamagePercent, 0.15)
	}
}
