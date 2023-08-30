package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill = "bailu-skill"
)

// Heals a single ally for 11.7% of Bailu's Max HP plus 312.
// Bailu then heals random allies 2 time(s). After each healing,
// HP restored from the next healing is reduced by 15%.

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	// record heal amt for next 2 random heals.
	healPercent := skillPercent[c.info.SkillLevelIndex()]
	healFlat := skillFlat[c.info.SkillLevelIndex()]
	// main targeted heal
	c.engine.Heal(info.Heal{
		Key:     Skill,
		Source:  c.id,
		Targets: []key.TargetID{target},
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: healPercent,
		},
		HealValue: healFlat,
	})

	// 2 randomized heals
	for i := 0; i < 2; i++ {
		// reduce heal amt for each subsequent heals
		healPercent = 0.85 * healPercent
		healFlat = 0.85 * healFlat
		chosenTarget := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Characters(),
			Max:     1,
		})
		c.engine.Heal(info.Heal{
			Key:     Skill,
			Source:  c.id,
			Targets: chosenTarget,
			BaseHeal: info.HealMap{
				model.HealFormula_BY_HEALER_MAX_HP: healPercent,
			},
			HealValue: healFlat,
		})
	}
}
