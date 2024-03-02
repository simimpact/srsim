package herta

import "github.com/simimpact/srsim/pkg/engine/modifier"

const (
	e1 = "herta-e1"
	e2 = "herta-e2"
	e6 = "herta-e6"
)

func init() {
	modifier.Register(e2, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		CountAddWhenStack: 1,
		Count:             1,
	})

	modifier.Register(e6, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Duration: 1,
	})
}
