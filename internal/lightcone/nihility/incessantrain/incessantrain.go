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
	rain key.Modifier = "incessant-rain"
	code key.Modifier = "aether-code"
)

type state struct {
	amt     float64
	targets []key.TargetID // set to nil in Create
}

// Desc : Increases the wearer's Effect Hit Rate by 24%.
// When the wearer deals DMG to an enemy that currently has 3 or more debuffs,
// increases the wearer's CRIT Rate by 12%. After the wearer uses their Basic ATK, Skill, or Ultimate,
// there is a 100% base chance to implant Aether Code on a random hit target that does not yet have it.
// Targets with Aether Code receive 12% increased DMG for 1 turn.

func init() {
	lightcone.Register(key.IncessantRain, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(rain, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: critRateBoost,
			OnAfterAttack:  fetchHitEnemies,
			OnAfterAction:  applyDebuffOnce,
		},
	})
	modifier.Register(code, modifier.Config{
		StatusType: model.StatusType_STATUS_DEBUFF,
		Stacking:   modifier.Replace,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// ehr constant + crit with condition
	ehrAmt := 0.20 + 0.04*float64(lc.Imposition)
	critAmt := 0.10 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   rain,
		Source: owner,
		Stats:  info.PropMap{prop.EffectHitRate: ehrAmt},
		State:  critAmt,
	})
	// Aether Code state setter
	dmgTakenAmt := 0.10 + 0.02*float64(lc.Imposition)
	state.amt = dmgTakenAmt
	state.targets = nil
}

// boost CR on current hit if enemy has >=3 debuffs
func critRateBoost(mod *modifier.Instance, e event.HitStart) {
	debuffCount := float64(e.Hit.Defender.StatusCount(model.StatusType_STATUS_DEBUFF))
	critRateAmt := mod.State().(float64)
	if debuffCount >= 3 {
		e.Hit.Defender.AddProperty(prop.CritChance, critRateAmt)
	}
}

// fetch the list of all enemies hit by this attack
func fetchHitEnemies(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*state)
	state.targets = e.Targets
}

// retarget with 1 chosen(among non-AC-applied). apply dmgTakenUp with 100% basechance. cooldown tick.
func applyDebuffOnce(mod *modifier.Instance, e event.ActionEnd) {
	state := mod.State().(*state)

	// make new array for possible targets, loop through state.targets to filter
	validEnemies := make([]key.TargetID, 0, len(state.targets))
	for _, t := range state.targets {
		if mod.Engine().HPRatio(t) > 0 && !mod.Engine().HasModifier(t, code) {
			validEnemies = append(validEnemies, t)
		}
	}

	if validEnemies != nil {
		// choose one enemy, apply debuff to them.
		chosenOne := validEnemies[mod.Engine().Rand().Intn(len(validEnemies))]
		mod.Engine().AddModifier(chosenOne, info.Modifier{
			Name:     code,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.AllDamageTaken: state.amt},
			Chance:   1.0,
			Duration: 1,
		})
	}
	// set targets to nil at end to reset
	state.targets = nil
}
