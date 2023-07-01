package asecretvow

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
	ASecretVow key.Modifier = "a_secret_vow"
)

// Increases DMG dealt by the wearer by 20%.
// The wearer also deals an extra 20% of DMG to enemies whose current HP percentage is equal to or higher than the wearer's current HP percentage.
func init() {
	lightcone.Register(key.ASecretVow, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(ASecretVow, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})
}

// Add dmg% modifier
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.15 + 0.05*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   ASecretVow,
		Source: owner,
		Stats:  info.PropMap{prop.AllDamagePercent: amt},
		State:  amt,
	})
}

// If the enemy hp ratio is greater or equal than the attackers hp ratio, add dmg%
func onBeforeHitAll(mod *modifier.Instance, e event.HitStartEvent) {
	if e.Hit.Attacker.CurrentHPRatio() <= e.Hit.Defender.CurrentHPRatio() {
		e.Hit.Attacker.AddProperty(prop.AllDamagePercent, mod.State().(float64))
	}
}
