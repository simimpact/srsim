package seele

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Resurgence  key.Modifier = "seele-resurgence"
	BuffedState key.Modifier = "seele-buffed-state"
)

type state struct {
	c      *char
	dmgAmt float64
}

// Enters the buffed state upon defeating an enemy with Basic ATK, Skill, or Ultimate,
// and receives an extra turn. While in the buffed state,
// the DMG of Seele's attacks increases by 80% for 1 turn(s).
// Enemies defeated in the extra turn provided by "Resurgence"
// will not trigger another "Resurgence."

func init() {
	modifier.Register(Resurgence, modifier.Config{
		Listeners: modifier.Listeners{
			OnTriggerDeath: applyResurgence,
		},
	})
	modifier.Register(BuffedState, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

// add mod to add and check for resurgence turns
func (c *char) talentActionEndListener(e event.ActionEnd) {
	modState := state{
		c:      c,
		dmgAmt: talent[c.info.TalentLevelIndex()],
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Resurgence,
		Source: e.Owner,
		State:  &modState,
	})
}

// add buffedState mod, add extra turn if turn not on resurgence
func applyResurgence(mod *modifier.Instance, target key.TargetID) {
	state := mod.State().(*state)
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	penAmt := 0.0
	if state.c.info.Traces["102"] {
		penAmt = 0.2
	}
	// add dmg% buff
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   BuffedState,
		Source: mod.Owner(),
		Stats: info.PropMap{
			prop.AllDamagePercent: state.dmgAmt,
			prop.QuantumPEN:       penAmt,
		},
	})
	switch state.c.resurgence {
	case true:
		state.c.resurgence = false
	case false:
		// enter resurgence turn.
		state.c.resurgence = true

		// TODO : implement extra turn mechanic here
		state.c.engine.InsertAction(state.c.id)
	}
}
