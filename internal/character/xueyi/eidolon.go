package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 = "xueyi-e1"
	E2 = "xueyi-e2"
	E4 = "xueyi-e4"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeHit: E1DamageBuff,
		},
	})

	modifier.Register(E4, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func E1DamageBuff(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_INSERT {
		e.Hit.Attacker.AddProperty(E1, prop.AllDamagePercent, 0.4)
	}
}
