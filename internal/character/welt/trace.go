package welt

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A4 key.Reason = "welt-a4"
	A6 key.Reason = "welt-a6"
)

func init() {

}

func (c *char) initTraces() {
	// A4 : Using Ultimate additionally regenerates 10 Energy.
	c.engine.Events().ActionEnd.Subscribe(func(e event.ActionEnd) {
		if e.Owner != c.id ||
			!c.info.Traces["102"] ||
			e.AttackType != model.AttackType_ULT {
			return
		}
		// add flat energy
		c.engine.ModifyEnergy(info.ModifyAttribute{
			Key:    A4,
			Target: c.id,
			Source: c.id,
			Amount: 10,
		})
	})

	// A6 : Deals 20% more DMG to enemies inflicted with Weakness Break.
	c.engine.Events().HitStart.Subscribe(func(e event.HitStart) {
		// TODO : DM uses modifier check for StanceBreakState
		if e.Attacker != c.id ||
			!c.info.Traces["103"] ||
			c.engine.Stance(e.Defender) != 0 {
			return
		}
		e.Hit.Attacker.AddProperty(A6, prop.AllDamagePercent, 0.2)
	})
}
