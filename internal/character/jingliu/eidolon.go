package jingliu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2     = "jingliu-e2"
	E2Buff = "jingliu-e2-buff"
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
}

func E2Listener(mod *modifier.Instance, e event.ActionStart) {
	tmp, _ := mod.Engine().CharacterInstance(mod.Owner())
	c := tmp.(*char)
	if c.isEnhanced {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   E2Buff,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: 0.8},
		})
		mod.RemoveSelf()
	}
}
