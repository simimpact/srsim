package modifier

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddModifier(target key.TargetID, modifier info.Modifier) (bool, error) {
	config := modifierCatalog[modifier.Name]

	if !mgr.engine.IsValid(target) {
		return false, fmt.Errorf("invalid target id: %v", target)
	}
	if !mgr.engine.IsValid(modifier.Source) {
		return false, fmt.Errorf("invalid source id: %v", modifier.Source)
	}

	chance, resisted := mgr.attemptResist(target, modifier, config.BehaviorFlags)
	if resisted {
		return false, nil
	}

	// prepare the instance to be added to the target
	instance := mgr.newInstance(target, modifier, mgr.turnCount)

	var result *Instance
	var newInstance bool
	switch config.Stacking {
	case Unique:
		result, newInstance = mgr.unique(target, instance)
	case ReplaceBySource:
		result = mgr.replaceBySource(target, instance)
		newInstance = true
	case Replace:
		result = mgr.replace(target, instance)
		newInstance = true
	case Multiple:
		result, newInstance = mgr.multiple(target, instance)
	case Refresh:
		result, newInstance = mgr.refresh(target, instance)
	case Prolong:
		result, newInstance = mgr.prolong(target, instance)
	case Merge:
		result = mgr.merge(target, instance)
		newInstance = true
	default:
		return false, fmt.Errorf("unsupported stacking method: %v", config.Stacking)
	}

	// When a matching instance exists, only the following will create a new instance:
	//	- Replace
	//	- ReplaceBySource
	//	- Multiple
	//	- Merge (special case, keeps old but triggers reset + add as if new)
	if newInstance {
		if len(result.stats) > 0 {
			// special case if new instance and Stats were predefined
			mgr.emitPropertyChange(target)
		}
		mgr.emitAdd(target, result, chance)
	}
	return true, nil
}

// returns the 1) chance to apply, 2) true if resisted, 3) error
func (mgr *Manager) attemptResist(
	target key.TargetID, mod info.Modifier, flags []model.BehaviorFlag) (float64, bool) {
	// if unspecified, do not resist
	if mod.Chance <= 0 {
		return -1, false
	}

	srcStats := mgr.engine.Stats(mod.Source)
	trgtStats := mgr.engine.Stats(target)

	effectHitRate := srcStats.EffectHitRate()
	effectRES := trgtStats.EffectRES()
	debuffRES := trgtStats.GetDebuffRES(flags...)

	chance := mod.Chance * (1 + effectHitRate) * (1 - effectRES) * (1 - debuffRES)
	if mgr.engine.Rand().Float64() < chance {
		return chance, false
	}

	// resisted, emit event
	mgr.engine.Events().ModifierResisted.Emit(event.ModifierResisted{
		Target:        target,
		Source:        mod.Source,
		Modifier:      mod.Name,
		Chance:        chance,
		BaseChance:    mod.Chance,
		EffectHitRate: effectHitRate,
		EffectRES:     effectRES,
		DebuffRES:     debuffRES,
	})
	return chance, true
}

func (mgr *Manager) unique(target key.TargetID, instance *Instance) (*Instance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			return mod, false
		}
	}
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) replaceBySource(target key.TargetID, instance *Instance) *Instance {
	for i, mod := range mgr.targets[target] {
		if mod.name == instance.name && mod.source == instance.source {
			// replace means this added instance is the new instance (can have new param values)
			instance.count = stackCount(instance, mod.count)
			mgr.targets[target][i] = instance
			return instance
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

func (mgr *Manager) replace(target key.TargetID, instance *Instance) *Instance {
	for i, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			// replace means this added instance is the new instance (can have new param values)
			instance.count = stackCount(instance, mod.count)
			mgr.targets[target][i] = instance
			return instance
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

func (mgr *Manager) multiple(target key.TargetID, instance *Instance) (*Instance, bool) {
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) refresh(target key.TargetID, instance *Instance) (*Instance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			// found a matching modifier, reset this modifier's duration
			old := mod.duration
			mod.duration = instance.duration
			mgr.emitExtendDuration(target, mod, old)
			return mod, false
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) prolong(target key.TargetID, instance *Instance) (*Instance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			// found a matching modifier, prolong the duration
			old := mod.duration
			mod.duration += instance.duration
			mgr.emitExtendDuration(target, mod, old)
			return mod, false
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) merge(target key.TargetID, instance *Instance) *Instance {
	for _, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			// found a matching modifier, merge
			mod.count = stackCount(instance, mod.count)
			if instance.duration > mod.duration {
				mod.duration = instance.duration
			}
			return mod
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

func stackCount(mod *Instance, prevCount float64) float64 {
	if prevCount < 0 || mod.count < 0 {
		return mod.count
	}

	count := prevCount + mod.count
	if mod.maxCount > 0 && count > mod.maxCount {
		count = mod.maxCount
	}
	return count
}
