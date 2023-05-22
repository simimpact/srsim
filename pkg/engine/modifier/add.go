package modifier

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddModifier(target key.TargetID, instance info.ModifierInstance) error {
	config, ok := modifierCatalog[instance.Name]

	if !ok {
		return fmt.Errorf("no modifier registered for key: %v", instance.Name)
	}
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

	// do some initial configuration to this instance based on the config and how it was passed in
	if instance.Params == nil {
		instance.Params = make(map[string]float64)
	}
	// if any of these fields are 0, use the value in the config instead
	if instance.CountAddWhenStack == 0 {
		instance.CountAddWhenStack = config.CountAddWhenStack
	}
	if instance.MaxCount == 0 {
		instance.MaxCount = config.MaxCount
	}
	if instance.Count == 0 {
		instance.Count = config.Count
	}
	if instance.Duration == 0 {
		instance.Duration = config.Duration
	}

	result, err := mgr.addByStacking(target, &instance, config.Stacking)
	if err != nil {
		return err
	}

	// reset the stats since we are calling OnAdd
	result.Reset(mgr.turnCount)

	f := config.Listeners.OnAdd
	if f != nil {
		f(mgr.engine, result)
	}
	mgr.engine.Events().ModifierAdded.Emit(event.ModifierAddedEvent{
		Target:   target,
		Modifier: result,
		Chance:   chance,
	})
	return nil
}

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

func (mgr *Manager) addByStacking(
	t key.TargetID, i *info.ModifierInstance, s StackingBehavior) (*info.ModifierInstance, error) {
	switch s {
	case ReplaceBySource:
		return mgr.replaceBySource(t, i), nil
	case Replace:
		return mgr.replace(t, i), nil
	case Multiple:
		return mgr.multiple(t, i), nil
	case Refresh:
		return mgr.refresh(t, i), nil
	case Prolong:
		return mgr.prolong(t, i), nil
	case Merge:
		return mgr.merge(t, i), nil
	default:
		return nil, fmt.Errorf("unsupported stacking method: %v", s)
	}
}

// stacks modifiers based on matching modifier name + source. Can have a stack of modifiers from
// each source on the target. When stacking, will use incoming instance rather than keeping old
func (mgr *Manager) replaceBySource(target key.TargetID, instance *info.ModifierInstance) *info.ModifierInstance {
	for i, mod := range mgr.targets[target] {
		if mod.Name == instance.Name && mod.Source == instance.Source {
			// found a matching modifier, attempt to stack
			count := instance.CountAddWhenStack + mod.Count
			if instance.MaxCount > 0 && count > instance.MaxCount {
				count = instance.MaxCount
			}
			instance.Count = count

			// replace means this added instance is the new instance (can have new param values)
			mgr.targets[target][i] = instance
			return instance
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

// stacks modifiers based on matching modifier name. Impossible to have duplicate modifiers on
// the given target. When stacking, will use incoming instance rather than keeping old
func (mgr *Manager) replace(target key.TargetID, instance *info.ModifierInstance) *info.ModifierInstance {
	for i, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, attempt to stack
			count := instance.CountAddWhenStack + mod.Count
			if instance.MaxCount > 0 && count > instance.MaxCount {
				count = instance.MaxCount
			}
			instance.Count = count

			// replace means this added instance is the new instance (can have new param values)
			mgr.targets[target][i] = instance
			return instance
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

// multiple does not check for duplicates and will just add this instance as a new modifier
func (mgr *Manager) multiple(target key.TargetID, instance *info.ModifierInstance) *info.ModifierInstance {
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

// if there is an existing modifier (by name), will KEEP that original instance and only reset
// the duration based on the incoming instance.
func (mgr *Manager) refresh(target key.TargetID, instance *info.ModifierInstance) *info.ModifierInstance {
	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, reset this modifier's duration
			mod.Duration = instance.Duration
			return mod
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

// if there is an existing modifier (by name), will KEEP that original instance and prolong
// this modifiers duration by adding this instance's duration to it
func (mgr *Manager) prolong(target key.TargetID, instance *info.ModifierInstance) *info.ModifierInstance {
	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, prolong the duration
			mod.Duration += instance.Duration
			return mod
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}

// if there is an existing modifier (by name), will KEEP that original instance and merge the new
// instance by performing the "replace" stack logic + using the max duration between old and new
func (mgr *Manager) merge(target key.TargetID, instance *info.ModifierInstance) *info.ModifierInstance {
	for _, mod := range mgr.targets[target] {
		if mod.Name == instance.Name {
			// found a matching modifier, merge
			count := instance.CountAddWhenStack + mod.Count
			if instance.MaxCount > 0 && count > instance.MaxCount {
				count = instance.MaxCount
			}
			mod.Count = count

			if instance.Duration > mod.Duration {
				mod.Duration = instance.Duration
			}
			return mod
		}
	}
	// no match, add this instance as a new modifier
	mgr.targets[target] = append(mgr.targets[target], instance)
	return instance
}
