package seele

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
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
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   1,
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
	// TODO : move onTriggerDeath logic to onDeath event subs to use the new resurgence func
}

// enter buffed state and add extra turn if not on resurgence
func applyResurgence(mod *modifier.Instance, target key.TargetID) {
	enterResurgence(true)
}

// TODO : refactor resurgence logic into a separate single func that all impl will use
// TODO : and merge buffedstate and resurgence appl.
func (c *char) enterResurgence(addExtraTurn bool) {
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	penAmt := 0.0
	if c.info.Traces["102"] {
		penAmt = 0.2
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   BuffedState,
		Source: c.id,
		Stats: info.PropMap{
			prop.AllDamagePercent: talent[c.info.TalentLevelIndex()],
			prop.QuantumPEN:       penAmt,
		},
	})
	// if addExtraTurn true, run extra turn logic. if not(just enter buffedState), then bypass.
	if addExtraTurn {
		if c.resurgence {
			c.resurgence = false
		} else {
			c.resurgence = true
			c.engine.InsertAction(c.id)
		}
	}
}
