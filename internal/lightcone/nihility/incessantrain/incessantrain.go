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
	rain             key.Modifier = "incessant-rain"
	aetherCodeApply  key.Modifier = "incessant-rain-aether-code-applier"
	aetherCodeDebuff key.Modifier = "incessant-rain-aether-code-debuff"
)

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
		},
	})
	modifier.Register(aetherCodeApply, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: applyDebuffOnce,
		},
	})
	modifier.Register(aetherCodeDebuff, modifier.Config{
		StatusType: model.StatusType_STATUS_DEBUFF,
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
	// Aether Code applier
	dmgTakenAmt := 0.10 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   aetherCodeApply,
		Source: owner,
		State:  dmgTakenAmt,
	})
}

// boost CR if enemy has >=3 debuffs
func critRateBoost(mod *modifier.Instance, e event.HitStart) {
	debuffCount := float64(e.Hit.Defender.StatusCount(model.StatusType_STATUS_DEBUFF))
	critRateAmt := mod.State().(float64)
	if debuffCount >= 3 {
		e.Hit.Defender.AddProperty(prop.CritChance, critRateAmt)
	}
}

// retarget with 1 chosen(among non-AC-applied). apply dmgTakenUp with 100% basechance. cooldown tick.
func applyDebuffOnce(mod *modifier.Instance, e event.AttackEnd) {
	// fetch enemy list hit by this attack
	enemyList := e.Targets
	dmgTakenAmt := mod.State().(float64)
	var validEnemyList []key.TargetID

	// validEnemyList should only contain non-dead, non-implanted enemies
	for _, enemy := range enemyList {
		// is enemy alive and does enemy not have aether code yet. if so, append.
		if mod.Engine().HPRatio(enemy) > 0 && !mod.Engine().HasModifier(enemy, aetherCodeDebuff) {
			validEnemyList = append(validEnemyList, enemy)
		}
	}
	if validEnemyList != nil {
		// choose one enemy, apply debuff to them.
		chosenOne := validEnemyList[mod.Engine().Rand().Intn(len(validEnemyList))]
		mod.Engine().AddModifier(chosenOne, info.Modifier{
			Name:     aetherCodeDebuff,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.AllDamageTaken: dmgTakenAmt},
			Chance:   1.0,
			Duration: 1,
		})
	}
}
