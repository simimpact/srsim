package natasha

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 key.Modifier = "natasha-a2"
	A4 key.Modifier = "natasha-a4"
	A6 key.Modifier = "natasha-a6"
)

func init() {

	//A2 dispel
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeBeingHeal: func(mod *modifier.Instance, e *event.HealStart) {
				mod.Engine().DispelStatus(e.Target.ID(), info.Dispel{
					Status: model.StatusType_STATUS_DEBUFF,
					Order:  model.DispelOrder_LAST_ADDED,
					Count:  1,
				})
			},
		},
	})

	//A4 outgoing heal buff
	modifier.Register(A4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})

}

func (c *char) initTraces() {

	//A2
	if c.info.Traces["1101101"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
		})
	}

	//A4
	if c.info.Traces["1101102"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A4,
			Source: c.id,
			State:  info.PropMap{prop.HealBoost: 0.1},
		})
	}

	//A6
	if c.info.Traces["1101103"] {
		c.engine.ExtendModifierDuration(c.id, Skill, 1)
	}
}
