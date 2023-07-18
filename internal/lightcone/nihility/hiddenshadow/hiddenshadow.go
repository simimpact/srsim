package hiddenshadow

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	shadow = "hidden-shadow"
)

type state struct {
	procced bool
	dmgMult float64
}

// DESC : After using Skill, the wearer's next Basic ATK deals
// Additional DMG equal to 60% of ATK to the target enemy.

// DM :
// OnAfterSkillUse : if skill, add _Sub mod
// _Sub def : OnAfterAttack = if flag = 1, retarget Max 1, includeLimbo.
// -> deal extra pursued dmg, param x% of lc holder's ATK.
// OnBeforeSkillUse : set flag to 1.
// OnAfterSkillUse : delete flag definition + modifier(wtf?)
// OnStart : add _Main mod.

// Impl NOTE :
// - no need to use subs. dmg param = lc holder atk.

func init() {
	lightcone.Register(key.HiddenShadow, lightcone.Config{
		CreatePassive: Create,
		Rarity:        3,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(shadow, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAction: skillProcced,
			OnAfterAttack: applyDamageOnBasic,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	modState := state{
		procced: false,
		dmgMult: 0.45 + 0.15*float64(lc.Imposition),
	}
	engine.AddModifier(owner, info.Modifier{
		Name:   shadow,
		Source: owner,
		State:  &modState,
	})
}

// if skill is used, set procced to true
func skillProcced(mod *modifier.Instance, e event.ActionEnd) {
	state := mod.State().(*state)
	if e.AttackType == model.AttackType_SKILL {
		state.procced = true
	}
}

// if attack is basic atk, apply extra dmg for 1 hit enemy randomly. set procced to false.
func applyDamageOnBasic(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*state)
	// return early if not basic atk or not procced
	if e.AttackType != model.AttackType_NORMAL || !state.procced {
		return
	}
	target := mod.Engine().Retarget(info.Retarget{
		Targets:      e.Targets,
		Max:          1,
		IncludeLimbo: true,
	})
	holderInfo, _ := mod.Engine().CharacterInfo(mod.Source())
	mod.Engine().Attack(info.Attack{
		Key:        shadow,
		Targets:    target,
		Source:     mod.Source(),
		AttackType: model.AttackType_PURSUED,
		DamageType: holderInfo.Element,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: state.dmgMult,
		},
		UseSnapshot: true,
	})
	// reset proc
	state.procced = false
}
