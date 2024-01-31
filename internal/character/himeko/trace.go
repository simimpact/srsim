package himeko

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	a2 = "himeko-a2"
	a6 = "himeko-a6"
)

func init() {
	modifier.Register(a2, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: a2Listener,
		},
	})

	modifier.Register(a6, modifier.Config{
		Listeners: modifier.Listeners{
			OnHPChange: a6Listener,
		},
	})
}

func (c *char) initTraces() {
	critRateAmt := 0.15
	if c.engine.HPRatio(c.id) <= 0.8 {
		critRateAmt = 0
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   a2,
		Source: c.id,
		Stats: info.PropMap{
			prop.CritChance: critRateAmt,
		},
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

func a6Listener(mod *modifier.Instance, e event.HPChange) {
	if e.NewHPRatio >= 0.8 {
		mod.SetProperty(prop.CritChance, 0.15)
	} else {
		mod.SetProperty(prop.CritChance, 0)
	}
}
