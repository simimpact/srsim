package danheng

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	E1 key.Modifier = "dan-heng-e1"
	E4 key.Modifier = "dan-heng-e4"
)

func init() {
	// When the target enemy's current HP percentage is greater than or equal to 50%, CRIT Rate increases by 12%.
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.Defender.CurrentHPRatio() >= 0.5 {
					e.Hit.Attacker.AddProperty(prop.CritChance, 0.12)
				}
			},
		},
	})

	// When Dan Heng uses his Ultimate to defeat an enemy, he will immediately take action again.
	// Note: this modifier is only active during ult
	modifier.Register(E4, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: func(mod *modifier.Instance, target key.TargetID) {
				mod.Engine().SetGauge(mod.Owner(), 0)
			},
		},
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E1,
			Source: c.id,
		})
	}
}
