package modifier

import (
	"github.com/simimpact/srsim/pkg/key"
)

func (mgr *Manager) ExtendDuration(target key.TargetID, modifier key.Modifier, amt int) {
	for _, mod := range mgr.itr(target) {
		if mod.name == modifier {
			old := mod.duration
			mod.duration += amt
			mgr.emitExtendDuration(target, mod, old)
		}
	}
}

func (mgr *Manager) ExtendCount(target key.TargetID, modifier key.Modifier, amt float64) {
	// update counts
	for _, mod := range mgr.targets[target] {
		if mod.name != modifier {
			continue
		}

		old := mod.count
		mod.count += amt
		if mod.maxCount > 0 && mod.count > mod.maxCount {
			mod.count = mod.maxCount
		}
		mgr.emitExtendCount(target, mod, old)
	}

	// remove any updated that have 0 counts
	i := 0
	var removedMods []*Instance
	for _, mod := range mgr.targets[target] {
		if mod.name == modifier && mod.count <= 0 {
			removedMods = append(removedMods, mod)
		} else {
			mgr.targets[target][i] = mod
			i++
		}
	}
	mgr.targets[target] = mgr.targets[target][:i]
	mgr.emitRemove(target, removedMods)
}
