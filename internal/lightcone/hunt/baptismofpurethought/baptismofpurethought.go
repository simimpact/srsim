package baptismofpurethought

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
	Check       = "baptism-of-pure-thought"
	Disputation = "baptism-of-pure-thought-disputation-buff"
)

type state struct {
	cdmgBonusInc float64
	dmgBonus     float64
	defignore    float64
}

// Increases the wearer's CRIT DMG by 20/23/26/29/32%. For every debuff on the enemy target,
// the wearer's CRIT DMG dealt against this target additionally increases by 8/9/10/11/12%, stacking up to 3 times.
// When using Ultimate to attack the enemy target, the wearer receives the Disputation effect,
// which increases DMG dealt by 36/42/48/54/60% and enables their follow-up attacks to ignore 24/28/32/36/40% of the target's DEF.
// This effect lasts for 2 turns.

func init() {
	lightcone.Register(key.BaptismofPureThought, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyCdmgInc,
			OnBeforeAttack: applyDisputation,
		},
	})

	modifier.Register(Disputation, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		// CanDispel: true,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyDefignore,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	cdmgAmt := 0.17 + 0.03*float64(lc.Imposition)
	cdmgIncAmt := 0.07 + 0.01*float64(lc.Imposition)
	dmgAmt := 0.3 + 0.06*float64(lc.Imposition)
	defignoreAmt := 0.2 + 0.04*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.CritDMG: cdmgAmt},
		State: &state{
			cdmgBonusInc: cdmgIncAmt,
			dmgBonus:     dmgAmt,
			defignore:    defignoreAmt,
		},
	})
}

func applyCdmgInc(mod *modifier.Instance, e event.HitStart) {
	debuffCount := mod.Engine().ModifierStatusCount(e.Defender, model.StatusType_STATUS_DEBUFF)
	if debuffCount > 3 {
		debuffCount = 3
	}
	cdmgBonusInc := mod.State().(state).cdmgBonusInc
	e.Hit.Attacker.AddProperty(Check, prop.CritDMG, float64(debuffCount)*cdmgBonusInc)
}

func applyDisputation(mod *modifier.Instance, e event.AttackStart) {
	// this might need to be a SkillType/ActionType check in the future
	if e.AttackType == model.AttackType_ULT {
		st := mod.State().(state)
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   Disputation,
			Source: mod.Owner(),
			Stats:  info.PropMap{prop.AllDamagePercent: st.dmgBonus},
			State:  st,
		})
	}
}

func applyDefignore(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_INSERT {
		e.Hit.Defender.AddProperty(Disputation, prop.DEFPercent, -mod.State().(state).defignore)
	}
}
