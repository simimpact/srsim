package izumo

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	name = "izumo-gensei-and-takama-divine-realm"
)

// Increases the wearer's ATK by 12%.
// When entering battle, if at least one other ally follows the same Path as the wearer, then the wearer's CRIT Rate increases by 12%.
func init() {
	relic.Register(key.IzumoGensei, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.ATKPercent: 0.12},
				CreateEffect: Create,
			},
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID) {
	ownerInfo, _ := engine.CharacterInfo(owner)

	engine.Events().BattleStart.Subscribe(func(e event.BattleStart) {
		samePath := 0

		for _, char := range engine.Characters() {
			charInfo, _ := engine.CharacterInfo(char)
			if char != owner && charInfo.Path == ownerInfo.Path {
				samePath++
				break
			}
		}

		if samePath >= 1 {
			engine.AddModifier(owner, info.Modifier{
				Name:   name,
				Source: owner,
				Stats:  info.PropMap{prop.CritChance: 0.12},
			})
		}
	})
}
