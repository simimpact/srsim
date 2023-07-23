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
	isResurgence bool
	dmgAmt       float64
}

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
		isResurgence: false,
		dmgAmt:       talent[c.info.TalentLevelIndex()],
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Resurgence,
		Source: c.id,
		State:  &modState,
	})
}

// add buffedState mod, add extra turn if turn not on resurgence
func applyResurgence(mod *modifier.Instance, target key.TargetID) {
	state := mod.State().(*state)
	// add dmg% buff
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   BuffedState,
		Source: mod.Owner(),
		Stats:  info.PropMap{prop.AllDamagePercent: state.dmgAmt},
	})
	if !state.isResurgence {
		// TODO : implement extra turn mechanic here
	}
}
