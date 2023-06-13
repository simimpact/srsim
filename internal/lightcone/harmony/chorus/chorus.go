package chorus

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Chorus key.Modifier = "chorus"
)

func init() {
	lightcone.Register(key.Chorus, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})

	modifier.Register(Chorus, modifier.Config{
		Stacking: modifier.Replace,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	mod := info.Modifier{
		Name:   Chorus,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: 0.08 + 0.01*float64(lc.Imposition)},
	}

	engine.Events().BattleStart.Subscribe(func(event event.BattleStartEvent) {
		for char := range event.CharInfo {
			engine.AddModifier(char, mod)
		}
	})
}
