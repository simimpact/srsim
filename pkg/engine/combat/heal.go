package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
)

func (mgr *Manager) Heal(heal info.Heal) {
	for _, t := range heal.Targets {
		baseHeal := make(info.HealMap, len(heal.BaseHeal))
		for k, v := range heal.BaseHeal {
			baseHeal[k] = v
		}

		e := &event.HealStart{
			Target:      mgr.attr.Stats(t),
			Healer:      mgr.attr.Stats(heal.Source),
			BaseHeal:    baseHeal,
			HealValue:   heal.HealValue,
			UseSnapshot: heal.UseSnapshot,
		}
		mgr.event.HealStart.Emit(e)

		// TODO: Perform Heal. Use the data in the event to perform the heal
		// TOOD: call ModifyHP to add the new HP to the healed target

		mgr.event.HealEnd.Emit(event.HealEnd{
			Target:             t,
			Healer:             heal.Source,
			HealAmount:         0, // TODO:
			OverflowHealAmount: 0, // TODO:
			UseSnapshot:        heal.UseSnapshot,
		})
	}
}
