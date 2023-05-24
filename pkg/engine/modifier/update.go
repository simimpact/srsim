package modifier

import (
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) ExtendDuration(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier {
			old := mod.duration
			mod.duration += amt
			mgr.emitExtendDuration(target, mod, old)
		}
	}
}

func (mgr *Manager) ExtendCount(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier {
			old := mod.count
			mod.count += amt
			if mod.maxCount > 0 && mod.count > mod.maxCount {
				mod.count = mod.maxCount
			}
			mgr.emitExtendCount(target, mod, old)
		}
	}
}
