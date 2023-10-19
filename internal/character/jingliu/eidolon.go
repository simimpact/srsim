package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1     = "jingliu-e1"
	E2     = "jingliu-e2-listener"
	E2Buff = "jingliu-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_UNKNOWN_STATUS,
		Listeners: modifier.Listeners{
			OnBeforeAction: E2Listener,
		},
	})
	modifier.Register(E2Buff, modifier.Config{
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
			Name:            E1,
			Source:          c.id,
			Stats:           info.PropMap{prop.CritDMG: 0.24},
			TickImmediately: true,
		})
	}
}

func E2Listener(mod *modifier.Instance, e event.ActionStart) {
	tmp, _ := mod.Engine().CharacterInstance(mod.Owner())
	c := tmp.(*char)
	if c.isEnhanced && e.AttackType == model.AttackType_SKILL {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   E2Buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: 0.8},
		})
		mod.RemoveSelf()
	}
}
