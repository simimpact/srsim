package ruanmei

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1         = "ruanmei-e1"
	E2         = "ruanmei-e2"
	E4         = "ruanmei-e4"
	E4Listener = "ruanmei-e4-listener"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyE1,
		},
	})
	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyE2,
		},
	})
	// E4 is summarized to 1 buff mod
	modifier.Register(E4, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})
	// Causes known bug (?) of the breaking hit not benefitting from E4 as this should apply with OnBeforeBeingBreak
	modifier.Register(E4Listener, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeingBreak: func(mod *modifier.Instance) {
				mod.Engine().AddModifier(mod.Source(), info.Modifier{
					Name:     E4,
					Source:   mod.Source(),
					Stats:    info.PropMap{prop.BreakEffect: 1},
					Duration: 3,
				})
			},
		},
	})
}

func applyE1(mod *modifier.Instance, e event.HitStart) {
	e.Hit.Defender.AddProperty(E1, prop.DEFPercent, -0.2)
}

func applyE2(mod *modifier.Instance, e event.HitStart) {
	// Needs to check for Break flag (uses workaround for now)
	if mod.Engine().Stance(mod.Owner()) == 0 {
		e.Hit.Attacker.AddProperty(E2, prop.ATKPercent, 0.4)
	}
}
