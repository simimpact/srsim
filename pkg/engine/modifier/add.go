package modifier

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddModifier(target key.TargetID, modifier info.Modifier) error {
	config := modifierCatalog[modifier.Name]

	if !mgr.engine.IsValid(target) {
		return fmt.Errorf("invalid target id: %v", target)
	}
	if !mgr.engine.IsValid(modifier.Source) {
		return fmt.Errorf("invalid source id: %v", modifier.Source)
	}

	chance, resisted, err := mgr.attemptResist(target, modifier, config.BehaviorFlags)
	if err != nil || resisted {
		return err
	}

	// prepare the instance to be added to the target
	instance := mgr.newInstance(target, modifier)

	var result *ModifierInstance
	var newInstance bool
	switch config.Stacking {
	case Unique:
		result, newInstance = mgr.unique(target, instance)
	case ReplaceBySource:
		result, newInstance = mgr.replaceBySource(target, instance)
	case Replace:
		result, newInstance = mgr.replace(target, instance)
	case Multiple:
		result, newInstance = mgr.multiple(target, instance)
	case Refresh:
		result, newInstance = mgr.refresh(target, instance)
	case Prolong:
		result, newInstance = mgr.prolong(target, instance)
	case Merge:
		result, newInstance = mgr.merge(target, instance)
	default:
		return fmt.Errorf("unsupported stacking method: %v", config.Stacking)
	}

	// When a matching instance exists, only the following will create a new instance:
	//	- Replace
	//	- ReplaceBySource
	//	- Multiple
	//	- Merge (special case, keeps old but triggers reset + add as if new)
	if newInstance {
		result.renewTurn = mgr.turnCount
		if len(result.stats) > 0 {
			// special case if new instance and Stats were predefined
			mgr.emitPropertyChange(target)
		}
		mgr.emitAdd(target, result, chance)
	}
	return nil
}

// returns the 1) chance to apply, 2) true if resisted, 3) error
func (mgr *Manager) attemptResist(
	target key.TargetID, mod info.Modifier, flags []model.BehaviorFlag) (float64, bool, error) {
	// if unspecified, do not resist
	if mod.Chance <= 0 {
		return -1, false, nil
	}

	srcStats := mgr.engine.Stats(mod.Source)
	trgtStats := mgr.engine.Stats(target)

	effectHitRate := srcStats.EffectHitRate()
	effectRES := trgtStats.EffectRES()
	debuffRES := trgtStats.GetDebuffRES(flags...)

	chance := mod.Chance * (1 + effectHitRate) * (1 - effectRES) * (1 - debuffRES)
	if mgr.engine.Rand().Float64() < chance {
		return chance, false, nil
	}

	// resisted, emit event
	mgr.engine.Events().ModifierResisted.Emit(event.ModifierResistedEvent{
		Target:     target,
		Source:     mod.Source,
		Modifier:   mod.Name,
		Chance:     chance,
		BaseChance: mod.Chance,
		EHR:        effectHitRate,
		EffectRES:  effectRES,
		DebuffRES:  debuffRES,
	})
	return chance, true, nil
}

func (mgr *Manager) unique(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			return mod, false
		}
	}
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) replaceBySource(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
	for i, mod := range mgr.targets[target] {
		if mod.name == instance.name && mod.source == instance.source {
			// replace means this added instance is the new instance (can have new param values)
			instance.count = stackCount(instance, mod.count)
			mgr.targets[target][i] = instance
			return instance, true
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) replace(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
	for i, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			// replace means this added instance is the new instance (can have new param values)
			instance.count = stackCount(instance, mod.count)
			mgr.targets[target][i] = instance
			return instance, true
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) multiple(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) refresh(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
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

func (mgr *Manager) prolong(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
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

func (mgr *Manager) merge(target key.TargetID, instance *ModifierInstance) (*ModifierInstance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.name == instance.name {
			// found a matching modifier, merge
			mod.count = stackCount(instance, mod.count)
			if instance.duration > mod.duration {
				mod.duration = instance.duration
			}
			return mod, true
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func stackCount(mod *ModifierInstance, prevCount float64) float64 {
	if prevCount < 0 || mod.count < 0 {
		return mod.count
	}

	count := prevCount + mod.count
	if mod.maxCount > 0 && count > mod.maxCount {
		count = mod.maxCount
	}
	return count
}
