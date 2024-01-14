package kafka

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	kafkaTalent   = "kafka-talent"
	kafkaFollowup = "kafka-followup"
)

func init() {

}

func (c *char) initTalent() {
	c.engine.Events().AttackEnd.Subscribe(c.talentTrigger)
}

var talentHitSplit = []float64{0.15, 0.15, 0.15, 0.15, 0.15, 0.25}

func (c *char) talentTrigger(e event.AttackEnd) {
	canAttack := c.engine.IsCharacter(e.Attacker) && e.AttackType == model.AttackType_NORMAL && e.Attacker != c.id && c.canUseTalent
	if canAttack {
		c.engine.InsertAbility(info.Insert{
			Key: kafkaTalent,
			Execute: func() {
				for index, hit := range talentHitSplit {
					c.engine.Attack(info.Attack{
						Key:      kafkaFollowup,
						HitIndex: index,
						HitRatio: hit,
						//
						Targets: e.Targets,
					})
				}
			},
			Source:     c.id,
			AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
			Priority:   info.CharInsertAttackSelf,
		})
	}
}
