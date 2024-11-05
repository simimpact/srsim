package gallagher

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	A2 = "gallagher-a2"
	A4 = "gallagher-a4"
	A6 = "gallagher-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Listeners: modifier.Listeners{
			OnPropertyChange: AdjustA2Buff,
		},
	})
}

func (c *char) initTraces() {
	if c.info.Traces["101"] {
		c.engine.Events().BattleStart.Subscribe(c.A2Init)
	}
}

func AdjustA2Buff(mod *modifier.Instance) {
	gallagher := mod.Engine().Stats(mod.Source())
	newHealBoost := 0.5 * gallagher.BreakEffect()
	if newHealBoost > 0.75 {
		newHealBoost = 0.75
	}
	mod.SetProperty(prop.HealBoost, newHealBoost)
}

func (c *char) A2Init(e event.BattleStart) {
	newHealBoost := 0.5 * c.engine.Stats(c.id).BreakEffect()
	if newHealBoost > 0.75 {
		newHealBoost = 0.75
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   A2,
		Source: c.id,
		Stats: info.PropMap{
			prop.HealBoost: newHealBoost,
		},
	})
}
