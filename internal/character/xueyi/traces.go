package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	A2 = "xueyi-a2"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnPropertyChange: checkBreakEffect,
		},
	})
}

func (c *char) initTraces() {
	initialBuff := c.engine.Stats(c.id).BreakEffect()
	if initialBuff > 2.4 {
		initialBuff = 2.4
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   A2,
		Source: c.id,
		Stats: info.PropMap{
			prop.BreakEffect: initialBuff,
		},
	})
}

func checkBreakEffect(mod *modifier.Instance) {
	stats := mod.OwnerStats()
	dmgbuff := stats.BreakEffect()
	if dmgbuff > 2.4 {
		dmgbuff = 2.4
	}
	mod.SetProperty(prop.AllDamagePercent, dmgbuff)
}
