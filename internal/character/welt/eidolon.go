package welt

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1            = "welt-e1"
	E2 key.Reason = "welt-e2"
	E6 key.Attack = "welt-e6"
)

type E1State struct {
	basicAtkMult  float64
	skillMult     float64
	enhancedCount int
}

// E1 : After Welt uses his Ultimate, his abilities are enhanced.
// The next 2 time(s) he uses his Basic ATK or Skill,
// deals Additional DMG to the target equal to 50% of his Basic ATK's
// DMG multiplier or 80% of his Skill's DMG multiplier respectively.
// E2 : When his Talent is triggered, Welt regenerates 3 Energy.
// E4 : Base chance for Skill to inflict SPD Reduction increases by 35%.
// E6 : When using Skill, deals DMG for 1 extra time to a random enemy.

func init() {
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeAttack: pursuedDmgOnEnhanced,
		},
		Stacking: modifier.ReplaceBySource,
	})
}

func (c *char) initEidolons() {
	// E1 : add enhanced mod
	modState := E1State{
		basicAtkMult:  0.5,
		skillMult:     0.8,
		enhancedCount: 0,
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:     E1,
		Source:   c.id,
		Duration: 2,
		State:    &modState,
	})
	// refresh counter on each ult cast
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		if e.Owner == c.id && e.AttackType == model.AttackType_ULT {
			modState.enhancedCount = 2
		}
	})
}

// E1 : add pursued attack on basic/skill if has enhanced mod.
func pursuedDmgOnEnhanced(mod *modifier.Instance, e event.AttackStart) {
	state := mod.State().(*E1State)
	if state.enhancedCount <= 0 {
		return
	}
	// setup variable for basic/skill
	atkMult := 0.0
	if e.AttackType == model.AttackType_NORMAL {
		atkMult = 0.5 * state.basicAtkMult
	} else if e.AttackType == model.AttackType_SKILL {
		atkMult = 0.8 * state.skillMult
	}

	// attack
	mod.Engine().Attack(info.Attack{
		Key:        E1,
		Targets:    e.Targets,
		Source:     mod.Owner(),
		AttackType: model.AttackType_PURSUED,
		DamageType: model.DamageType_IMAGINARY,
		BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: atkMult},
	})
	state.enhancedCount--
}
