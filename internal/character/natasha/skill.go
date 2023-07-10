package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill key.Modifier = "natasha-skill"
)

func init() {
	modifier.Register(
		Skill,
		modifier.Config{
			Duration:          2,
			Stacking:          modifier.Unique,
			StatusType:        model.StatusType_STATUS_BUFF,
			CanModifySnapshot: true,
			Listeners: modifier.Listeners{
				OnPhase1: func(mod *modifier.Instance) {
					mod.Engine().Heal(info.Heal{
						Source:   mod.Source(),
						Targets:  []key.TargetID{mod.Owner()},
						BaseHeal: info.HealMap{model.HealFormula_BY_HEALER_MAX_HP: skillContinuous[mod.Source()]},
					})
				}}})
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {

	healTargets := []key.TargetID{target}

	source := c.id

	BaseHeal := info.HealMap{
		model.HealFormula_BY_HEALER_MAX_HP: skill[state.CharacterInfo().SkillLevelIndex()],
	}

	//Stats of the heal
	heal := info.Heal{
		Targets:     healTargets,
		Source:      source,
		BaseHeal:    BaseHeal,
		HealValue:   skillFlatHeal[state.CharacterInfo().SkillLevelIndex()],
		UseSnapshot: true,
	}

	//Check for A2
	//Dispel a debuff
	/*c.engine.DispelStatus(target, info.Dispel{
		Status: model.StatusType_STATUS_DEBUFF,
		Order:  model.DispelOrder_LAST_ADDED,
		Count:  1,
	})*/

	//The actual act of healing
	c.engine.Heal(heal)

	natMaxHp := c.engine.Stats(c.id).MaxHP()
	continuousHealPercentage := skillContinuous[state.CharacterInfo().SkillLevelIndex()]
	continuousFlatAmount := skillContinuousFlat[state.CharacterInfo().SkillLevelIndex()]
	continuousHealAmt := continuousHealPercentage*natMaxHp + continuousFlatAmount

	//The continuous heal applied
	continuousHeal := info.Modifier{
		Name:     Skill,
		Source:   c.id,
		Stats:    info.PropMap{prop.HealTaken: continuousHealAmt},
		Duration: 2,
	}

	//Create the continuous heal modification
	c.engine.AddModifier(target, continuousHeal)
}
