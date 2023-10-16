package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult key.Heal     = "bailu-ult"
	E2  key.Modifier = "bailu-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:          modifier.Replace,
		StatusType:        model.StatusType_STATUS_BUFF,
		CanModifySnapshot: true,
	})
}

// Heals all allies for 13.5% of Bailu's Max HP plus 360.
// Bailu applies Invigoration to allies that are not already Invigorated.
// For those already Invigorated, Bailu extends the duration of their Invigoration by 1 turn.
// The effect of Invigoration can last for 2 turn(s). This effect cannot stack.

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	healPercent := ultPercent[c.info.UltLevelIndex()]
	healFlat := ultFlat[c.info.UltLevelIndex()]

	// use retarget to filter out characters on limbo
	filteredTargets := c.engine.Retarget(info.Retarget{
		Targets: c.engine.Characters(),
	})
	// main team heal
	for _, char := range filteredTargets {
		c.addHeal(Ult, healPercent, healFlat, []key.TargetID{char})
	}

	// add team invigoration, already invigorated get extended duration by 1.
	for _, char := range c.engine.Characters() {
		if c.engine.HasModifier(char, invigoration) {
			// if target char already has invigoration, extend duration.
			c.engine.ExtendModifierDuration(char, invigoration, 1)
		} else {
			c.addInvigoration(char, 2)
		}
	}

	// E2 : After using her Ultimate, Bailu's Outgoing Healing increases
	// by an additional 15% for 2 turn(s).
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     E2,
			Source:   c.id,
			Stats:    info.PropMap{prop.HealBoost: 0.15},
			Duration: 2,
		})
	}
}
