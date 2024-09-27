package worrisomeblissful

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
	check = "worrisome-blissful"
	tame  = "worrisome-blissful-cdmg-debuff"
)

type state struct {
	fuaflag   bool
	dmgBonus  float64
	cdmgBonus float64
}

// Increase the wearer's CRIT Rate by 18/21/24/27/30% and increases DMG dealt by follow-up attack by 30/35/40/45/50%.
// After the wearer uses a follow-up attack, inflicts the target with the Tame state, stacking up to 2 time(s).
// When allies hit enemy targets under the Tame state, each Tame stack increases the CRIT DMG dealt by 12/14/16/18/20%.

func init() {
	lightcone.Register(key.WorrisomeBlissful, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: buffFuaDmg,
			OnAfterAttack:  applyTame,
		},
	})

	modifier.Register(tame, modifier.Config{
		StatusType:        model.StatusType_STATUS_DEBUFF,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          2,
		CountAddWhenStack: 1,
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll: buffCdmg,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	crAmt := 0.15 + 0.03*float64(lc.Imposition)
	dmgAmt := 0.25 + 0.05*float64(lc.Imposition)
	cdmgAmt := 0.1 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   check,
		Source: owner,
		Stats: info.PropMap{
			prop.CritChance: crAmt,
		},
		State: &state{
			fuaflag:   false,
			dmgBonus:  dmgAmt,
			cdmgBonus: cdmgAmt,
		},
	})
}

func buffFuaDmg(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_INSERT {
		st := mod.State().(*state)
		e.Hit.Attacker.AddProperty(check, prop.AllDamagePercent, st.dmgBonus)
		// flag for checking whether to apply Tame debuff
		st.fuaflag = true
	}
}

func applyTame(mod *modifier.Instance, e event.AttackEnd) {
	st := mod.State().(*state)
	if st.fuaflag {
		for _, trg := range e.Targets {
			mod.Engine().AddModifier(trg, info.Modifier{
				Name:   tame,
				Source: mod.Owner(),
				State: state{
					fuaflag:   st.fuaflag,
					dmgBonus:  st.dmgBonus,
					cdmgBonus: st.cdmgBonus,
				},
			})
		}
		st.fuaflag = false
	}
}

func buffCdmg(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().IsCharacter(e.Attacker) {
		cdmgBonus := mod.State().(*state).cdmgBonus
		e.Hit.Attacker.AddProperty(tame, prop.CritDMG, mod.Count()*cdmgBonus)
	}
}
