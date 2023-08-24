package seele

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Resurgence key.Modifier = "seele-resurgence"
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
	// enter buffed state and add extra turn if not on resurgence
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {

	})
}

// enter "buffed state", insert extra turn right after if addExtraTurn is true.
func (c *char) enterResurgence(addExtraTurn bool) {
	// A4 : While Seele is in the buffed state, her Quantum RES PEN increases by 20%.
	penAmt := 0.0
	if c.info.Traces["102"] {
		penAmt = 0.2
	}
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Resurgence,
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
