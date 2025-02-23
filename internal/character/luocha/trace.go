package luocha

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const A6 = "luocha-a6"

func init() {
	modifier.Register(A6, modifier.Config{})
}

func (c *char) initTraces() {
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:      A6,
			Source:    c.id,
			DebuffRES: info.DebuffRESMap{model.BehaviorFlag_STAT_CTRL: 0.7},
		})
	}
}
