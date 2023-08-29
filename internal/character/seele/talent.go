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
	// listen to death events. If killer is seele, enter buffed state.
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Killer == c.id {
			c.enterBuffedState()
		}
		// set hasKilled trigger to true
		c.hasKilled = true
	})

	// only add extra turn after an action ends.
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		// insert resurgence turn only if action is seele's, hasKilled is triggered,
		// and not already in an insert/resurgence turn.
		// TODO : fix hasKilled trigger reset.
		// SCENARIO : seele skill killed an enemy -> trigger targetDeath sub -> set .hasKilled T
		// -> enter actionEnd, insert condition met -> add extra turn. (IDEAL, works now.)
		// INSIGHT :
		// make sure before she killed an enemy .hasKilled always = FALSE
		// .hasKilled currently only set to FALSE ONLY IF last resurgence is triggered.
		// -> only good for battleStart -> action -> kill -> action end -> extra turn -> action
		// -> kill -> action end, IF not met because .IsInsert
		// CONCLUSION :
		// need to guard against :
		// extra turn -> action -> kill, .hasKilled = T -> IF not met -> action
		// -> IF met, insertAction. even though has not killed this action.
		// figure out WHEN to reset .hasKilled
		// if e.Owner == c.id && c.hasKilled && !e.IsInsert {
		// 	c.hasKilled = false
		// 	c.engine.InsertAction(c.id)
		// }
		if e.Owner == c.id { // only run func on seele's actions
			if c.hasKilled { // only allows insert logic to run ONLY IF seele has killed an enemy
				c.hasKilled = false
				if !e.IsInsert { // allow insert only if not already on insert turn.
					c.engine.InsertAction(c.id)
				}
			}
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
