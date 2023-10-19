package bailu

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	A2 key.Modifier = "bailu-a2"
)

func init() {
	// register here
	modifier.Register(A2, modifier.Config{
		Stacking:          modifier.ReplaceBySource,
		CanModifySnapshot: true,
	})
}

func (c *char) initTraces() {
	// A2 : When Bailu heals a target ally above their normal Max HP,
	//      the target's Max HP increases by 10% for 2 turns.
	c.engine.Events().HealEnd.Subscribe(func(e event.HealEnd) {
		// return early if no A2, healer isn't bailu, or overflow heal amt is 0 or less.
		if !c.info.Traces["101"] || e.Healer != c.id || e.OverflowHealAmount <= 0 {
			return
		}
		c.engine.AddModifier(e.Target, info.Modifier{
			Name:     A2,
			Source:   c.id,
			Duration: 2,
			Stats:    info.PropMap{prop.HPPercent: 0.1},
		})
	})
}
