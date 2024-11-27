package serval

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A4:
//
//	At the start of the battle, immediately regenerates 15 Energy.
//
// A6:
//
//	Upon defeating an enemy, ATK is increased by 20% for 2 turn(s).
const (
	A4      = "serval-a4"
	A6      = "serval-a6"
	A6Check = "serval-a6-check"
)

func init() {
	modifier.Register(A6Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: a6Buff,
		},
	})
	modifier.Register(A6, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
	})
}

func (c *char) initTraces() {
	c.engine.Events().BattleStart.Subscribe(func(event event.BattleStart) {
		if c.info.Traces["102"] {
			c.engine.ModifyEnergy(info.ModifyAttribute{
				Key:    A4,
				Target: c.id,
				Source: c.id,
				Amount: 15.0,
			})
		}
	})

	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6Check,
			Source: c.id,
		})
	}
}

func a6Buff(mod *modifier.Instance, target key.TargetID) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:     A6,
		Source:   mod.Owner(),
		Duration: 2,
		Stats:    info.PropMap{prop.ATKPercent: 0.2},
	})
}
