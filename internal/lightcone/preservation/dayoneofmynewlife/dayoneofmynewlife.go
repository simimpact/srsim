package dayoneofmynewlife

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
	mod     key.Modifier = "dayone"
	modaura key.Modifier = "dayone-aura"
)

// Increases the wearer's DEF by 16%/18%/20%/22%/24%. After entering battle,
// increases All-Type RES of all allies by 8%/9%/10%/11%/12%. Effects of the
// same type cannot stack.
func init() {
	lightcone.Register(key.DayOneofMyNewLife, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(modaura, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeDying: onWearerDeath,
		},
		Stacking: modifier.ReplaceBySource,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// wearer DEF%
	amt := 0.14 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: amt},
	})

	// allies AllType RES
	amtAura := 0.07 + 0.01*float64(lc.Imposition)
	mod := info.Modifier{
		Name:   modaura,
		Source: owner,
		Stats:  info.PropMap{prop.AllDamageRES: amtAura},
	}

	// NOTE: fine to leave this as-is, but this may have to change in the
	// future when adding support for multiple waves. `BattleStart` happens
	// at the beginning of each wave
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		for char := range event.CharInfo {
			engine.AddModifier(char, mod)
		}
	})
}

// removes team aura on wearer's death
func onWearerDeath(mod *modifier.Instance) {
	for _, char := range mod.Engine().Characters() {
		mod.Engine().RemoveModifierFromSource(char, mod.Owner(), modaura)
	}
}
