package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TalentCheck key.Modifier = "silverwolf-talent"
	BugATK      key.Modifier = "silverwolf-bug-atk"
	BugDEF      key.Modifier = "silverwolf-bug-def"
	BugSPD      key.Modifier = "silverwolf-bug-speed"
)

type talentState struct {
	penAmt float64
	cd     int
}

func init() {
	modifier.Register(BugATK, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				char, _ := mod.Engine().CharacterInfo(mod.Source())
				mod.SetProperty(prop.ATKPercent, -talentATK[char.AbilityLevel.Talent-1])
			},
		},
	})

	modifier.Register(BugDEF, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_DEF_DOWN},
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				char, _ := mod.Engine().CharacterInfo(mod.Source())
				mod.SetProperty(prop.DEFPercent, -talentDEF[char.AbilityLevel.Talent-1])
			},
		},
	})

	modifier.Register(BugSPD, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_DOWN},
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.ModifierInstance) {
				char, _ := mod.Engine().CharacterInfo(mod.Source())
				mod.SetProperty(prop.SPDPercent, -talentSPD[char.AbilityLevel.Talent-1])
			},
		},
	})

	modifier.Register(TalentCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.ModifierInstance, e event.AttackEndEvent) {
				char, _ := mod.Engine().CharacterInfo(e.Attacker)
				for _, trg := range e.Targets {
					bugs := []key.Modifier{}
					// get list of bugs not present on target
					for _, b := range []key.Modifier{BugATK, BugDEF, BugSPD} {
						if len(mod.Engine().GetModifiers(trg, b)) == 0 {
							bugs = append(bugs, b)
						}
					}
					mod.Engine().AddModifier(trg, info.Modifier{
						Name:     bugs[mod.Engine().Rand().Intn(len(bugs))],
						Source:   e.Attacker,
						Duration: 3,
						Chance:   talentChance[char.AbilityLevel.Talent-1],
					})
				}
			},
		},
	})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   TalentCheck,
		Source: c.id,
	})
}
