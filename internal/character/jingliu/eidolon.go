package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 = "jingliu-e1"
	E2 = "jingliu-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAfterAction: removeWhenEndAttack,
		},
	})
	modifier.Register(E1, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   1,
	})
}

func (c *char) E1Listener(e event.ActionStart) {
	if (c.isEnhanced && e.AttackType == model.AttackType_SKILL) || e.AttackType == model.AttackType_ULT {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E1,
			Source: c.id,
			Stats:  info.PropMap{prop.CritDMG: 0.24},
		})
	}
}

func (c *char) E2Listener(e event.HitStart) {
	if c.isEnhanced && c.afterUlt && e.Hit.AttackType == model.AttackType_SKILL {
		e.Hit.Attacker.AddProperty(E2, prop.AllDamagePercent, 0.8)
	}
}
