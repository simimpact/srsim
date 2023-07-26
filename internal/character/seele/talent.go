package seele

import (
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

// add mod to check for resurgence turns and add extra action
func (c *char) initTalent() {
	modState := state{
		c:      c,
		dmgAmt: talent[c.info.TalentLevelIndex()],
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Resurgence,
		Source: c.id,
		State:  &modState,
	})
}

// enter buffed state and add extra turn if not on resurgence
func applyResurgence(mod *modifier.Instance, target key.TargetID) {
	state := mod.State().(*state)
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	penAmt := 0.0
	if state.c.info.Traces["102"] {
		penAmt = 0.2
	}
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   BuffedState,
		Source: mod.Owner(),
		Stats: info.PropMap{
			prop.AllDamagePercent: state.dmgAmt,
			prop.QuantumPEN:       penAmt,
		},
	})
	// if current turn is resurgence, don't add extra turn. otherwise add extra turn.
	if state.c.resurgence {
		state.c.resurgence = false
	} else {
		state.c.resurgence = true
		state.c.engine.InsertAction(state.c.id)
	}
}
