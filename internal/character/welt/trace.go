package welt

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 key.Modifier = "welt-a2"
	A4 key.Reason   = "welt-a4"
	A6 key.Reason   = "welt-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
	})
}

func (c *char) initTraces() {
	// A2 : When using Ultimate, there is a 100% base chance to
	// 			increase the DMG received by the targets by 12% for 2 turn(s).
	// TODO : consider moving this logic to Ult.go
	c.engine.Events().AttackStart.Subscribe(func(e event.AttackStart) {
		// negative condition for early return
		if e.Attacker != c.id ||
			!c.info.Traces["101"] ||
			e.AttackType != model.AttackType_ULT {
			return
		}
		// inflict allDmgTaken debuff to all targets
		for _, target := range e.Targets {
			c.engine.AddModifier(target, info.Modifier{
				Name:     A2,
				Source:   c.id,
				Chance:   1,
				Duration: 2,
				Stats:    info.PropMap{prop.AllDamageTaken: 0.12},
			})
		}
	})

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
