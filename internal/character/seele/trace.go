package seele

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	A2Check key.Modifier = "seele-a2-check"
	A2Aggro key.Modifier = "seele-a2-aggro"
)

// A2: When current HP percentage is 50% or lower, reduces the chance of being attacked by enemies

func init() {
	modifier.Register(A2Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnHPChange: reduceAggro,
		},
	})
	modifier.Register(A2Aggro, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func reduceAggro(mod *modifier.Instance, e event.HPChange) {
	if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   A2Aggro,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AggroPercent: -0.5},
		})
	}
}
