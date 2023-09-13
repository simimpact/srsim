package march7th

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 = "march7th-e2"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.Engine().AddShield(E2, info.Shield{
					Source: mod.Source(),
					Target: mod.Owner(),
					BaseShield: info.ShieldMap{
						model.ShieldFormula_SHIELD_BY_SHIELDER_DEF: 0.24,
					},
					ShieldValue: 320,
				})
			},
			OnRemove: func(mod *modifier.Instance) {
				mod.Engine().RemoveShield(E2, mod.Owner())
			},
		},
	})
}

func (c *char) addE2Shield(e event.BattleStart) {
	lowestHpRatio := 1.0
	var lowestHpTarget key.TargetID
	for _, Target := range c.engine.Characters() {
		canidateHp := c.engine.HPRatio(Target)
		if canidateHp <= lowestHpRatio && canidateHp > 0.0 {
			lowestHpRatio = canidateHp
			lowestHpTarget = Target
		}
	}

	c.engine.AddModifier(lowestHpTarget, info.Modifier{
		Name:     E2,
		Source:   c.id,
		Duration: 3,
	})
}
