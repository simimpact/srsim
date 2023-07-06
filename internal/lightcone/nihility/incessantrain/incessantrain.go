package incessantrain

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Desc : Increases the wearer's Effect Hit Rate by 24%.
// When the wearer deals DMG to an enemy that currently has 3 or more debuffs,
// increases the wearer's CRIT Rate by 12%. After the wearer uses their Basic ATK, Skill, or Ultimate,
// there is a 100% base chance to implant Aether Code on a random hit target that does not yet have it.
// Targets with Aether Code receive 12% increased DMG for 1 turn.
// Apparent mods : EHR, CRBoost, Aether Code implant (dmgTakenIncrease)

// DM Listeners : OnBeforeHitAll = add CR, onAfterAttack = Retarget, choose 1, 100% base hitrate AetherCode,
// OnBeforeSkillUse = set turn cooldown to 1, OnAfterSkillUse = apply AC if not yet have it.
// => set turn cd to 0.
// DmgTakenUp = OnStack, modifier, AllDmgTypeTakenRatio, stacking = replace, OnStart = AddModifier Main

// Quick Imp Plan :
// Modifiers :
// - Crit Rate : Listener = OnBeforeHitAll, check each enemy if it has 3 debuffs => add critrate boost
// - AetherCode : Listener = OnAfterAction, choose 1 among hit targets on last atk, AetherCode 100% BC
//	 -> check first if target already have AC
//	=> OnBeforeAction = set cd to 1
//  => set AetherCode as enemy modifier DmgTakenUp.

func init() {
	lightcone.Register(key.IncessantRain, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {

}
