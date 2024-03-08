package silverwolf

import (
	"github.com/simimpact/srsim/pkg/engine"
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

func init() {
	modifier.Register(BugATK, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				char, _ := mod.Engine().CharacterInfo(mod.Source())
				mod.SetProperty(prop.ATKPercent, -talentATK[char.TalentLevelIndex()])
			},
		},
	})

	modifier.Register(BugDEF, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_DEF_DOWN},
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				char, _ := mod.Engine().CharacterInfo(mod.Source())
				mod.SetProperty(prop.DEFPercent, -talentDEF[char.TalentLevelIndex()])
			},
		},
	})

	modifier.Register(BugSPD, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_DOWN},
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				char, _ := mod.Engine().CharacterInfo(mod.Source())
				mod.SetProperty(prop.SPDPercent, -talentSPD[char.TalentLevelIndex()])
			},
		},
	})

	modifier.Register(TalentCheck, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: func(mod *modifier.Instance, e event.AttackEnd) {
				if len(e.Targets) == 0 {
					return
				}
				mod.Engine().AddModifier(e.Targets[0], newRandomBug(mod.Engine(), e.Targets[0], e.Attacker))
			},
		},
	})
}

func newRandomBug(engine engine.Engine, target, source key.TargetID) info.Modifier {
	char, _ := engine.CharacterInfo(source)
	bugs := []key.Modifier{}
	// get list of bugs not present on target
	for _, b := range []key.Modifier{BugATK, BugDEF, BugSPD} {
		if !engine.HasModifier(target, b) {
			bugs = append(bugs, b)
		}
	}
	// if all bugs on target, application is random
	if len(bugs) == 0 {
		bugs = []key.Modifier{BugATK, BugDEF, BugSPD}
	}
	duration := 3
	if char.Traces["101"] {
		duration += 1
	}
	return info.Modifier{
		Name:     bugs[engine.Rand().Intn(len(bugs))],
		Source:   source,
		Duration: duration,
		Chance:   talentChance[char.TalentLevelIndex()],
	}
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   TalentCheck,
		Source: c.id,
	})
}
