package arlan

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1     key.Modifier = "arlan-e1"
	E4     key.Modifier = "arlan-e4"
	E6     key.Modifier = "arlan-e6"
	Revive key.Insert   = "arlan-revive"
)

func init() {
	// When HP is lower than or equal to 50% of Max HP, increases Skill's DMG by 10%.
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.AttackType != model.AttackType_SKILL {
					return
				}

				if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
					e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.1)
				}
			},
		},
	})

	// When struck by a killing blow after entering battle, instead of becoming knocked down,
	// Arlan immediately restores his HP to 25% of his Max HP. This effect is automatically
	// removed after it is triggered once or after 2 turn(s) have elapsed.
	modifier.Register(E4, modifier.Config{
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnLimboWaitHeal: func(mod *modifier.Instance) bool {
				// Dispel all debuffs
				mod.Engine().DispelStatus(mod.Owner(), info.Dispel{
					Status: model.StatusType_STATUS_DEBUFF,
					Order:  model.DispelOrder_LAST_ADDED,
				})

				// Queue Heal
				mod.Engine().InsertAbility(info.Insert{
					Execute: func() {
						mod.Engine().SetHP(
							mod.Owner(), mod.Owner(), mod.OwnerStats().MaxHP()*0.25)
					},
					Key:        Revive,
					Source:     mod.Owner(),
					Priority:   info.CharReviveSelf,
					AbortFlags: nil,
				})

				mod.RemoveSelf()
				return true
			},
		},
	})

	// When HP drops to 50% or below, Ultimate deals 20% more DMG. The DMG multiplier of
	// DMG taken by the target enemy now applies to adjacent enemies as well.
	modifier.Register(E6, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.AttackType != model.AttackType_ULT {
					return
				}

				if mod.Engine().HPRatio(mod.Owner()) <= 0.5 {
					e.Hit.Attacker.AddProperty(prop.AllDamagePercent, 0.2)
				}
			},
		},
	})
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.DispelStatus(c.id, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E1,
			Source: c.id,
		})
	}

	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
		})
	}

	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}
