package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent             = "march7th-talent"
	TalentCount        = "march7th-talentcount"
	MarchAllyMark      = "march7th-shield-counter"
	MarchCounterMark   = "march7th-counter-mark"
	MarchCounterAttack = "march7th-counter"
)

/*
Note:
High level breakdown of how March's talent works, as far as I can tell from the DM:
When March joins batttle, she is given a modifier that represents her actual talent.
This modifier, when added, adds another modifier to all of March's teammates, which I called MarchAllyMark
MarchAllyMark listens for whenever the person it is attached to is about to be attacked. If at least one of the targets in
that attack is shielded, AND the attacker is an enenmy, add a modifier called MarchCounterMark to the attacker
MarchCounterMark listens for when whoever it is attached to is done attacking.
Afterwards, it tells March to counter attack. If successful, March will counter, and (?) remove
one stack from the modifier that determines if march can counter or not
MarchCounterMark will not be added to an enemy if March does not have enough counters left.
March regenerates her counters on Phase2End.
*/

func init() {
	modifier.Register(MarchCounterMark, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: talentCounterAttack,
		},
	})

	modifier.Register(MarchAllyMark, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingAttacked: checkToApplyCounterMark,
		},
	})

	modifier.Register(TalentCount, modifier.Config{
		Count:    2,
		MaxCount: 3,
	})

	// The actual talent modifier
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				for _, teammate := range mod.Engine().Characters() {
					mod.Engine().AddModifier(teammate, info.Modifier{
						Name:   MarchAllyMark,
						Source: mod.Owner(),
					})
				}
			},
			OnPhase2: func(mod *modifier.Instance) {
				march7th, _ := mod.Engine().CharacterInfo(mod.Source())
				talentCount := 2
				if march7th.Eidolon >= 4 {
					talentCount = 3
				}
				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name:   TalentCount,
					Source: mod.Source(),
					Count:  float64(talentCount),
				})
				// TODO : Implement logic for keeping track of how many counters march has left, note below
			},
			OnBeforeDying: func(mod *modifier.Instance) {
				for _, teammate := range mod.Engine().Characters() {
					mod.Engine().RemoveModifier(teammate, MarchAllyMark)
				}
			},
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})
}

func checkToApplyCounterMark(mod *modifier.Instance, e event.AttackStart) {
	hasEnoughCounterStacks := mod.Engine().HasModifier(mod.Source(), TalentCount)
	qualifiesForCounter := mod.Engine().IsShielded(mod.Owner()) && mod.Engine().IsEnemy(e.Attacker) && hasEnoughCounterStacks
	//TODO: Add support for checking how many counters March's talent has left:
	/*
	   Thinking there are two main ways to do this:
	   First would be to have a state struct that keeps track of how many counters March has already shot. Do not think
	   this is a good way since there's no easy way to pass this information to each of the instances of the modifier march gives to her                teammates.

	   Other better way I am leaning towards is to have March's talent modifier itself keep track of how many counters she has remaining                via the count/stacks, or the count/stacks of another modifier it creates/adds to March every Phase2 (Similar to DM)
	*/

	if qualifiesForCounter {
		mod.Engine().AddModifier(e.Attacker, info.Modifier{
			Name:   MarchCounterMark,
			Source: mod.Source(),
		})
	}
}

// Actual counter attack logic
func talentCounterAttack(mod *modifier.Instance, e event.AttackEnd) {
	// Counter attack
	mod.Engine().InsertAbility(info.Insert{
		Source:   mod.Source(),
		Priority: info.CharInsertAttackSelf,
		Key:      MarchCounterAttack,
		Execute: func() {
			march7th, _ := mod.Engine().CharacterInfo(mod.Source())
			mod.Engine().Attack(info.Attack{
				Key:        MarchCounterAttack,
				Source:     mod.Source(),
				Targets:    []key.TargetID{mod.Owner()},
				AttackType: model.AttackType_INSERT,
				DamageType: model.DamageType_ICE,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: talent[march7th.TalentLevelIndex()],
					model.DamageFormula_BY_DEF: 0.3,
				},
			})
			// Remove this modifier from the enemy it is attached to
			mod.Engine().RemoveModifier(mod.Owner(), MarchCounterMark)
			// Probably not right
			mod.Engine().ExtendModifierCount(mod.Source(), TalentCount, -1)
		},
		AbortFlags: []model.BehaviorFlag{
			model.BehaviorFlag_DISABLE_ACTION,
			model.BehaviorFlag_STAT_CTRL,
		},
	})
}
