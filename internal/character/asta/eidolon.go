package asta

import "github.com/simimpact/srsim/pkg/engine/modifier"

const (
	e2 = "asta-e2-flag"
	e4 = "asta-e4"
)

func init() {
	modifier.Register(e2, modifier.Config{
		Duration: 1,
	})

	modifier.Register(e4, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}
