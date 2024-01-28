package himeko

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
)

const (
	a2 = "himeko-a2"
)

func init() {
	modifier.Register(a2, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: a2Listener,
		},
	})
}

func (c *char) initTraces() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   a2,
		Source: c.id,
	})
}

func a2Listener(mod *modifier.Instance, e event.AttackEnd) {
	for _, target := range e.Targets {
		mod.Engine().AddModifier(target, info.Modifier{
			Name:   common.Burn,
			Source: e.Attacker,
			State: &common.BurnState{
				DamagePercentage:    0.3,
				DamageValue:         0,
				DEFDamagePercentage: 0,
			},
			Chance:   1,
			Duration: 2,
		})
	}

}
