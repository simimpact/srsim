package herta

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent = "herta-talent"
)

func (c *char) initTalent() {
	c.engine.Events().HPChange.Subscribe(c.talentListener)
}

var (
	hertaCountInsert = 0
	hertaCount       = 0
	hertaCountATK    = 0
	passiveCooldown  = 0
)

func (c *char) talentListener(e event.HPChange) {
	// if herta insert count = 1, set herta atk count to 0, herta insert count to 0
	if hertaCountInsert == 1 {
		hertaCountATK = 0
		hertaCountInsert = 0
	}

	if e.NewHPRatio <= 0.5 {
		if c.engine.IsEnemy(e.Target) && passiveCooldown == 0 && !c.engine.HasBehaviorFlag(c.id, model.BehaviorFlag_STAT_CTRL) {
			if len(c.engine.Enemies()) > 0 {
				c.engine.Events().AttackEnd.Subscribe(c.talentAfterAttackListener)
				hertaCount += 1
			}
		}
	}
}

func (c *char) talentAfterAttackListener(e event.AttackEnd) {
	if hertaCountATK == 0 && hertaCount > 0 && c.engine.IsCharacter(e.Attacker) && !c.passiveFlag {
		if len(c.engine.Enemies()) > 0 {
			hertaCountATK = 1
			hertaCountInsert = 1
			c.engine.InsertAbility(info.Insert{
				Source: c.id,
				AbortFlags: []model.BehaviorFlag{
					model.BehaviorFlag_STAT_CTRL,
					model.BehaviorFlag_DISABLE_ACTION,
				},
				Key:      Talent,
				Execute:  c.talentInsert,
				Priority: info.CharInsertAttackSelf,
			})
		}
	}
}

func (c *char) talentInsert() {
	c.passiveFlag = true
	hertaCountInsert = 0
	for hertaCount > 0 && len(c.engine.Enemies()) > 0 {
		hertaCount -= 1

		if c.info.Eidolon >= 2 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   e2,
				Source: c.id,
			})
		}

		if c.info.Eidolon >= 4 {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   e4,
				Source: c.id,
				Stats: info.PropMap{
					prop.AllDamagePercent: 0.1,
				},
			})
		}

		c.talentInsertAttack()
	}

	c.engine.RemoveModifier(c.id, e4)
	c.engine.EndAttack()

	c.passiveFlag = false
}

func (c *char) talentInsertAttack() {
	c.engine.Attack(info.Attack{
		Key:        Talent,
		Source:     c.id,
		Targets:    c.engine.Enemies(),
		AttackType: model.AttackType_INSERT,
		DamageType: model.DamageType_ICE,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()],
		},
		EnergyGain:   5,
		StanceDamage: 15,
	})
}
