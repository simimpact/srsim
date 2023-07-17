package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	Talent                              = "natasha-talent"
	talentHpThresholdPercentage float64 = 0.3
)

func init() {
	modifier.Register(Talent,
		modifier.Config{
			CanModifySnapshot: true,
			Listeners: modifier.Listeners{
				OnBeforeDealHeal: talentHealListener,
			},
		},
	)
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Talent,
		Source: c.id,
	})
}

// Listens to any heals done by nat to see if the target qualifies for the talent boost, including HOT
func talentHealListener(mod *modifier.Instance, e *event.HealStart) {
	if e.Target.CurrentHPRatio() <= talentHpThresholdPercentage {
		healer, _ := mod.Engine().CharacterInfo(e.Healer.ID())
		e.Healer.AddProperty(Talent, prop.HealBoost, talent[healer.TalentLevelIndex()])
	}
}
