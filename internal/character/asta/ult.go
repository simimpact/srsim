package asta

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ult = "asta-ult"
)

func init() {
	modifier.Register(ult, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
		Duration:      2,
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for _, ally := range c.engine.Characters() {
		c.engine.AddModifier(ally, info.Modifier{
			Name:     ult,
			Source:   c.id,
			Duration: 2,
			Stats: info.PropMap{
				prop.SPDPercent: ultimate[c.info.UltLevelIndex()],
			},
		})
	}

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    ult,
		Source: c.id,
		Target: c.id,
		Amount: 5,
	})

	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:     e2,
			Source:   c.id,
			Duration: 1,
		})
	}
}
