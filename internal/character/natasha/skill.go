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
	// Nat HOT from skill
	modifier.Register(
		Skill,
		modifier.Config{
			Duration:          2,
			Stacking:          modifier.ReplaceBySource,
			StatusType:        model.StatusType_STATUS_BUFF,
			CanModifySnapshot: true,
			Listeners: modifier.Listeners{
				OnPhase1: natHot,
			},
			TickMoment: modifier.ModifierPhase1End,
		},
	)
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// Nat dispel (Checks if nat is A2)
	if c.info.Traces["1101101"] {
		c.engine.DispelStatus(target, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	// Stats of the heal
	heal := info.Heal{
		Targets: []key.TargetID{target},
		Source:  c.id,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: skill[c.info.SkillLevelIndex()],
		},
		HealValue:   skillFlatHeal[c.info.SkillLevelIndex()],
		UseSnapshot: true,
	}

	// The actual act of healing
	c.engine.Heal(heal)

	hotDuration := 2

	// A6
	if c.info.Traces["1101103"] {
		hotDuration = 3
	}

	// Create the continuous heal modification
	c.engine.AddModifier(target, info.Modifier{
		Name:     Skill,
		Source:   c.id,
		Duration: hotDuration,
	})

	c.engine.ModifyEnergy(c.id, 30)
}

func natHot(mod *modifier.Instance) {
	char, _ := mod.Engine().CharacterInfo(mod.Source())
	mod.Engine().Heal(info.Heal{
		Source:    mod.Source(),
		Targets:   []key.TargetID{mod.Owner()},
		BaseHeal:  info.HealMap{model.HealFormula_BY_HEALER_MAX_HP: skillContinuous[char.SkillLevelIndex()]},
		HealValue: skillContinuousFlat[char.SkillLevelIndex()],
	},
	)
}
