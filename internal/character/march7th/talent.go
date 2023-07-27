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

func init() {
	modifier.Register(MarchCounterMark, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				mod.Engine().InsertAbility(info.Insert{
					Key: MarchCounterAttack,
					Execute: func() {
						mod.Engine().Attack(info.Attack{
							Source:     mod.Source(),
							Targets:    []key.TargetID{mod.Owner()},
							AttackType: model.AttackType_INSERT,
							DamageType: model.DamageType_ICE,
						})
					},
					AbortFlags: []model.BehaviorFlag{
						model.BehaviorFlag_DISABLE_ACTION,
						model.BehaviorFlag_STAT_CTRL,
					},
				})
			},
		},
	})

	modifier.Register(MarchAllyMark, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingAttacked: func(mod *modifier.Instance, e event.AttackStart) {

				/*oneTargetWasShielded := false
				for _, target := range e.Targets {
					//oneTargetWasShielded = mod.Engine().HasShield(target, key.Shield()) || oneTargetWasShielded
				}*/
			},
		},
	})

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
				mod.Engine().AddModifier(mod.Owner(), info.Modifier{
					Name: TalentCount,
				})

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
