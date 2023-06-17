package returntodarkness

import (
	"math/rand"

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
	ReturntoDarkness key.Modifier = "return_to_darkness"
)

var wasTriggered bool = false

// Increases the wearer's CRIT Rate by 12/15/18/21/24%. After a CRIT Hit, there
// is a 16/20/24/28/32% fixed chance to dispel 1 buff on the target enemy. This
// effect can only trigger once per attack.
func init() {
	lightcone.Register(key.ReturntoDarkness, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(ReturntoDarkness, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHitAll: onAfterHitAll,
			OnAfterAttack: onAfterAttack,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	cr := 0.09 + 0.03*float64(lc.Imposition)
	chance := 0.12 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   ReturntoDarkness,
		Source: owner,
		Stats: info.PropMap{
			prop.CritChance: cr,
		},
		State: chance,
	})
}

func onAfterHitAll(mod *modifier.ModifierInstance, e event.HitEndEvent) {
	if e.IsCrit && !wasTriggered && rand.Float64() < mod.State().(float64) {
		mod.Engine().DispelStatus(e.Defender, info.Dispel{
			Status: model.StatusType_STATUS_BUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
		wasTriggered = true
	}
}

func onAfterAttack(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
	wasTriggered = false
}
