package goodnightandsleepwell

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
	Check key.Modifier = "good_night_and_sleep_well_check"
)

func init() {
	lightcone.Register(key.GoodNightandSleepWell, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		CanModifySnapshot: true,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitAll,
		},
	})
}

// For every debuff the target enemy has, the DMG dealt by the wearer increases by 12%/15%/18%/21%/24%,
// stacking up to 3 time(s). This effect also applies to DoT.
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.09 + 0.03*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		State:  amt,
	})
}

func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	debuffCount := float64(e.Hit.Defender.StatusCount(model.StatusType_STATUS_DEBUFF))
	if debuffCount > 3 {
		debuffCount = 3
	}
	amt := mod.State().(float64) * debuffCount
	e.Hit.Attacker.AddProperty(prop.AllDamagePercent, amt)
}
