package hanya

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ultimate = "hanya-ult"
)

func init() {
	modifier.Register(Ultimate, modifier.Config{
		StatusType:    model.StatusType_STATUS_BUFF,
		Stacking:      modifier.Replace,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	ultdur := 2
	if c.info.Eidolon >= 4 {
		ultdur = 3
	}
	hanyaStats := c.engine.Stats(c.id)
	c.engine.AddModifier(target, info.Modifier{
		Name:     Ultimate,
		Source:   c.id,
		Duration: ultdur,
		Stats: info.PropMap{
			prop.ATKPercent: ultAtkBuff[c.info.UltLevelIndex()],
			prop.SPDConvert: ultSpdBuff[c.info.UltLevelIndex()]*hanyaStats.SPD() - hanyaStats.GetProperty(prop.SPDConvert),
		},
	})

	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(target, info.Modifier{
			Name:     E1,
			Source:   c.id,
			Duration: ultdur,
		})
	}
}
