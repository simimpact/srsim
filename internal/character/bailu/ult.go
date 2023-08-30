package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Ult          key.Heal     = "bailu-ult"
	invigoration key.Modifier = "invigoration"
)

// Heals all allies for 13.5% of Bailu's Max HP plus 360.
// Bailu applies Invigoration to allies that are not already Invigorated.
// For those already Invigorated, Bailu extends the duration of their Invigoration by 1 turn.
// The effect of Invigoration can last for 2 turn(s). This effect cannot stack.

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	healPercent := ultPercent[c.info.UltLevelIndex()]
	healFlat := ultFlat[c.info.UltLevelIndex()]

	// main team heal
	for _, char := range c.engine.Characters() {
		c.addHeal(Ult, healPercent, healFlat, []key.TargetID{char})
	}

	// add team invigoration, already invigorated get extended duration by 1.
	for _, char := range c.engine.Characters() {
		duration := 2
		if c.engine.HasModifier(char, invigoration) {
			duration = 1
		}
		c.engine.AddModifier(char, info.Modifier{
			Name:     invigoration,
			Source:   c.id,
			Duration: duration,
		})
	}
}
