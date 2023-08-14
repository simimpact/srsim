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

// TODO : try to implement for each element cases.

type element struct {
}
type elementList struct {
	// fire

	// ice

	// img

	// lightning

	// phys

	// qua

	// wind
}

// After entering battle, if an ally deals the same DMG Type as the wearer,
// DMG dealt increases by 12%.

// IMPL NOTES :
// DM is a mess.
// OnBattleStart : check for each char on the field. use subscribe.
// if char dmg type is same as lc holder, add permanent dmg buff modifier.

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
			// check current char element against lc holder
			charInfo, _ := engine.CharacterInfo(char)
			if charInfo.Element == holderInfo.Element {
				// if element matches. add perm dmg buff.
				// TODO : confirm if the buff is all dmg percent or specific element percent
				// is it mechanically different or nah?
				engine.AddModifier(char, info.Modifier{
					Name:   dmgBuff,
					Source: owner,
					Stats:  info.PropMap{prop.AllDamagePercent: dmgAmt},
				})
			}
		}
	})
}
