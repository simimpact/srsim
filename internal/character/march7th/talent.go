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
	MarchCounterMark   = "march7th-counter-mark"
	MarchCounterAttack = "march7th-counter"
)

type talentState struct {
	maxCounters float64
}

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
				// Perhaps keep the max count allowed stored in a state of the Talent modifier, to avoid having to grab with CharacterInfo
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
		talentCount += 1
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

	haveCounters := c.engine.ModifierStackCount(c.id, c.id, TalentCount) > 0

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
			mod.Engine().ExtendModifierCount(mod.Source(), TalentCount, -1)
		},
		AbortFlags: []model.BehaviorFlag{
			model.BehaviorFlag_DISABLE_ACTION,
			model.BehaviorFlag_STAT_CTRL,
		},
	})

	// Remove this modifier from the enemy it is attached to
	mod.Engine().RemoveModifier(mod.Owner(), MarchCounterMark)
}
