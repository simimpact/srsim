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

type counterState struct {
	countersLeft   int
	counterScaling float64
}

/*
Note to self:
High level breakdown of how March's talent works, as far as I can tell from the DM:
When March joins batttle, she is given a modifier that represents her actual talent.
This modifier, when added, adds another modifier to all of March's teammates, which I called MarchAllyMark
MarchAllyMark listens for whenever the person it is attached to is about to be attacked. If at least one of the targets in
that attack is shielded, AND the attacker is an enenmy, add a modifier called MarchCounterMark to the attacker
MarchCounterMark listens for when whoever it is attached to is done attacking.
Afterwards, it tells March to counter attack. If successful, March will counter
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

	//The actual talent modifier
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
				/*mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name: TalentCount,
				})*/
				//TODO : Implement logic for keeping track of how many counters march has left, note below

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

func (c *char) talentCounterListener(e event.CharactersAdded) {

}

func checkToApplyCounterMark(mod *modifier.Instance, e event.AttackStart) {
	//Check all targets of any attack an enemy makes on someone holding the MarchAllyMark
	atLeastOneTargetShielded := false
	for _, target := range e.Targets {
		atLeastOneTargetShielded = mod.Engine().IsShielded(target) || atLeastOneTargetShielded
	}

	//TODO: Add support for checking how many counters March's talent has left:
	/*
		Thinking there are two main ways to do this:
		First would be to have a state struct that keeps track of how many counters March has already shot. Do not think
		this is a good way since there's no easy way to pass this information to each of the instances of the modifier march gives to her
		teammates.

		Other better way I am leaning towards is to have March's talent modifier itself keep track of how many counters she has remaining
		via the count/stacks, or the count/stacks of another modifier it creates/adds to March every Phase2 (Similar to DM)
	*/

	if atLeastOneTargetShielded && mod.Engine().IsEnemy(e.Attacker) {
		mod.Engine().AddModifier(e.Attacker, info.Modifier{
			Name:   MarchCounterMark,
			Source: mod.Source(),
		})
	}
}

// Actual counter attack logic
func talentCounterAttack(mod *modifier.Instance, e event.AttackEnd) {
	//Counter attack
	mod.Engine().InsertAbility(info.Insert{
		Key: MarchCounterAttack,
		Execute: func() {
			mod.Engine().Attack(info.Attack{
				Source:     mod.Source(),
				Targets:    []key.TargetID{mod.Owner()},
				AttackType: model.AttackType_INSERT,
				DamageType: model.DamageType_ICE,
			})
			//Remove this modifier from the enemy it is attached to
			mod.Engine().RemoveModifier(mod.Owner(), MarchCounterMark)
		},
		AbortFlags: []model.BehaviorFlag{
			model.BehaviorFlag_DISABLE_ACTION,
			model.BehaviorFlag_STAT_CTRL,
		},
	})
}
