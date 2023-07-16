package timewaitsfornoone

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
	time = "time-waits-for-no-one"
	cd   = "time-waits-for-no-one-extra-dmg-cd"
)

// Desc : Increases the wearer's Max HP by 18% and Outgoing Healing by 12%.
// When the wearer heals allies, record the amount of Outgoing Healing.
// When any ally launches an attack, a random attacked enemy takes Additional DMG
// equal to 36% of the recorded Outgoing Healing value.
// The type of this Additional DMG is of the same Type as the wearer's.
// This Additional DMG is not affected by other buffs, and can only occur 1 time per turn.

func init() {
	lightcone.Register(key.TimeWaitsforNoOne, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_ABUNDANCE,
		Promotions:    promotions,
	})
	modifier.Register(time, modifier.Config{})
	// added to automatically track extraDmg cooldown
	modifier.Register(cd, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// modifier only adds the hp and outgoing heal buffs
	engine.AddModifier(owner, info.Modifier{
		Name:   time,
		Source: owner,
		Stats: info.PropMap{
			prop.HPPercent: 0.15 + 0.03*float64(lc.Imposition),
			prop.HealBoost: 0.10 + 0.02*float64(lc.Imposition),
		},
	})

	extraDmgMult := 0.30 + 0.06*float64(lc.Imposition)
	healAmt := 0.0

	// record lc holder's heals
	engine.Events().HealEnd.Subscribe(func(e event.HealEnd) {
		if e.Healer == owner {
			healAmt += e.HealAmount
		}
	})

	// apply extra dmg if not on cooldown, then reset recorded heals and set to cooldown
	engine.Events().AttackEnd.Subscribe(func(e event.AttackEnd) {
		if engine.IsCharacter(e.Attacker) && !engine.HasModifier(owner, cd) {
			applyExtraDmg(engine, e.Targets, owner, healAmt*extraDmgMult)

			healAmt = 0.0                            // reset heal amount
			engine.AddModifier(owner, info.Modifier{ // add 1 turn cd modifier
				Name:     cd,
				Source:   owner,
				Duration: 1,
			})
		}
	})
}

// choose 1 in targets list, apply pure pursued damage
func applyExtraDmg(engine engine.Engine, targets []key.TargetID, owner key.TargetID, dmgAmt float64) {
	chosenEnemy := engine.Retarget(info.Retarget{
		Targets: targets,
		Max:     1,
	})
	holderInfo, _ := engine.CharacterInfo(owner)
	engine.Attack(info.Attack{
		Key:          time,
		Targets:      chosenEnemy,
		Source:       owner,
		AttackType:   model.AttackType_PURSUED,
		DamageType:   holderInfo.Element,
		DamageValue:  dmgAmt,
		AsPureDamage: true,
		// NOTE : might need to later change BaseDmg field to optional later
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: 0.0,
		},
		UseSnapshot: true,
	})
}
