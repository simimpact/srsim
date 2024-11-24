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
		CanDispel:  true,
	})
}

func (c *char) initTalent() {
	// if seele kills in her turn, set extra turn trigger and add buffedState mod.
	c.engine.Events().TargetDeath.Subscribe(func(e event.TargetDeath) {
		if e.Killer == c.id && c.id == c.actionOwner && !c.isInsertTurn {
			c.shouldInsert = true
			c.enterBuffedState()
		}
	})

	// actionStart subs to enable adding buffedState mod effectively after-kill
	c.engine.Events().ActionStart.Subscribe(func(e event.ActionStart) {
		if c.id == e.Owner {
			if e.IsInsert {
				c.isInsertTurn = true
			} else {
				c.actionOwner = e.Owner
				c.isInsertTurn = false
			}
		}
	})

	// extra turn logic
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		if c.shouldInsert {
			c.engine.InsertAction(c.id)
		}
		c.shouldInsert = false
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
