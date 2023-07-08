package pela

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 key.Modifier = "pela-e1"
	E2 key.Modifier = "pela-e2"
	E4 key.Modifier = "pela-e4"
	E6 key.Modifier = "pela-e6"
)

func init() {
	// When an enemy is defeated, Pela regenerates 5 Energy.
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: func(mod *modifier.Instance, target key.TargetID) {
				if !mod.Engine().IsEnemy(target) {
					return
				}
				mod.Engine().ModifyEnergy(mod.Owner(), 5)
			},
		},
	})

	// Using Skill to remove buff(s) increases SPD by 10% for 2 turn(s).
	modifier.Register(E2, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})

	// When using Skill, there is a 100% base chance to reduce the target enemy's Ice RES by 12% for 2 turn(s).
	modifier.Register(E4, modifier.Config{
		TickMoment: modifier.ModifierPhase1End,
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
	})

	// When Pela attacks a debuffed enemy, she deals Additional Ice DMG equal to 40% of Pela's ATK to the enemy.
	modifier.Register(E6, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				for _, trg := range e.Targets {
					if mod.Engine().ModifierStatusCount(trg, model.StatusType_STATUS_DEBUFF) >= 1 {
						mod.Engine().Attack(info.Attack{
							Source:     mod.Owner(),
							Targets:    []key.TargetID{trg},
							DamageType: model.DamageType_ICE,
							AttackType: model.AttackType_PURSUED,
							BaseDamage: info.DamageMap{
								model.DamageFormula_BY_ATK: 0.4,
							},
						})
					}
				}
			},
		},
	})
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     E2,
			Source:   c.id,
			Duration: 2,
			Stats:    info.PropMap{prop.SPDPercent: 0.1},
		})
	}
}

func (c *char) e4(target key.TargetID) {
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(target, info.Modifier{
			Name:     E4,
			Source:   c.id,
			Chance:   1,
			Duration: 2,
			Stats:    info.PropMap{prop.IceDamageRES: -0.12},
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

	if c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6,
			Source: c.id,
		})
	}
}
