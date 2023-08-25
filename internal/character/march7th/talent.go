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

type talentState struct {
	maxCounters float64
}

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

	modifier.Register(TalentCount, modifier.Config{
		Stacking: modifier.Replace,
		Count:    2,
		MaxCount: 2,
	})

	// The actual talent modifier
	modifier.Register(Talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase2: func(mod *modifier.Instance) {
				//Perhaps keep the max count allowed stored in a state of the Talent modifier, to avoid having to grab with CharacterInfo
				maxCounters := mod.State().(talentState).maxCounters
				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name:     TalentCount,
					Source:   mod.Source(),
					Count:    maxCounters,
					MaxCount: maxCounters,
				})
			},
		},
	})
}

func (c *char) initTalent() {
	talentCount := 2.0
	if c.info.Eidolon >= 4 {
		talentCount += 1.0
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
		State: talentState{
			maxCounters: talentCount,
		},
	})
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     TalentCount,
		Source:   c.id,
		Count:    talentCount,
		MaxCount: talentCount,
	})
}

func (c *char) applyCounterMark(e event.AttackStart) {
	hasShieldedTarget := false
	for _, ally := range c.engine.Characters() {
		hasShieldedTarget = hasShieldedTarget || c.engine.IsShielded(ally)
	}

	haveCounters := c.engine.HasModifier(c.id, TalentCount)

	canCounter := hasShieldedTarget && c.engine.IsEnemy(e.Attacker) && haveCounters

	if canCounter {
		c.engine.AddModifier(e.Attacker, info.Modifier{
			Name:   MarchCounterMark,
			Source: c.id,
		})
	}

}

// Actual counter attack logic
func talentCounterAttack(mod *modifier.Instance, e event.AttackEnd) {
	mod.Engine().InsertAbility(info.Insert{
		Source:   mod.Source(),
		Priority: info.CharInsertAttackSelf,
		Key:      MarchCounterAttack,
		Execute: func() {
			march7th, _ := mod.Engine().CharacterInfo(mod.Source())
			e4Scaling := 0.0
			if march7th.Eidolon >= 4 {
				e4Scaling = 0.30
			}
			mod.Engine().Attack(info.Attack{
				Key:        MarchCounterAttack,
				Source:     mod.Source(),
				Targets:    []key.TargetID{mod.Owner()},
				AttackType: model.AttackType_INSERT,
				DamageType: model.DamageType_ICE,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: talent[march7th.TalentLevelIndex()],
					model.DamageFormula_BY_DEF: e4Scaling,
				},
			})
			// Remove a count/stack from the talent counter
			mod.Engine().ExtendModifierCount(mod.Source(), TalentCount, -1.0)
			// Remove this modifier from the enemy it is attached to
			mod.Engine().RemoveModifier(mod.Owner(), MarchCounterMark)
		},
		AbortFlags: []model.BehaviorFlag{
			model.BehaviorFlag_DISABLE_ACTION,
			model.BehaviorFlag_STAT_CTRL,
		},
	})
}
