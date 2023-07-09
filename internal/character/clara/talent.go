package clara

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TalentMark key.Modifier = "clara-talent-mark"    // MAvatar_Klara_00_PassiveATK_Mark for BPSkill_Revenge
	TalentRes  key.Modifier = "clara-talent-dmg-res" // MAvatar_Klara_00_Passive_DamageReduce
)

// Under the protection of Svarog, DMG taken by Clara when hit by enemy
// attacks is reduced by 10%. Svarog will mark enemies who attack Clara with
// his Mark of Counter and retaliate with a Counter, dealing Physical DMG
// equal to *% of Clara's ATK.
func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   TalentRes,
		Source: c.id,
		Stats:  info.PropMap{prop.AllDamageReduce: 0.1},
	})
}

// MAvatar_Klara_00_PassiveATK_Mark
//
// listens to AttackEnd events and check if clara is eligible for counter
func (c *char) talentActionEndListener(e event.AttackEnd) {
	attackerID := e.Attacker

	// won't counter non-enemies (dominated ally)
	if !c.engine.IsEnemy(attackerID) {
		return
	}

	// check for mark actually disregards the 50% fixed chance at E6, meaning
	// at E6 it's guaranteed any ally is hit, hence we can't use this in
	// canCounter block
	if c.canMark(e) {
		// add marker modifier on enemy attacker
		c.engine.AddModifier(attackerID, info.Modifier{
			Name:   TalentMark,
			Source: c.id,
		})
	}

	// canCounter (clara targeted or e6 + an ally winning 50/50)
	if c.canCounter(e) {
		c.doCounter(attackerID)
	}
}

// executes the follow-up attack,
// NOTE: follow-up attack is same for both talent and e6
func (c *char) doCounter(attackerID key.TargetID) {
	c.engine.InsertAbility(info.Insert{
		Execute: func() {
			hasUlt := c.engine.HasModifier(c.id, UltCounter)

			percent := talent[c.info.TalentLevelIndex()]
			if hasUlt {
				percent += ultDmgBoost[c.info.UltLevelIndex()]
			}

			// normal counter, damage
			c.engine.Attack(info.Attack{
				Source:       c.id,
				Targets:      []key.TargetID{attackerID},
				DamageType:   model.DamageType_PHYSICAL,
				AttackType:   model.AttackType_INSERT,
				BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: percent},
				StanceDamage: 30.0,
				EnergyGain:   5.0,
			})

			// enhanced counter, damage to adjacent
			// NOTE: in-game wording = damage value is based on main target's def
			// mhy memes ?
			if hasUlt {
				c.engine.Attack(info.Attack{
					Source:       c.id,
					Targets:      c.engine.AdjacentTo(attackerID),
					DamageType:   model.DamageType_PHYSICAL,
					AttackType:   model.AttackType_INSERT,
					BaseDamage:   info.DamageMap{model.DamageFormula_BY_ATK: percent * 0.5},
					StanceDamage: 30.0,
				})

				// counter done, decrease counter stack
				c.engine.ExtendModifierCount(c.id, UltCounter, -1.0)
			}
		},
		Source:     c.id,
		Priority:   info.CharInsertAttackSelf,
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
	})
}

func (c *char) canCounter(e event.AttackEnd) bool {
	// clara has ult counter stacks left
	hasUlt := c.engine.HasModifier(c.id, UltCounter)

	// attack targets clara's ally, 50/50 roll for each ally to counter
	for _, target := range e.Targets {
		if !c.engine.IsCharacter(target) {
			// only select allies
			continue
		}
		// attack targets clara, guarantees a counter
		isClara := target == c.id

		// attack targets ally and winning the 50% fixed
		allyRoll := c.info.Eidolon >= 6 && c.engine.Rand().Float32() < 0.5

		if isClara || hasUlt || allyRoll {
			return true
		}
	}
	return false
}

// helper method to decide if it's eligible to add talent marker, being either:
// 1. clara is included in the attack
// 2. E6 and any ally is attacked
func (c *char) canMark(e event.AttackEnd) bool {
	for _, target := range e.Targets {
		// clara or E6 + ally
		if c.id == target || (c.info.Eidolon >= 6 && c.engine.IsCharacter(target)) {
			return true
		}
	}
	return false
}
