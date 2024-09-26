package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A6 = "jingliu-a6"
)

func (c *char) A6Listener(e event.HitStart) {
	if c.info.Traces["103"] && e.Hit.AttackType == model.AttackType_ULT && c.isEnhanced {
		e.Hit.Attacker.AddProperty(A6, prop.AllDamagePercent, 0.2)
	}
}
