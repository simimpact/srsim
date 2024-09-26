package solitaryhealing

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
	check    = "solitary-healing-check"
	solitary = "solitary-healing"
)

// Increases the wearer's Break Effect by 20%. When the wearer uses their Ultimate,
// increases DoT dealt by the wearer by 24%, lasting for 2 turn(s).
// When a target enemy suffering from DoT imposed by the wearer is defeated,
// regenerates 4 Energy for the wearer.

func init() {
	lightcone.Register(key.SolitaryHealing, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: buffDotOnUlt,
		},
	})
	modifier.Register(solitary, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	// apply ult checker + static BE buff
	breakEffectAmt := 0.15 + 0.05*float64(lc.Imposition)
	dotBuffAmt := 0.18 + 0.06*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   check,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: breakEffectAmt},
		State:  dotBuffAmt,
	})

	// energy from killed enemy suffering owner's DOT
	// TODO : DM uses OnListenCharacterDie. confirm is equivalent to .onDeath subs.
	engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if !engine.HasBehaviorFlag(e.Target, model.BehaviorFlag_STAT_DOT) ||
			!engine.IsEnemy(e.Target) ||
			!engine.IsCharacter(e.Killer) {
			return
		}
		// add flat energy to owner
		energyAmt := 3.5 + 0.5*float64(lc.Imposition)
		engine.ModifyEnergy(info.ModifyAttribute{
			Key:    solitary,
			Source: owner,
			Target: owner,
			Amount: energyAmt,
		})
	})
}

func buffDotOnUlt(mod *modifier.Instance, e event.ActionStart) {
	if e.AttackType != model.AttackType_ULT {
		return
	}
	dotBuffAmt := mod.State().(float64)
	// buff DOTs for 2 turns
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     solitary,
		Source:   mod.Owner(),
		Stats:    info.PropMap{prop.DOTDamagePercent: dotBuffAmt},
		Duration: 2,
	})
}
