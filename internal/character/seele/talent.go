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

// Enters the buffed state upon defeating an enemy with Basic ATK, Skill, or Ultimate,
// and receives an extra turn. While in the buffed state,
// the DMG of Seele's attacks increases by 80% for 1 turn(s).
// Enemies defeated in the extra turn provided by "Resurgence"
// will not trigger another "Resurgence."

func init() {
	modifier.Register(Resurgence, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) initTalent() {
	// listen to death events. If killer is seele,
	// enter buffed state and set trigger for extra/resurgence turn.
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Killer == c.id {
			c.hasKilled = true
		}
	})

	// only runs on seele basic, skill, or ult. effectively excluding E6 pursued dmg etc.
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		// if action is seele's and hasKilled is triggered, reset and enter buffed state.
		// then if not on insert turn, insertAction. reset .hasKilled on all actionEnd.
		if e.Owner == c.id && c.hasKilled {
			c.hasKilled = false
			c.enterBuffedState()
			if !e.IsInsert {
				c.engine.InsertAction(c.id)
			}
		}
		c.hasKilled = false
	})
}

func (c *char) enterBuffedState() {
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
		Duration: 1,
	})
}
