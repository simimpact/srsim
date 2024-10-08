package pastselfinmirror

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
	psim        = "past-self-in-mirror"
	psimDmgBuff = "past-self-in-mirror-dmg-buff"
	psimEnergy  = "past-self-in-mirror-energy"
)

// Increases the wearer's Break Effect by 60%.
// When the wearer uses their Ultimate, increases all allies' DMG by 24%, lasting for 3 turn(s).
// Should the wearer's Break Effect exceed or equal 150%, 1 Skill Point will be recovered.
// At the start of each wave, all allies regenerate 10 Energy immediately. Abilities of the same type cannot stack.
func init() {
	lightcone.Register(key.PastSelfinMirror, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_HARMONY,
		Promotions:    promotions,
	})
	modifier.Register(psimDmgBuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: afterUlt,
		},
	})
}

// Break Effect Buff
func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.5 + 0.1*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   psim,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: amt},
		State:  float64(lc.Imposition),
	})

	energyAmt := 7.5 + 2.5*float64(lc.Imposition)
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		for _, char := range engine.Characters() {
			engine.ModifyEnergy(info.ModifyAttribute{
				Key:    psimEnergy,
				Target: char,
				Source: owner,
				Amount: energyAmt,
			})
		}
	})
}

func afterUlt(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType != model.AttackType_ULT {
		return
	}

	// Restore 1 SP if BE >= 150%
	if mod.OwnerStats().BreakEffect() >= 1.5 {
		mod.Engine().ModifySP(info.ModifySP{
			Key:    psim,
			Source: mod.Owner(),
			Amount: 1,
		})
	}

	// DMG% buff for 3 turns
	amt := 0.2 + 0.04*mod.State().(float64)
	for _, char := range mod.Engine().Characters() {
		mod.Engine().AddModifier(char, info.Modifier{
			Name:     psimDmgBuff,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.AllDamagePercent: amt},
			Duration: 3,
		})
	}
}
