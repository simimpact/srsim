package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	E2 = "sparkle-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: E2Callback,
		},
		MaxCount: 3,
	})
}

func E2Callback(mod *modifier.Instance, e event.HitStart) {
	e.Hit.Defender.AddProperty(E2, prop.DEFPercent, -0.08*mod.Count())
}
