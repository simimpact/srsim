package modifier

import (
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	i := 0
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier {
			mgr.emitRemove(target, mod)
		} else {
			mgr.targets[target][i] = mod
			i++
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
}

func (mgr *Manager) RemoveModifierFromSource(target, source key.TargetID, modifier key.Modifier) {
	i := 0
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier && mod.source == source {
			mgr.emitRemove(target, mod)
		} else {
			mgr.targets[target][i] = mod
			i++
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
}
