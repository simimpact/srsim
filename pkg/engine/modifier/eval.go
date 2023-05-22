package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// TODO: move to info?

// Evaluate all the modifiers for a given target and return their current modifier state:
//   - sum of all applied properties/stats
//   - current Debuff RES
//   - all active behavior flags
//   - count of active buffs & debuffs
//   - list of all active modifiers (deduped)
func (mgr *Manager) EvalModifiers(target key.TargetID) info.ModifierState {
	props := info.NewPropMap()
	debuffRES := info.NewDebuffRESMap()
	counts := make(map[model.StatusType]int)
	flagSet := make(map[model.BehaviorFlag]struct{})
	modSet := make(map[key.Modifier]struct{})

	for _, mod := range mgr.targets[target] {
		props.AddAllFromInstance(mod)
		debuffRES.AddAllFromInstance(mod)
		counts[modifierCatalog[mod.Name].StatusType] += 1
		modSet[mod.Name] = struct{}{}

		for _, flag := range modifierCatalog[mod.Name].BehaviorFlags {
			flagSet[flag] = struct{}{}
		}
	}

	return info.ModifierState{
		Props:     props,
		DebuffRES: debuffRES,
		Counts:    counts,
		Flags:     toList(flagSet),
		Modifiers: toList(modSet),
	}
}

func toList[T comparable](m map[T]struct{}) []T {
	out := make([]T, len(m))
	i := 0
	for k := range m {
		out[i] = k
		i++
	}
	return out
}
