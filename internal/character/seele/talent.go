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
	// listen to death events. If killer is seele, enter resurgence.
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Killer == c.id {
			c.enterBuffedState()
		}
	})

	// TODO : add turnEnd sub/afterAction listener to check for resurgence and insert turn
	// if resurgence == true.
	// TODO : logic need to ask if seele killed from her action(not from pursued atk.)
	// TODO : add mechanism to check for isInsert and block inserts until after action ends.

	// wait and listen for actionEnd to insert a resurgence turn.
	// further check if seele killed an enemy on this turn
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		// TODO : figure out how to check for enemy death too. or do i make proc-trigger listener?
		if c.resurgence {
			c.resurgence = false
		} else {
			c.resurgence = true
			c.engine.InsertAction(c.id)
		}
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
