package clara

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TalentCounter key.Modifier = "clara-talent-counter"
	TalentMark    key.Modifier = "clara-talent-mark"    // MAvatar_Klara_00_PassiveATK_Mark for BPSkill_Revenge
	TalentRes     key.Modifier = "clara-talent-dmg-res" // MAvatar_Klara_00_Passive_DamageReduce
)

// Under the protection of Svarog, DMG taken by Clara when hit by enemy
// attacks is reduced by 10%. Svarog will mark enemies who attack Clara with
// his Mark of Counter and retaliate with a Counter, dealing Physical DMG
// equal to *% of Clara's ATK.
//
// flow:
//
// onBeforeBeingAttacked: + TalentMark, + TalentRes
//
// onAfterAttack: + follow-up attack
func init() {
	modifier.Register(TalentCounter, modifier.Config{
		Listeners: modifier.Listeners{
			// add marker modifier on attacker enemy
			OnBeforeBeingAttacked: func(mod *modifier.Instance, e event.AttackStart) {
				mod.Engine().AddModifier(e.Attacker, info.Modifier{
					Name:   TalentMark,
					Source: mod.Source(),
				})
			},

			// remove marker modifier on all enemies when clara dies
			OnBeforeDying: func(mod *modifier.Instance) {
				for _, enemyID := range mod.Engine().Enemies() {
					mod.Engine().RemoveModifier(enemyID, TalentMark)
				}
			},
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   TalentCounter,
		Source: c.id,
	})

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

	// canCounter (clara targeted or e6 + an ally winning 50/50)
	if c.canCounter(e) {
		// add marker modifier on enemy attacker
		c.engine.AddModifier(attackerID, info.Modifier{
			Name:   TalentMark,
			Source: c.id,
			State:  State{skillLevelIndex: c.info.SkillLevelIndex(), ultLevelIndex: c.info.UltLevelIndex()},
		})

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
	hasUlt := c.engine.HasModifier(c.id, UltCounter) && c.engine.ModifierStackCount(c.id, c.id, UltCounter) > 0

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
