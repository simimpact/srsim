package kafka

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	kafkaTalent = "kafka-talent"
	followup    = "kafka-followup"
)

func init() {
	modifier.Register(kafkaTalent, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase2: restoreTalent,
		},
	})
}

func restoreTalent(mod *modifier.Instance) {
	kafka, _ := mod.Engine().CharacterInstance(mod.Owner())
	c := kafka.(*char)
	c.canUseTalent = true
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   kafkaTalent,
		Source: c.id,
	})
	c.engine.Events().AttackEnd.Subscribe(c.talentTrigger)
	c.engine.Events().ActionEnd.Subscribe(c.talentAttack)
}

var talentHitSplit = []float64{0.15, 0.15, 0.15, 0.15, 0.15, 0.25}
var talentTargs = []key.TargetID{}

func (c *char) talentTrigger(e event.AttackEnd) {
	isAlly := c.engine.IsCharacter(e.Attacker)
	isBasicAtk := e.AttackType == model.AttackType_NORMAL
	isNotKafka := e.Attacker != c.id
	if isAlly && isBasicAtk && isNotKafka && c.canUseTalent {
		talentTargs = e.Targets
	}
}

func (c *char) talentAttack(e event.ActionEnd) {
	isAlly := c.engine.IsCharacter(e.Owner)
	isBasicAtk := e.AttackType == model.AttackType_NORMAL
	isNotKafka := e.Owner != c.id
	target := talentTargs[0]

	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(target, info.Modifier{
			Name:     E1,
			Source:   c.id,
			Chance:   1,
			Duration: 2,
		})
	}

	if isBasicAtk && c.canUseTalent && c.engine.HPRatio(target) <= 0 && isNotKafka && isAlly {
		c.engine.InsertAbility(info.Insert{
			Key: followup,
			Execute: func() {
				for index, hit := range talentHitSplit {
					c.engine.Attack(info.Attack{
						Key:          followup,
						HitIndex:     index,
						HitRatio:     hit,
						Source:       c.id,
						Targets:      []key.TargetID{target},
						EnergyGain:   10,
						StanceDamage: 30,
						DamageType:   model.DamageType_THUNDER,
						AttackType:   model.AttackType_INSERT,
						BaseDamage: info.DamageMap{
							model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()],
						},
					})
				}
			},
			Source:     c.id,
			AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
			Priority:   info.CharInsertAttackSelf,
		})
	}

	c.applyShock([]key.TargetID{target})

	c.canUseTalent = false
}
