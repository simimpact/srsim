package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
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
	c.addHeal(Skill, healPercent, healFlat, []key.TargetID{target})

	// 2 randomized heals
	for i := 0; i < 2; i++ {
		// reduce heal amt for each subsequent random heals
		healPercent = 0.85 * healPercent
		healFlat = 0.85 * healFlat
		chosenTarget := c.engine.Retarget(info.Retarget{
			Targets: c.engine.Characters(),
			Max:     1,
		})
		c.addHeal(Skill, healPercent, healFlat, chosenTarget)
	}

	// energy gained after skill usage
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Skill,
		Source: c.id,
		Target: c.id,
		Amount: 30.0,
	})
}
