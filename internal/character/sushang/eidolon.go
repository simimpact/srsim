package sushang

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 key.Modifier = "sushang-e2"
	E4 key.Modifier = "sushang-e4"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4, modifier.Config{})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Source: c.id,
			Stats: info.PropMap{
				prop.BreakEffect: 0.4,
			},
		})
	}

	if c.info.Eidolon >= 6 {
		c.addTalentBuff()
	}
}

func (c *char) e2() {
	if c.info.Eidolon >= 2 {
		if !c.engine.HasModifier(c.id, E2) {
			c.engine.AddModifier(c.id, info.Modifier{
				Name:   E2,
				Source: c.id,
				Stats: info.PropMap{
					prop.AllDamageReduce: 0.2,
				},
				Duration: 1,
			})
		}
	}
}
