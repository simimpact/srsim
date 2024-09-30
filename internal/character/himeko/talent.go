package himeko

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	himekoTalent = "himeko-talent"
)

func (c *char) initTalent() {
	c.engine.Events().StanceBreak.Subscribe(c.talentBreakListener)
	c.engine.Events().AttackEnd.Subscribe(c.talentAttackListener)
	c.engine.Events().BattleStart.Subscribe(c.talentBattleStartListener)
	c.engine.Events().ModifierRemoved.Subscribe(c.talentModifierRemoveListener)
	c.engine.Events().EnemiesAdded.Subscribe(c.talentNewEnemiesListener)
}

func (c *char) talentBreakListener(e event.StanceBreak) {
	targ, _ := c.engine.EnemyInfo(e.Target)

	// Himeko talent immediately maxes out if an elite/boss is broken
	switch targ.Rank {
	case model.EnemyRank_BIG_BOSS:
		fallthrough
	case model.EnemyRank_ELITE:
		fallthrough
	case model.EnemyRank_LITTLE_BOSS:
		c.talentStacks += 3
	default:
		if e.Source == c.id && e.Key == skill {
			c.talentStacks += 2
		} else {
			c.talentStacks++
		}
	}
}

func (c *char) talentAttackListener(e event.AttackEnd) {
	if c.engine.IsCharacter(e.Attacker) && c.canAttack {
		// If we still have alive enemies
		if len(c.engine.Enemies()) > 0 {
			c.insertTalentAttack(e.Targets)
		}
	}
}

func (c *char) talentBattleStartListener(e event.BattleStart) {
	if c.canAttack {
		if c.talentStacks >= 3 && len(c.engine.Enemies()) > 0 {
			c.insertTalentAttack(c.engine.Enemies())
		}
	}
}

func (c *char) talentModifierRemoveListener(e event.ModifierRemoved) {
	if c.engine.IsCharacter(e.Target) && c.canAttack {
		if c.talentStacks >= 3 && len(c.engine.Enemies()) != 0 {
			c.insertTalentAttack(c.engine.Enemies())
		}
	}
}

func (c *char) talentNewEnemiesListener(e event.EnemiesAdded) {
	if c.canAttack && c.talentStacks >= 3 {
		c.insertTalentAttack(c.engine.Enemies())
	}
}

func (c *char) insertTalentAttack(targets []key.TargetID) {
	c.engine.InsertAbility(info.Insert{
		Key: himekoTalent,
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

var talentHitSplit = []float64{0.2, 0.2, 0.2, 0.4}

/*
*

	Execute the actual attack part of the talent;
	Seperated for readability

*
*/
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
			Key:        himekoTalent,
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
