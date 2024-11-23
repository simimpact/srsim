package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) RemoveModifier(target key.TargetID, modifier key.Modifier) {
	i := 0
	var removedMods []*Instance
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
	var removedMods []*Instance
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

func (mgr *Manager) RemoveSelf(target key.TargetID, instance *Instance) {
	for i, mod := range mgr.targets[target] {
		if mod != instance {
			continue
		}

		last := len(mgr.targets[target]) - 1
		mgr.targets[target][i] = mgr.targets[target][last]
		mgr.targets[target] = mgr.targets[target][:last]
		mgr.emitRemove(target, []*Instance{instance})
		return
	}
}

func (mgr *Manager) DispelStatus(target key.TargetID, dispel info.Dispel) {
	idx := 0
	idsToRemove := mgr.dispelIDs(target, dispel)
	removedMods := make([]*Instance, 0, len(idsToRemove))

	for i, mod := range mgr.targets[target] {
		if _, ok := idsToRemove[i]; ok {
			removedMods = append(removedMods, mod)
		} else {
			mgr.targets[target][idx] = mod
			idx++
		}
	}

	mgr.targets[target] = mgr.targets[target][:idx]
	mgr.emitDispel(target, removedMods)
}

func (mgr *Manager) dispelIDs(target key.TargetID, dispel info.Dispel) map[int]struct{} {
	if dispel.Count <= 0 {
		dispel.Count = len(mgr.targets[target])
	}

	l := len(mgr.targets[target])
	out := make(map[int]struct{})

	switch dispel.Order {
	case model.DispelOrder_FIRST_ADDED:
		for i := 0; i < l && len(out) < dispel.Count; i++ {
			if mgr.targets[target][i].statusType == dispel.Status && mgr.targets[target][i].canDispel {
				out[i] = struct{}{}
			}
		}
	case model.DispelOrder_LAST_ADDED:
		for i := len(mgr.targets[target]) - 1; i > 0 && len(out) < dispel.Count; i-- {
			if mgr.targets[target][i].statusType == dispel.Status && mgr.targets[target][i].canDispel {
				out[i] = struct{}{}
			}
		}
	case model.DispelOrder_RANDOM:
		var options []int
		for i, mod := range mgr.targets[target] {
			if mod.statusType == dispel.Status && mod.canDispel {
				options = append(options, i)
			}
		}

		mgr.engine.Rand().Shuffle(len(options), func(i, j int) {
			options[i], options[j] = options[j], options[i]
		})

		for i := 0; i < len(options) && i < dispel.Count; i++ {
			out[options[i]] = struct{}{}
		}
	}
	return out
}
