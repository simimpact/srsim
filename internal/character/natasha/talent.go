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
				OnBeforeDealHeal: func(mod *modifier.Instance, e *event.HealStart) {
					char, _ := mod.Engine().CharacterInfo(mod.Owner())
					if e.Target.CurrentHPRatio() <= talentHpThresholdPercentage {
						mod.AddProperty(prop.HealBoost, talent[char.TalentLevelIndex()])
					}
				},
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
