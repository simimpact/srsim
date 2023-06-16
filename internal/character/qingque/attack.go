package qingque

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	if c.tiles[0] != 4 {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_QUANTUM,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
			},
			StanceDamage: 30.0,
			EnergyGain:   20.0,
		})
		c.engine.RemoveModifier(c.id, Talent)
		if c.engine.HasModifier(c.id, Autarky) {
			c.engine.InsertAbility(info.Insert{
				Execute: func() {
					c.engine.Attack(info.Attack{
						Source:     c.id,
						Targets:    []key.TargetID{target},
						DamageType: model.DamageType_QUANTUM,
						AttackType: model.AttackType_INSERT,
						BaseDamage: info.DamageMap{
							model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
						},
						StanceDamage: 30.0,
						// Energy gain might change
						EnergyGain: 20.0,
					})
				},
				Source:     c.id,
				Priority:   info.CharInsertAttackOthers,
				AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
			})
			c.engine.RemoveModifier(c.id, Autarky)
		}
		c.engine.ModifySP(1)
	} else {
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    []key.TargetID{target},
			DamageType: model.DamageType_QUANTUM,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: 2.4 * atk[c.info.AttackLevelIndex()],
			},
			StanceDamage: 60.0,
			EnergyGain:   20.0,
		})
		c.engine.Attack(info.Attack{
			Source:     c.id,
			Targets:    c.engine.AdjacentTo(target),
			DamageType: model.DamageType_QUANTUM,
			AttackType: model.AttackType_NORMAL,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
			},
			StanceDamage: 30.0,
		})
		c.engine.RemoveModifier(c.id, Talent)
		c.a6()
		if c.engine.HasModifier(c.id, Autarky) {
			c.engine.InsertAbility(info.Insert{
				Execute: func() {
					c.engine.Attack(info.Attack{
						Source:     c.id,
						Targets:    []key.TargetID{target},
						DamageType: model.DamageType_QUANTUM,
						AttackType: model.AttackType_INSERT,
						BaseDamage: info.DamageMap{
							model.DamageFormula_BY_ATK: 2.4 * atk[c.info.AttackLevelIndex()],
						},
						StanceDamage: 60.0,
						// Energy gain might change
						EnergyGain: 20.0,
					})
					c.engine.Attack(info.Attack{
						Source:     c.id,
						Targets:    c.engine.AdjacentTo(target),
						DamageType: model.DamageType_QUANTUM,
						AttackType: model.AttackType_INSERT,
						BaseDamage: info.DamageMap{
							model.DamageFormula_BY_ATK: atk[c.info.AttackLevelIndex()],
						},
						StanceDamage: 30.0,
					})
				},
				Source:     c.id,
				Priority:   info.CharInsertAttackOthers,
				AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
			})
			c.engine.RemoveModifier(c.id, Autarky)
		}
		if c.info.Eidolon >= 6 {
			c.engine.ModifySP(1)
		}
		c.engine.RemoveModifier(c.id, Talent)
	}

	state.EndAttack()
}
