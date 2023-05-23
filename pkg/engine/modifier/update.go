package modifier

import (
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) ExtendDuration(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == modifier {
			old := mod.Duration
			mod.Duration += amt
			mgr.emitExtendDuration(target, mod, old)
		}
	}
}

func (mgr *Manager) ExtendCount(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == modifier {
			old := mod.Count
			mod.Count += amt
			if mod.MaxCount > 0 && mod.Count > mod.MaxCount {
				mod.Count = mod.MaxCount
			}
			mgr.emitExtendCount(target, mod, old)
		}
	}
}
