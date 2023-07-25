package gepard

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	A2 key.Modifier = "gepard-a2"
	A6 key.Modifier = "gepard-a6"
)

func init() {
	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: func(mod *modifier.Instance) {
				mod.SetProperty(prop.ATKConvert, mod.Engine().Stats(mod.Owner()).DEF()*0.35)
			},
		},
	})
}

func (c *char) initTraces() {
	// A2
	if c.info.Traces["101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
			Stats:  info.PropMap{prop.AggroPercent: 3},
		})
	}

	// A6
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}
