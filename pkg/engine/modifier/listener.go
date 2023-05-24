package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/key"
)

type Listeners struct {
	OnAdd            func(mod *ModifierInstance)
	OnRemove         func(mod *ModifierInstance)
	OnExtendDuration func(mod *ModifierInstance)
	OnExtendCount    func(mod *ModifierInstance)
	OnPropertyChange func(mod *ModifierInstance)
}

func (mgr *Manager) subscribe() {

}

func (mgr *Manager) emitPropertyChange(target key.TargetID) {
	for _, mod := range mgr.targets[target] {
		f := mod.listeners.OnPropertyChange
		if f != nil {
			f(mod)
		}
	}
}

func (mgr *Manager) emitAdd(target key.TargetID, mod *ModifierInstance, chance float64) {
	f := mod.listeners.OnAdd
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierAdded.Emit(event.ModifierAddedEvent{
		Target:   target,
		Modifier: mod.ToModel(),
		Chance:   chance,
	})
}

func (mgr *Manager) emitRemove(target key.TargetID, mods []*ModifierInstance) {
	for _, mod := range mods {
		if len(mod.stats) > 0 {
			mgr.emitPropertyChange(target)
		}

		f := mod.listeners.OnRemove
		if f != nil {
			f(mod)
		}
		mgr.engine.Events().ModifierRemoved.Emit(event.ModifierRemovedEvent{
			Target:   target,
			Modifier: mod.ToModel(),
		})
	}
}

func (mgr *Manager) emitExtendDuration(target key.TargetID, mod *ModifierInstance, old int) {
	f := mod.listeners.OnExtendDuration
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtended.Emit(event.ModifierExtendedEvent{
		Target:    target,
		Modifier:  mod.ToModel(),
		Operation: "ExtendDuration",
		OldValue:  old,
		NewValue:  mod.duration,
	})
}

func (mgr *Manager) emitExtendCount(target key.TargetID, mod *ModifierInstance, old int) {
	f := mod.listeners.OnExtendCount
	if f != nil {
		f(mod)
	}
	mgr.engine.Events().ModifierExtended.Emit(event.ModifierExtendedEvent{
		Target:    target,
		Modifier:  mod.ToModel(),
		Operation: "ExtendCount",
		OldValue:  old,
		NewValue:  mod.count,
	})
}
