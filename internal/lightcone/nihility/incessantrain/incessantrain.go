package incessantrain

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
	ehrNCritCheck key.Modifier = "incessant-rain-ehr-crit-check"
	critMod       key.Modifier = "incessant-rain-crit-rate"
	aetherCode    key.Modifier = "incessant-rain-aether-code"
)

// Desc : Increases the wearer's Effect Hit Rate by 24%.
// When the wearer deals DMG to an enemy that currently has 3 or more debuffs,
// increases the wearer's CRIT Rate by 12%. After the wearer uses their Basic ATK, Skill, or Ultimate,
// there is a 100% base chance to implant Aether Code on a random hit target that does not yet have it.
// Targets with Aether Code receive 12% increased DMG for 1 turn.
// Apparent mods : EHR, CRBoost, Aether Code implant (dmgTakenIncrease)

// Quick Imp Plan :
// EHR perm buff : add it on create.
// Modifiers :
// - Crit Rate : Listener = OnBeforeHitAll, check each enemy if it has 3 debuffs => add critrate boost
// - aetherCode : Listener = OnAfterAction, choose 1 among hit targets on last atk, aetherCode 100% BC
//	 -> check first if target already have AC
//	=> OnBeforeAction = set cd to 1
//  => set aetherCode as enemy modifier DmgTakenUp.

func init() {
	lightcone.Register(key.IncessantRain, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(ehrNCritCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: critRateBoost,
		},
	})
	modifier.Register(aetherCode, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: resetCooldown,
			OnAfterAction:  applyDebuffOnce,
		},
	})
	modifier.Register(critMod, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF, // is this critrate boost a mod?
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	ehrAmt := 0.20 + 0.04*float64(lc.Imposition)
	critAmt := 0.10 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   ehrNCritCheck,
		Source: owner,
		Stats:  info.PropMap{prop.EffectHitRate: ehrAmt},
		State:  critAmt,
	})
}

// boost CR if enemy has >=3 debuffs
func critRateBoost(mod *modifier.Instance, e event.HitStart) {
	debuffCount := float64(e.Hit.Defender.StatusCount(model.StatusType_STATUS_DEBUFF))
	critRateAmt := mod.State().(float64)
	if debuffCount >= 3 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   critMod,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.CritChance: critRateAmt},
		})
	}
}
// reset aethercode cooldown when char turn is up
func resetCooldown(mod *modifier.Instance, e event.ActionStart) {

}

// retarget with 1 chosen(among non-AC-applied). apply dmgTakenUp with 100% basechance. cooldown tick.
func applyDebuffOnce(mod *modifier.Instance, e event.ActionEnd) {

}
