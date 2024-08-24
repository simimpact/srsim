package tingyun

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 = "tingyun-a2"
	A4 = "tingyun-a4"
	A6 = "tingyun-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		Duration:   1,
		TickMoment: modifier.ModifierPhase1End,
	})

	modifier.Register(A4, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: addA4,
		},
	})

	modifier.Register(A6, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: addA6,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["103"] {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A6,
			Source: c.id,
		})
	}
}

func addA4(mod *modifier.Instance, e event.HitStart) {
	ting, _ := mod.Engine().CharacterInfo(mod.Owner())
	if ting.Traces["102"] {
		e.Hit.Attacker.AddProperty(A4, prop.AllDamagePercent, 0.4)
	}
}

func addA6(mod *modifier.Instance) {
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    A6,
		Target: mod.Owner(),
		Source: mod.Owner(),
		Amount: 5.0,
	})
}
