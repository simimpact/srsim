package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	i := 0
	for _, mod := range mgr.targets[target] {
		if mod.Name != modifier {
			// keep
			mgr.targets[target][i] = mod
			i++
		} else {
			mgr.remove(target, mod)
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
}

func (mgr *Manager) RemoveModifierFromSource(target, source key.TargetID, modifier key.Modifier) {
	i := 0
	for _, mod := range mgr.targets[target] {
		if mod.Name != modifier && mod.Source != source {
			// keep
			mgr.targets[target][i] = mod
			i++
		} else {
			mgr.remove(target, mod)
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
}

func (mgr *Manager) remove(target key.TargetID, mod *info.ModifierInstance) {
	f := modifierCatalog[mod.Name].Listeners.OnRemove
	if f != nil {
		f(mgr.engine, mod)
	}
	mgr.engine.Events().ModifierRemoved.Emit(event.ModifierRemovedEvent{
		Target:   target,
		Modifier: mod,
	})
}
