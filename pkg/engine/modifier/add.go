package modifier

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddModifier(target key.TargetID, instance info.ModifierInstance) error {
	config := modifierCatalog[instance.Name]

	// TODO: not sure if should have this check?
	// if !ok {
	// 	return fmt.Errorf("no modifier registered for key: %v", instance.Name)
	// }
	if !mgr.engine.IsValid(target) {
		return fmt.Errorf("invalid target id: %v", target)
	}
	if !mgr.engine.IsValid(instance.Source) {
		return fmt.Errorf("invalid source id: %v", instance.Source)
	}

	chance, resisted, err := mgr.attemptResist(target, instance, config.BehaviorFlags)
	if err != nil || resisted {
		return err
	}

	// prepare the instance to be added to the target
	setDefaults(&instance)

	// cases that can stack. Need to ensure count is stackable pending what was specified on instance
	if config.Stacking == Replace || config.Stacking == ReplaceBySource || config.Stacking == Merge {
		if instance.Count <= 0 && instance.CountAddWhenStack > 0 {
			instance.Count = instance.CountAddWhenStack
		}
	}

	var result *info.ModifierInstance
	var newInstance bool
	switch config.Stacking {
	case Unique:
		result, newInstance = mgr.unique(target, &instance)
	case ReplaceBySource:
		result, newInstance = mgr.replaceBySource(target, &instance)
	case Replace:
		result, newInstance = mgr.replace(target, &instance)
	case Multiple:
		result, newInstance = mgr.multiple(target, &instance)
	case Refresh:
		result, newInstance = mgr.refresh(target, &instance)
	case Prolong:
		result, newInstance = mgr.prolong(target, &instance)
	case Merge:
		result, newInstance = mgr.merge(target, &instance)
	default:
		return fmt.Errorf("unsupported stacking method: %v", config.Stacking)
	}

	// When a matching instance exists, only the following will create a new instance:
	//	- Replace
	//	- ReplaceBySource
	//	- Multiple
	if newInstance {
		result.Reset(mgr.turnCount)
		mgr.emitAdd(target, result, chance)
	}
	return nil
}

// returns the 1) chance to apply, 2) true if resisted, 3) error
func (mgr *Manager) attemptResist(
	target key.TargetID, instance info.ModifierInstance, flags []model.BehaviorFlag) (float64, bool, error) {
	// if unspecified, do not resist
	if instance.Chance <= 0 {
		return -1, false, nil
	}

	srcStats := mgr.engine.Stats(instance.Source)
	trgtStats := mgr.engine.Stats(target)

	effectHitRate := srcStats.EffectHitRate()
	effectRES := trgtStats.EffectRES()
	debuffRES := trgtStats.GetDebuffRES(flags...)

	chance := instance.Chance * (1 + effectHitRate) * (1 - effectRES) * (1 - debuffRES)
	if mgr.engine.Rand().Float64() < chance {
		return chance, false, nil
	}

	// resisted, emit event
	mgr.engine.Events().ModifierResisted.Emit(event.ModifierResistedEvent{
		Target:     target,
		Source:     instance.Source,
		Modifier:   instance.Name,
		Chance:     chance,
		BaseChance: instance.Chance,
		EHR:        effectHitRate,
		EffectRES:  effectRES,
		DebuffRES:  debuffRES,
	})
	return chance, true, nil
}

func (mgr *Manager) unique(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			return mod, false
		}
	}
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) replaceBySource(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	for i, mod := range mgr.targets[target] {
		if mod.Name == instance.Name && mod.Source == instance.Source {
			// replace means this added instance is the new instance (can have new param values)
			instance.Count = replaceCount(instance, mod.Count)
			mgr.targets[target][i] = instance
			return instance, true
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) replace(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	for i, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// replace means this added instance is the new instance (can have new param values)
			instance.Count = replaceCount(instance, mod.Count)
			mgr.targets[target][i] = instance
			return instance, true
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) multiple(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) refresh(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, reset this modifier's duration
			old := mod.Duration
			mod.Duration = instance.Duration
			mgr.emitExtendDuration(target, mod, old)
			return mod, false
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) prolong(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, prolong the duration
			old := mod.Duration
			mod.Duration += instance.Duration
			mgr.emitExtendDuration(target, mod, old)
			return mod, false
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func (mgr *Manager) merge(target key.TargetID, instance *info.ModifierInstance) (*info.ModifierInstance, bool) {
	// set count to be CountAddWhenStack for replace case if Count is unspecified
	if instance.Count <= 0 && instance.CountAddWhenStack > 0 {
		instance.Count = instance.CountAddWhenStack
	}

	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, merge
			mod.Count = replaceCount(instance, mod.Count)
			if instance.Duration > mod.Duration {
				mod.Duration = instance.Duration
			}
			return mod, true
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance, true
}

func replaceCount(mod *info.ModifierInstance, prevCount int) int {
	if prevCount < 0 {
		return mod.Count
	}

	count := prevCount + mod.Count
	if mod.MaxCount > 0 && count > mod.MaxCount {
		count = mod.MaxCount
	}
	return count
}
