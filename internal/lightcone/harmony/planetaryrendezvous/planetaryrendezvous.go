package planetaryrendezvous

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
	dmgBuff key.Modifier = "planetary-rendezvous"
)

// After entering battle, if an ally deals the same DMG Type as the wearer,
// DMG dealt increases by 12%.

func init() {
	lightcone.Register(key.PlanetaryRendezvous, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
	modifier.Register(dmgBuff, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	dmgAmt := 0.09 + 0.03*float64(lc.Imposition)

	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		holderInfo, _ := engine.CharacterInfo(owner)
		for _, char := range engine.Characters() {
			// add lcholder's element's buff to all chars
			engine.AddModifier(char, info.Modifier{
				Name:   dmgBuff,
				Source: owner,
				Stats:  info.PropMap{prop.DamagePercent(holderInfo.Element): dmgAmt},
			})
		}
	})
}
