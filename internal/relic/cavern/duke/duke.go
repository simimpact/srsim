package duke

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	dukeFuaBuff     = "ashblazing-grand-duke-boost"
	dukeAttackCheck = "ashblazing-grand-duke-effect"
	dukeAttackBuff  = "ashblazing-grand-duke-attack-buff"
)

// 2-Pc: Increases the DMG dealt by follow-up attack by 20%.
// 4-Pc: When the wearer uses a follow-up attack, increases the wearer's ATK by 6%
// for every time the follow-up attack deals DMG. This effect can stack up to 8 time(s)
// and lasts for 3 turn(s). This effect is removed the next time the wearer uses a follow-up attack.

func init() {
	relic.Register(key.AshblazingGrandDuke, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount: 2,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   dukeFuaBuff,
						Source: owner,
					})
				},
			},
			{
				MinCount: 4,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   dukeAttackCheck,
						Source: owner,
					})
				},
			},
		},
	})

	modifier.Register(dukeFuaBuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: addFuaDamage,
		},
	})

	modifier.Register(dukeAttackCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: onBeforeHitBuff,
			OnBeforeAttack: onBeforeAttack,
		},
	})

	modifier.Register(dukeAttackBuff, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          8,
		CountAddWhenStack: 1,
		Duration: 3,
		Listeners: modifier.Listeners{
			OnAdd: onAddStack,
		},
	})
}

// Increases the DMG dealt by follow-up attack by 20%.
func addFuaDamage(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_INSERT {
		e.Hit.Attacker.AddProperty(dukeFuaBuff, prop.AllDamagePercent, 0.2)
	}
}

// When the wearer uses a follow-up attack, increases the wearer's ATK by 6%
// for every time the follow-up attack deals DMG. This effect can stack up to 8 time(s)
// and lasts for 3 turn(s). This effect is removed the next time the wearer uses a follow-up attack.
func onBeforeHitBuff(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_INSERT {
		if mod.Engine().HasModifierFromSource(e.Attacker, mod.Owner(), dukeAttackBuff) {
			if !(mod.Engine().ModifierStackCount(e.Attacker, mod.Owner(), dukeAttackBuff) == 8) {
				e.Hit.Attacker.AddProperty(dukeAttackBuff, prop.ATKPercent, mod.Engine().ModifierStackCount(e.Attacker, mod.Owner(), dukeAttackBuff)*0.06)
				mod.Engine().AddModifier(e.Attacker, info.Modifier{
					Name: dukeAttackBuff,
					Source: mod.Owner(),
				})
			}
		} else {
			mod.Engine().AddModifier(e.Attacker, info.Modifier{
				Name:   dukeAttackBuff,
				Source: mod.Owner(),
			})
			e.Hit.Attacker.AddProperty(dukeAttackBuff, prop.ATKPercent, mod.Engine().ModifierStackCount(e.Attacker, mod.Owner(), dukeAttackBuff)*0.06)
		}
	}
}

func onBeforeAttack(mod *modifier.Instance, e event.AttackStart) {
	if e.AttackType == model.AttackType_INSERT {
		if mod.Engine().HasModifierFromSource(e.Attacker, mod.Owner(), dukeAttackBuff) {
			mod.OwnerStats().AddProperty(dukeAttackBuff, prop.ATKPercent, -mod.Engine().ModifierStackCount(e.Attacker, mod.Owner(), dukeAttackBuff)*0.06)
			mod.Engine().RemoveModifierFromSource(e.Attacker, mod.Owner(), dukeAttackBuff)
		}
	}
}

func onAddStack(mod *modifier.Instance) {
	mod.AddProperty(prop.ATKPercent, mod.Count()*0.06)
}
