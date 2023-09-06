package genius

import (
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	genius key.Modifier = "genius-of-brilliant-stars"
)

// 2pc : Increases Quantum DMG by 10%
// 4pc : When the wearer deals DMG to the target enemy, ignores 10% DEF.
//       If the target enemy has Quantum Weakness, the wearer additionally ignores 10% DEF.

func init() {
	relic.Register(key.GeniusOfBrilliantStars, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
			},
			{
				MinCount: 4,
			},
		},
	})
	modifier.Register(genius, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: addResPen,
		},
	})
}

func addResPen(mod *modifier.Instance, e event.HitStart) {

}
