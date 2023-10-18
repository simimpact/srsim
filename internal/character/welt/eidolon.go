package welt

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 key.Attack = "welt-e1"
	E6 key.Attack = "welt-e6"
)

// E1 : After Welt uses his Ultimate, his abilities are enhanced.
// The next 2 time(s) he uses his Basic ATK or Skill,
// deals Additional DMG to the target equal to 50% of his Basic ATK's
// DMG multiplier or 80% of his Skill's DMG multiplier respectively.
// E6 : When using Skill, deals DMG for 1 extra time to a random enemy.

func (c *char) initEidolons() {
	// E1 : refresh counter on each ult cast
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		if e.Owner == c.id && e.AttackType == model.AttackType_ULT {
			c.enhancedCount = 2
		}
	})
}

// E1 : add pursued attack on basic/skill if has enhanced mod.
func (c *char) applyE1Pursued(target key.TargetID, multiplier float64) {
	if c.enhancedCount <= 0 {
		return
	}
	c.engine.Attack(info.Attack{
		Key:        E1,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_IMAGINARY,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: multiplier,
		},
	})
	c.enhancedCount--
}
