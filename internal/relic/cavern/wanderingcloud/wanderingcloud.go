package wanderingcloud

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "passerby-of-wandering-cloud"
)

// 2pc:
// Increases Outgoing Healing by 10%
// 4pc:
// At the start of the battle, immediately regenerates 1 Skill Point.

func init() {
	relic.Register(key.PasserbyOfWanderingCloud, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.HealBoost: 0.10},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.ModifySP(1)
				},
			},
		},
	})
}
