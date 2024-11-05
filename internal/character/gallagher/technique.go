package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, enemy := range c.engine.Enemies() {
		c.engine.AddModifier(enemy, info.Modifier{
			Name: Besotted,
			State: &BesottedState{
				a6Active:  c.info.Traces["103"],
				breakVuln: talent[c.info.TalentLevelIndex()],
				healAmt:   talent_heal[c.info.TalentLevelIndex()],
			},
			// Technique is always duration of 2
			Duration: 2,
		})
	}
}
