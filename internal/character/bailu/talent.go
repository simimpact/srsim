package bailu

import "github.com/simimpact/srsim/pkg/engine/modifier"

func init() {
	modifier.Register(invigoration, modifier.Config{
		Stacking: modifier.Prolong,
	})
}
