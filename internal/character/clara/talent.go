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
				for _, enemyId := range mod.Engine().Enemies() {
					mod.Engine().RemoveModifier(enemyId, TalentMark)
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
	attackerId := e.Attacker

	// won't counter non-enemies (dominated ally)
	if !c.engine.IsEnemy(attackerId) {
		return
	}

	if c.canCounter(e) {
		c.doCounter(attackerId)
	}
}

// executes the follow-up attack,
// NOTE: follow-up attack is same for both talent and e6
func (c *char) doCounter(attackerId key.TargetID) {
	c.engine.InsertAbility(info.Insert{
		Execute: func() {
			// DamageMaps
			normalDamage := info.DamageMap{
				model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()],
			}
			enhancedDamage := info.DamageMap{
				model.DamageFormula_BY_ATK: talent[c.info.TalentLevelIndex()] +
					ult_dmg_boost[c.info.UltLevelIndex()],
			}
			mainTargetDamage := normalDamage

			if c.engine.HasModifier(c.id, UltCounter) {
				mainTargetDamage = enhancedDamage
			}

			// normal counter, damage
			c.engine.Attack(info.Attack{
				Source:       c.id,
				Targets:      []key.TargetID{attackerId},
				DamageType:   model.DamageType_PHYSICAL,
				AttackType:   model.AttackType_INSERT,
				BaseDamage:   mainTargetDamage,
				StanceDamage: 30.0,
				EnergyGain:   5.0,
			})

			// enhanced counter, damage to adjacent
			// NOTE: in-game wording = damage value is based on main target's def
			// mhy memes ?
			if c.engine.HasModifier(c.id, UltCounter) {
				splashDamage := info.DamageMap{model.DamageFormula_BY_ATK: (talent[c.info.TalentLevelIndex()] + ult_dmg_boost[c.info.UltLevelIndex()]) * 0.5}

				c.engine.Attack(info.Attack{
					Source:       c.id,
					Targets:      c.engine.AdjacentTo(attackerId),
					DamageType:   model.DamageType_PHYSICAL,
					AttackType:   model.AttackType_INSERT,
					BaseDamage:   splashDamage,
					StanceDamage: 30.0,
				})
			}
		},
		Source:     c.id,
		Priority:   info.CharInsertAttackSelf,
		AbortFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_CTRL, model.BehaviorFlag_DISABLE_ACTION},
	})

	// add marker modifier on enemy attacker
	c.engine.AddModifier(attackerId, info.Modifier{
		Name:   TalentMark,
		Source: c.id,
		State:  State{skillLevelIndex: c.info.SkillLevelIndex()},
	})
}

func enemyBeingAttacked(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(State)
	// clara is the source of TalentMark
	clara := mod.Source()
	// attacked by clara's skill
	if e.AttackType == model.AttackType_SKILL && e.Attacker == clara {
		mod.Engine().Attack(info.Attack{
			Source:     clara,
			Targets:    []key.TargetID{mod.Owner()},
			DamageType: model.DamageType_PHYSICAL,
			AttackType: model.AttackType_PURSUED,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: skill[state.skillLevelIndex],
			},
		})
	}
}

// helper fn to see if a counter is eligible
//
// - if clara is not targeted AND < E6 -> no counter
//
// - clara targeted OR E6 + winning E6 50/50 -> counter
func (c *char) canCounter(e event.AttackEnd) bool {
	isClara := false
	for _, target := range e.Targets {
		if target == c.id {
			isClara = true
		}
	}
	if isClara || (c.engine.HasModifier(c.id, E6) && c.engine.Rand().Float32() < 0.5) {
		return true
	}
	return false
}
