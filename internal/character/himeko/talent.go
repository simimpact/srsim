package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	himeko_talent = "himeko-talent"
)

func init() {
	modifier.Register(himeko_talent, modifier.Config{
		Listeners: modifier.Listeners{},
	})
}

func (c *char) initTalent() {
	c.engine.Events().StanceBreak.Subscribe(c.breakListener)
	c.engine.Events().AttackEnd.Subscribe(c.attackListener)
	//c.engine.Events().ModifierRemoved.Subscribe()

}

func (c *char) breakListener(e event.StanceBreak) {
	targ, _ := c.engine.EnemyInfo(e.Target)

	// Himeko talent immediately maxes out if an elite/boss is broken
	if targ.Rank == model.EnemyRank_BIG_BOSS || targ.Rank == model.EnemyRank_ELITE || targ.Rank == model.EnemyRank_LITTLE_BOSS {
		c.talentStacks += 3
	} else if c.info.Eidolon >= 4 && e.Source == c.id && e.Key == skill { // slightly hacky check for the skill
		c.talentStacks += 2
	} else {
		c.talentStacks++
	}

	if c.talentStacks >= 3 {

	}
}

func (c *char) attackListener(e event.AttackEnd) {
	if c.engine.IsCharacter(e.Attacker) && c.canAttack {
		if len(c.engine.Enemies()) != 0 {
			c.insertTalentAttack(e.Targets)
		}
	}
}

var talentHitSplit = []float64{0.2, 0.2, 0.2, 0.4}

func (c *char) insertTalentAttack(targets []key.TargetID) {
	c.engine.InsertAbility(info.Insert{
		Key: himeko_talent,
		Execute: func() {
			c.executeTalentAttack(targets)
		},
		Source: c.id,
		AbortFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_CTRL,
			model.BehaviorFlag_DISABLE_ACTION,
		},
		Priority: info.CharInsertAttackSelf,
	})
}

func (c *char) executeTalentAttack(targets []key.TargetID) {
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     e1,
			Source:   c.id,
			Duration: 2,
			Stats: info.PropMap{
				prop.SPDPercent: 0.2,
			},
		})
	}

	for index, ratio := range talentHitSplit {
		c.engine.Attack(info.Attack{
			Key:        himeko_talent,
			HitIndex:   index,
			HitRatio:   ratio,
			Targets:    targets,
			Source:     c.id,
			AttackType: model.AttackType_INSERT,
			DamageType: model.DamageType_FIRE,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()],
			},
			StanceDamage: 30,
			EnergyGain:   10,
		})
	}

	c.talentStacks = 0
}
