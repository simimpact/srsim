package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Modifier = "natasha-skill"
)

func init() {

	//Nat HOT from skill
	modifier.Register(
		Skill,
		modifier.Config{
			Duration:          2,
			Stacking:          modifier.Unique,
			StatusType:        model.StatusType_STATUS_BUFF,
			CanModifySnapshot: true,
			Listeners: modifier.Listeners{
				OnPhase1: func(mod *modifier.Instance) {
					char, _ := mod.Engine().CharacterInfo(mod.Source())
					mod.Engine().Heal(info.Heal{
						Source:    mod.Source(),
						Targets:   []key.TargetID{mod.Owner()},
						BaseHeal:  info.HealMap{model.HealFormula_BY_HEALER_MAX_HP: skillContinuous[char.SkillLevelIndex()]},
						HealValue: skillContinuousFlat[char.SkillLevelIndex()],
					},
					)
				},
			},
		},
	)
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {

	healTarget := []key.TargetID{target}

	natasha := c.id

	healPercentage := info.HealMap{
		model.HealFormula_BY_HEALER_MAX_HP: skill[c.info.SkillLevelIndex()],
	}

	//Stats of the heal
	heal := info.Heal{
		Targets:     healTarget,
		Source:      natasha,
		BaseHeal:    healPercentage,
		HealValue:   skillFlatHeal[c.info.SkillLevelIndex()],
		UseSnapshot: true,
	}

	//The actual act of healing
	c.engine.Heal(heal)

	//Create the continuous heal modification
	c.engine.AddModifier(target, info.Modifier{
		Name:     Skill,
		Source:   c.id,
		Duration: 2,
	})
}
