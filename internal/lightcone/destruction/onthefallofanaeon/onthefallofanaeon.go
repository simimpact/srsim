package onthefallofanaeon

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
	Check        key.Modifier = "on-the-fall-of-an-aeon"
	BuffAtk      key.Modifier = "on-the-fall-of-an-aeon-atk-buff"
	BuffDmgBonus key.Modifier = "on-the-fall-of-an-aeon-dmg-bonus-buff"
)

type state struct {
	atkBuff  float64
	dmgBonus float64
}

// Whenever the wearer attacks, their ATK is increased by 8/10/12/14/16% in this battle.
// This effect can stack up to 4 time(s).
// When the wearer inflicts Weakness Break on enemies, the wearer's DMG increases by 12/15/18/21/24% for 2 turn(s).
func init() {
	lightcone.Register(key.OntheFallofanAeon, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAttack: onBeforeAttack,
			OnTriggerBreak: onTriggerBreak,
		},
	})

	modifier.Register(BuffAtk, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          4,
		CountAddWhenStack: 1,
		Listeners: modifier.Listeners{
			OnAdd: onAdd,
		},
	})

	modifier.Register(BuffDmgBonus, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		State: state{
			atkBuff:  0.06 + 0.02*float64(lc.Imposition),
			dmgBonus: 0.09 + 0.03*float64(lc.Imposition),
		},
	})
}

func onBeforeAttack(mod *modifier.ModifierInstance, e event.AttackStartEvent) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   BuffAtk,
		Source: mod.Owner(),
		State:  mod.State().(state).atkBuff,
	})
}

func onAdd(mod *modifier.ModifierInstance) {
	atkBuff := mod.State().(state).atkBuff
	mod.AddProperty(prop.ATKPercent, atkBuff*mod.Count())
}

func onTriggerBreak(mod *modifier.ModifierInstance, target key.TargetID) {
	dmgBonus := mod.State().(state).dmgBonus
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   BuffDmgBonus,
		Source: mod.Owner(),
		Stats:  info.PropMap{prop.AllDamagePercent: dmgBonus},
	})
}
