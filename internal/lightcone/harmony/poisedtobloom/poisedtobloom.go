package poisedtobloom

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
	poised        key.Modifier = "poised-to-bloom"
	poisedCritDmg key.Modifier = "poised-to-bloom-crit-dmg"
)

// Lose Not, Forget Not
// Increases the wearer's ATK by 16/20/24/28/32%. Upon entering battle,
// if two or more characters follow the same Path,
// then these characters' CRIT DMG increases by 16/20/24/28/32%.
// Abilities of the same type cannot stack.

func init() {
	lightcone.Register(key.PoisedToBloom, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})

	modifier.Register(poised, modifier.Config{})

	modifier.Register(poisedCritDmg, modifier.Config{
		Stacking: modifier.Replace,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	atkAmt := 0.12 + 0.04*float64(lc.Imposition)
	critDmgAmt := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   poised,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: atkAmt},
	})

	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		// This is probably slow, but I can't think of other good ways to iterate
		// and store paths that don't involve allocating memory
		// that might not be used, which I suspect would be slower
		// It's still only called once per iteration though, so it should
		// be fine.

		// Iterating over all the characters
		for _, charA := range engine.Characters() {
			charAInfo, _ := engine.CharacterInfo(charA)
			// Checking for pairs with them
			for _, charB := range engine.Characters() {
				charBInfo, _ := engine.CharacterInfo(charB)
				// If there's a pair, we apply the crit dmg buff to charA
				// we could also apply it to charB, but with no way to remove
				// them from the future iterations, that would actually be slower
				if charB != charA && charAInfo.Path == charBInfo.Path {
					engine.AddModifier(charA, info.Modifier{
						Name:   poisedCritDmg,
						Source: owner,
						Stats:  info.PropMap{prop.CritDMG: critDmgAmt},
					})
					break
				}
			}
		}
	})
}
