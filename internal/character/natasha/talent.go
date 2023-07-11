package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Talent                      key.Modifier = "natasha-talent"
	talentHpThresholdPercentage float64      = 0.3
)

func init() {
	modifier.Register(Talent,
		modifier.Config{
			StatusType:        model.StatusType_STATUS_BUFF,
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
		e.Healer.AddProperty(prop.HealBoost, 0.5)
	}
}
