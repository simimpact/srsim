package modifier

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Evaluate all the modifiers for a given target and return their current modifier state:
//   - sum of all applied properties/stats
//   - current Debuff RES
//   - all active behavior flags
//   - count of active buffs & debuffs
//   - list of all active modifiers (deduped)
func (mgr *Manager) EvalModifiers(target key.TargetID) *info.ModifierState {
	totalProps := make(info.PropMap, prop.Total())
	totalDebuffRES := make(info.DebuffRESMap, len(model.BehaviorFlag_name))
	totalWeakness := make(info.WeaknessMap, len(model.DamageFormula_name))

	counts := make(map[model.StatusType]int)
	flagSet := make(map[model.BehaviorFlag]struct{})
	mods := make([]info.ModifierChangeSet, 0, len(mgr.targets[target]))

	for _, mod := range mgr.targets[target] {
		props := make(info.PropMap, len(mod.stats))
		props.AddAll(mod.stats)
		totalProps.AddAll(mod.stats)

		debuffRES := make(info.DebuffRESMap, len(mod.debuffRES))
		debuffRES.AddAll(mod.debuffRES)
		totalDebuffRES.AddAll(mod.debuffRES)

		weakness := make(info.WeaknessMap, len(mod.weakness))
		weakness.AddAll(mod.weakness)
		totalWeakness.AddAll(mod.weakness)

		counts[mod.statusType] += 1
		for _, flag := range mod.flags {
			flagSet[flag] = struct{}{}
		}

		mods = append(mods, info.ModifierChangeSet{
			Name:      mod.name,
			Props:     props,
			DebuffRES: debuffRES,
			Weakness:  weakness,
		})
	}

	return &info.ModifierState{
		Props:     totalProps,
		DebuffRES: totalDebuffRES,
		Weakness:  totalWeakness,
		Modifiers: mods,
		Counts:    counts,
		Flags:     toList(flagSet),
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
