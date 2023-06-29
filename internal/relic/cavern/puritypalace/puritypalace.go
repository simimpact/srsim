package puritypalace

import (
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "knight-of-purity-palace"
)

// 2pc: Increases DEF by 15%.
// 4pc: Increases the max DMG that can be absorbed by the Shield created by the wearer by 20%.
func init() {
	relic.Register(key.KnightOfPurityPalace, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.DEFPercent: 0.15},
			},
			{
				MinCount: 4,
				Stats:    info.PropMap{prop.ShieldBoost: 0.20},
			},
		},
	})
}
