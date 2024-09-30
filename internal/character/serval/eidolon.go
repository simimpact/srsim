package serval

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	e6 key.Modifier = "serval-e6"
)

func init() {
	modifier.Register(e6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: e6Listener,
		},
	})
}

func e6Listener(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HasModifier(e.Defender, common.Shock) {
		e.Hit.Attacker.AddProperty("serval-e6", prop.AllDamagePercent, 0.3)
	}
}
