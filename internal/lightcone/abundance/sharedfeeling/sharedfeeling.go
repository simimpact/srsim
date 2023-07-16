package sharedfeeling

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
	SFMod = "shared-feeling"
)

func init() {
	lightcone.Register(key.SharedFeeling, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	// merge heal buff OnStart and energy top up checker
	modifier.Register(SFMod, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: giveTeamEnergy,
		},
	})
}

// Increases the wearer's Outgoing Healing by 10%.
// When using Skill, regenerates 2 Energy for all allies.
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	healBuffAmt := 0.075 + 0.025*float64(lc.Imposition)
	// Checker : + applies energy top up here.
	engine.AddModifier(owner, info.Modifier{
		Name:   SFMod,
		Source: owner,
		Stats:  info.PropMap{prop.HealBoost: healBuffAmt},
		State:  1.5 + 0.5*float64(lc.Imposition),
	})
}

func giveTeamEnergy(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_SKILL {
		// apply team energy top up.
		for _, char := range mod.Engine().Characters() {
			mod.Engine().ModifyEnergy(info.ModifyAttribute{
				Key:    SFMod,
				Target: char,
				Source: mod.Owner(),
				Amount: mod.State().(float64),
			})
		}
	}
}
