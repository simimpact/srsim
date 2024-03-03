package herta

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	e1 = "herta-e1"
	e2 = "herta-e2"
	e4 = "herta-e4"
	e6 = "herta-e6"
)

func init() {
	modifier.Register(e2, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		CountAddWhenStack: 1,
		Count:             1,
		Listeners: modifier.Listeners{
			OnAdd: e2OnStack,
		},
	})

	modifier.Register(e6, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Duration: 1,
	})
}

func e2OnStack(mod *modifier.Instance) {
	mod.SetProperty(prop.CritChance, 0.03*mod.Count())
}
