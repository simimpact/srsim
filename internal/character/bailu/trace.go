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
		Stacking: modifier.ReplaceBySource,
	})
}

func (c *char) initTraces() {
	// A2 : When Bailu heals a target ally above their normal Max HP,
	//      the target's Max HP increases by 10% for 2 turns.
	// TODO : implement OnSnapshotCreate logic. look into natasha talent impl for example.
	c.engine.Events().HealEnd.Subscribe(func(e event.HealEnd) {
		if !c.info.Traces["101"] {
			return
		}
		if e.Healer == c.id && e.OverflowHealAmount > 0 {
			c.engine.AddModifier(e.Target, info.Modifier{
				Name:     A2,
				Source:   c.id,
				Duration: 2,
				Stats:    info.PropMap{prop.HPPercent: 0.1},
			})
		}
	})
}
