package huohuo

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	TechDebuff key.Modifier = "huohuo-technique"
)

func init() {
	modifier.Register(TechDebuff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_DEBUFF,
	})
}

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, target := range c.engine.Enemies() {
		c.engine.AddModifier(target, info.Modifier{
			Name:     TechDebuff,
			Source:   c.id,
			Chance:   1,
			Duration: 2,
			Stats:    info.PropMap{prop.ATKPercent: -0.25},
		})
	}
}
