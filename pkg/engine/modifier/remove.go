package modifier

import (
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	i := 0
	var removedMods []*ModifierInstance
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier {
			removedMods = append(removedMods, mod)
		} else {
			mgr.targets[target][i] = mod
			i++
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
	mgr.emitRemove(target, removedMods)
}

func (mgr *Manager) RemoveModifierFromSource(target, source key.TargetID, modifier key.Modifier) {
	i := 0
	var removedMods []*ModifierInstance
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier && mod.source == source {
			removedMods = append(removedMods, mod)
		} else {
			mgr.targets[target][i] = mod
			i++
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
	mgr.emitRemove(target, removedMods)
}

func (mgr *Manager) RemoveSelf(target key.TargetID, instance *ModifierInstance) {
	for i, mod := range mgr.targets[target] {
		if mod == instance {
			last := len(mgr.targets[target]) - 1
			mgr.targets[target][i] = mgr.targets[target][last]
			mgr.targets[target] = mgr.targets[target][:last]
			mgr.emitRemove(target, []*ModifierInstance{instance})
			return
		}
	}
}
