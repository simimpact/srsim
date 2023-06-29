package vonwacq

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	mod key.Modifier = "sprightly-vonwacq"
)

// 2pc:
// Increases the wearer's Energy Regeneration Rate by 5%.
// When the wearer's SPD reaches 120 or higher,
// the wearer's action is Advanced Forward by 40% immediately upon entering battle.

// Multi Wave Support not implemented

func init() {
	relic.Register(key.SprightlyVonwacq, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				Stats:    info.PropMap{prop.EnergyRegen: 0.05},
			},
			{
				MinCount:     4,
				CreateEffect: Create,
			},
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID) {
	engine.Events().BattleStart.Subscribe(func(e event.BattleStartEvent) {
		for _, char := range e.CharStats {
			if char.ID() == owner && char.SPD() >= 120 {
				engine.ModifyGaugeNormalized(char.ID(), -0.4)
			}
		}
	})
}
