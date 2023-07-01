package shatteredhome

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
	ShatteredHomeCheck key.Modifier = "shattered-home-check"
	ShatteredHomeBuff  key.Modifier = "shattered-home-buff"
)

func init() {
	lightcone.Register(key.ShatteredHome, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(ShatteredHomeCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})
}

// The wear deals 20%/25%/30%/35%/40% more DMG to enemy targets whose HP percentage is greater than 50%.
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   ShatteredHomeCheck,
		Source: owner,
		State:  0.15 + 0.05*float64(lc.Imposition),
	})
}

func onBeforeHitAll(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	if e.Hit.Defender.CurrentHPRatio() > 0.5 {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, mod.State().(float64))
	}
}
