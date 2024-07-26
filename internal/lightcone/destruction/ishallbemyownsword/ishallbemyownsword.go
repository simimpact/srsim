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
			OnBeforeAttack: onBeforeAttack,
			OnBeforeHitAll: onBeforeHitAll,
			OnAfterAttack:  onAfterAttack,
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
					State: state{
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

// helper function to handle stacks
func addStack(mod *modifier.Instance) {
	st := mod.State().(*state)
	mod.Engine().AddModifier(mod.Source(), info.Modifier{
		Name:   Eclipse,
		Source: mod.Owner(),
		State: state{
			dmgBonus:  st.dmgBonus,
			defIgnore: st.defIgnore,
			flag:      false,
		},
	})
}

// set flag
func onBeforeAttack(mod *modifier.Instance, e event.AttackStart) {
	st := mod.State().(*state)
	st.flag = true
}

// if flag, apply Eclipse buff(s)
func onBeforeHitAll(mod *modifier.Instance, e event.HitStart) {
	st := mod.State().(*state)
	if st.flag {
		e.Hit.Attacker.AddProperty(EclipseDmgBonus, prop.AllDamagePercent, mod.Count()*st.dmgBonus)
		if mod.Count() == 3 {
			e.Hit.Defender.AddProperty(EclipseDefIgnore, prop.DEFPercent, -st.defIgnore)
		}
	}
}

// remove mod after attack
func onAfterAttack(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}
