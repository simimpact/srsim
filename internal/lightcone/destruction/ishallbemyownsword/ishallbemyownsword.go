package ishallbemyownsword

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
	OwnSword           key.Modifier = "i-shall-be-my-own-sword"
	Eclipse            key.Modifier = "i-shall-be-my-own-sword-eclipse"
	EclipseAllyMonitor key.Modifier = "i-shall-be-my-own-sword-eclipse-ally-monitor"
	EclipseDmgBonus                 = "i-shall-be-my-own-sword-dmg-bonus-buff"
	EclipseDefIgnore                = "i-shall-be-my-own-sword-def-ignore-buff"
)

type state struct {
	flag      bool
	dmgBonus  float64
	defIgnore float64
}

// ncreases the wearer's CRIT DMG by 20/23/26/29/32%.
// When an ally (excluding the wearer) gets attacked or loses HP, the wearer gains 1 stack of Eclipse, up to a max of 3 stack(s).
// Each stack of Eclipse increases the DMG of the wearer's next attack by 14/16.5/19/21.5/24%.
// When 3 stack(s) are reached, additionally enables that attack to ignore 12/14/16/18/20% of the enemy's DEF.
// This effect will be removed after the wearer uses an attack.

func init() {
	lightcone.Register(key.IShallBeMyOwnSword, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})

	modifier.Register(OwnSword, modifier.Config{})

	modifier.Register(EclipseAllyMonitor, modifier.Config{
		Listeners: modifier.Listeners{
			OnHPChange:           allyOnHPChange,
			OnAfterBeingAttacked: onAfterBeingAttacked,
		},
	})

	modifier.Register(Eclipse, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		CountAddWhenStack: 1,
		CanModifySnapshot: true,
		Listeners: modifier.Listeners{
			OnBeforeAttack: setFlag,
			OnBeforeHitAll: applyEclipse,
			OnAfterAttack:  removeEclipse,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	cdmgBuff := 0.17 + 0.03*float64(lc.Imposition)
	dmgBonus := 0.115 + 0.025*float64(lc.Imposition)
	defIgnore := 0.1 + 0.02*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   OwnSword,
		Source: owner,
		Stats:  info.PropMap{prop.CritDMG: cdmgBuff},
	})

	// apply modifier to all other allies
	engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		for char := range event.CharInfo {
			if char != owner {
				engine.AddModifier(char, info.Modifier{
					Name:   EclipseAllyMonitor,
					Source: owner,
					State: &state{
						dmgBonus:  dmgBonus,
						defIgnore: defIgnore,
						flag:      false,
					},
				})
			}
		}
	})
}

func allyOnHPChange(mod *modifier.Instance, e event.HPChange) {
	if !e.IsHPChangeByDamage && e.NewHP < e.OldHP {
		addStack(mod)
	}
}

func onAfterBeingAttacked(mod *modifier.Instance, e event.AttackEnd) {
	addStack(mod)
}

// helper function to handle stacks: adds a new Eclipse modifier,
// handing over values from the Monitor mod to the Eclipse mod
func addStack(mod *modifier.Instance) {
	st := mod.State().(*state)
	mod.Engine().AddModifier(mod.Source(), info.Modifier{
		Name:   Eclipse,
		Source: mod.Owner(),
		State: &state{
			dmgBonus:  st.dmgBonus,
			defIgnore: st.defIgnore,
			flag:      st.flag,
		},
	})
}

// set flag that makes sure to only apply Eclipse on an attack
func setFlag(mod *modifier.Instance, e event.AttackStart) {
	st := mod.State().(*state)
	st.flag = true
}

// if flag, apply Eclipse buff(s)
func applyEclipse(mod *modifier.Instance, e event.HitStart) {
	st := mod.State().(*state)
	if st.flag {
		e.Hit.Attacker.AddProperty(EclipseDmgBonus, prop.AllDamagePercent, mod.Count()*st.dmgBonus)
		if mod.Count() == 3 {
			e.Hit.Defender.AddProperty(EclipseDefIgnore, prop.DEFPercent, -st.defIgnore)
		}
	}
}

// remove mod after attack
func removeEclipse(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}
