package themoleswelcomeyou

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
	Check key.Modifier = "the-moles-welcome-you"
	Buff  key.Modifier = "mischievous"
)

type checkState struct {
	attacks map[model.AttackType]struct{}
	atkBuff float64
}

// When the wearer uses Basic ATK, Skill, or Ultimate to attack enemies, the wearer gains
// one stack of Mischievous. Each stack increases the wearer's ATK by 12%/15%/18%/21%/24%.
func init() {
	lightcone.Register(key.TheMolesWelcomeYou, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: onAfterAttack,
		},
	})

	modifier.Register(Buff, modifier.Config{
		MaxCount:          3,
		CountAddWhenStack: 1,
		Stacking:          modifier.ReplaceBySource,
		StatusType:        model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: buffOnAdd,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		State: checkState{
			attacks: make(map[model.AttackType]struct{}, 3),
			atkBuff: 0.09 + 0.03*float64(lc.Imposition),
		},
	})
}

// after an attack, add 1 stack iff new attack type
func onAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(checkState)

	// must be normal, skill, or ult
	if e.AttackType != model.AttackType_NORMAL &&
		e.AttackType != model.AttackType_SKILL &&
		e.AttackType != model.AttackType_ULT {
		return
	}

	if _, has := state.attacks[e.AttackType]; !has {
		state.attacks[e.AttackType] = struct{}{}
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   Buff,
			Source: mod.Owner(),
			State:  state.atkBuff,
		})
	}
}

// each stack increases by amt
func buffOnAdd(mod *modifier.Instance) {
	amt := mod.State().(float64)
	mod.AddProperty(prop.ATKPercent, amt*mod.Count())
}
