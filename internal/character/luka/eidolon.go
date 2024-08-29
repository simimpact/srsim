package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	e1 = "luka-e1"
	e2 = "luka-e2"
)

func (c *char) initEidolons() {

}

func init() {
	modifier.Register(e1, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		Count:      1,
		StatusType: model.StatusType_STATUS_BUFF,
	})

}

func (c *char) e1Check(target key.TargetID) {
	if c.info.Eidolon >= 1 && c.engine.HasBehaviorFlag(target, model.BehaviorFlag_STAT_DOT_BLEED) {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   e1,
			Source: c.id,
			Stats: info.PropMap{
				prop.AllDamagePercent: 0.15,
			},
			Duration: 2,
		})
	}
}
