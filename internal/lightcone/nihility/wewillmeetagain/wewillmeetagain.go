package wewillmeetagain

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
	meet = "we-will-meet-again"
)

type state struct {
	proceed     bool
	extraDmgAmt float64
}

// After the wearer uses Basic ATK or Skill, deals Additional DMG
// equal to 48% of the wearer's ATK to a random enemy that has been attacked.

func init() {
	lightcone.Register(key.WeWillMeetAgain, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(meet, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAction: activateTrigger,
			OnAfterAttack:  addExtraDmgOnTrigger,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	extraDmgAmt := 0.36 + 0.12*float64(lc.Imposition)
	modState := state{
		proceed:     false,
		extraDmgAmt: extraDmgAmt,
	}

	engine.AddModifier(owner, info.Modifier{
		Name:   meet,
		Source: owner,
		State:  &modState,
	})
}

func activateTrigger(mod *modifier.Instance, e event.ActionStart) {
	state := mod.State().(*state)
	// if action is basic atk/skill, set proceed to true. else set to false.
	state.proceed = e.AttackType == model.AttackType_NORMAL ||
		e.AttackType == model.AttackType_SKILL
}

func addExtraDmgOnTrigger(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*state)
	// if proceed, add pursued atk to 1 random attacked enemy.
	if !state.proceed {
		return
	}
	chosenTarget := mod.Engine().Retarget(info.Retarget{
		Targets: e.Targets,
		Max:     1,
	})

	mod.Engine().Attack(info.Attack{
		Key:        meet,
		Targets:    chosenTarget,
		Source:     mod.Owner(),
		AttackType: model.AttackType_PURSUED,
		DamageType: e.DamageType,
		BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: state.extraDmgAmt},
	})
	// reset proc value
	state.proceed = false
}
