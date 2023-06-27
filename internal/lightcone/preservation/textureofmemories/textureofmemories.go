package textureofmemories

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
	mod    key.Modifier = "texture"
	shield key.Modifier = "texture-shield"

	modshield0      key.Modifier = "texture-attacked-noshield"
	cooldownshield0 key.Modifier = "texture-attacked-noshield-cooldown"
	modshield1      key.Modifier = "texture-attacked-shield"
)

// Increases the wearer's Effect RES by 8%/10%/12%/14%/16%. If the wearer is
// attacked and has no Shield, they gain a Shield equal to
// 16%/20%/24%/28%/32% of their Max HP for 2 turn(s). This effect can only be
// triggered once every 3 turn(s). If the wearer has a Shield when attacked,
// the DMG they receive decreases by 12%/15%/18%/21%/24%.
func init() {
	lightcone.Register(key.TextureofMemories, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(cooldownshield0, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// wearer Effect RES
	amt := 0.06 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   mod,
		Source: owner,
		Stats:  info.PropMap{prop.EffectRES: amt},
		State:  amt,
	})

}

// wearer has Shield before attacked
func onBeforeHit(mod *modifier.ModifierInstance, lc info.LightCone) {
	amtshield1 := 0.09 + 0.03*float64(lc.Imposition)
	if mod.Engine().HasShield(mod.Owner(), "any" /* any shield key here */) {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   modshield1,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamageRES: amtshield1},
			State:  amtshield1,
		})
	}
}

// wearer doesn't have Shield after attacked
func onAfterHit(mod *modifier.ModifierInstance, lc info.LightCone) {
	if !mod.Engine().HasShield(mod.Owner(), "any") {
		mod.Engine().AddShield(key.Shield(shield), info.Shield{
			Source: mod.Owner(),
			Target: mod.Owner(),
			// TODO: VALUES
		})
	}
}
