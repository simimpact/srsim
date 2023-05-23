package modifier

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type Listeners struct {
	OnAdd            func(engine engine.Engine, modifier *info.ModifierInstance)
	OnRemove         func(engine engine.Engine, modifier *info.ModifierInstance)
	OnExtendDuration func(engine engine.Engine, modifier *info.ModifierInstance)
	OnExtendCount    func(engine engine.Engine, modifier *info.ModifierInstance)
}

func (mgr *Manager) subscribe() {

}

func (mgr *Manager) emitAdd(target key.TargetID, mod *info.ModifierInstance, chance float64) {
	f := modifierCatalog[mod.Name].Listeners.OnAdd
	if f != nil {
		f(mgr.engine, mod)
	}
	mgr.engine.Events().ModifierAdded.Emit(event.ModifierAddedEvent{
		Target:   target,
		Modifier: mod,
		Chance:   chance,
	})
}

func (mgr *Manager) emitRemove(target key.TargetID, mod *info.ModifierInstance) {
	f := modifierCatalog[mod.Name].Listeners.OnRemove
	if f != nil {
		f(mgr.engine, mod)
	}
	mgr.engine.Events().ModifierRemoved.Emit(event.ModifierRemovedEvent{
		Target:   target,
		Modifier: mod,
	})
}

func (mgr *Manager) emitExtendDuration(target key.TargetID, mod *info.ModifierInstance, old int) {
	f := modifierCatalog[mod.Name].Listeners.OnExtendDuration
	if f != nil {
		f(mgr.engine, mod)
	}
	mgr.engine.Events().ModifierExtended.Emit(event.ModifierExtendedEvent{
		Target:    target,
		Modifier:  mod,
		Operation: "ExtendDuration",
		OldValue:  old,
		NewValue:  mod.Duration,
	})
}

func (mgr *Manager) emitExtendCount(target key.TargetID, mod *info.ModifierInstance, old int) {
	f := modifierCatalog[mod.Name].Listeners.OnExtendCount
	if f != nil {
		f(mgr.engine, mod)
	}
	mgr.engine.Events().ModifierExtended.Emit(event.ModifierExtendedEvent{
		Target:    target,
		Modifier:  mod,
		Operation: "ExtendCount",
		OldValue:  old,
		NewValue:  mod.Count,
	})
}
