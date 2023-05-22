package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) ExtendDuration(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == modifier {
			old := mod.Duration
			mod.Duration += amt

			f := modifierCatalog[modifier].Listeners.OnExtendDuration
			if f != nil {
				f(mgr.engine, mod)
			}
			mgr.engine.Events().ModifierExtended.Emit(event.ModifierExtendedEvent{
				Target:    target,
				Modifier:  modifier,
				Operation: "ExtendDuration",
				OldValue:  old,
				NewValue:  mod.Duration,
			})
		}
	}
}

func (mgr *Manager) ExtendCount(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == modifier {
			old := mod.Count
			mod.Count += amt
			if mod.Count > mod.MaxCount {
				mod.Count = mod.MaxCount
			}

			f := modifierCatalog[modifier].Listeners.OnExtendCount
			if f != nil {
				f(mgr.engine, mod)
			}
			mgr.engine.Events().ModifierExtended.Emit(event.ModifierExtendedEvent{
				Target:    target,
				Modifier:  modifier,
				Operation: "ExtendCount",
				OldValue:  old,
				NewValue:  mod.Count,
			})
		}
	}
}
