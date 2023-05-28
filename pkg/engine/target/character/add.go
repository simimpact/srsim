package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddCharacter(id key.TargetID, char *model.Character) error {
	config, ok := characterCatalog[key.Character(char.Key)]
	if !ok {
		return fmt.Errorf("invalid character: %v", char.Key)
	}

	lvl := int(char.Level)
	asc := config.ascension(int(char.MaxLevel))
	baseStats := newBaseStats(config.Promotions[asc], lvl)
	traces := processTraces(config.Traces, baseStats, char.Traces, asc, lvl)
	baseDebuffRES := info.NewDebuffRESMap()

	err := mgr.attr.AddTarget(id, attribute.BaseStats{
		Stats:     baseStats,
		DebuffRES: baseDebuffRES,
		MaxEnergy: config.MaxEnergy,
	})
	if err != nil {
		return err
	}

	// TODO: lightcone + relic initialization (before or after character init?)

	info := info.Character{
		Key:       key.Character(char.Key),
		Level:     lvl,
		Ascension: asc,
		Eidolon:   int(char.Eidols),
		Path:      config.Path,
		Element:   config.Element,
		BaseStats: baseStats,
		Traces:    traces,
	}

	mgr.info[id] = info
	mgr.instances[id] = config.Create(mgr.engine, id, info)

	// TODO: emit CharacterAddedEvent
	return nil
}

func newBaseStats(data PromotionData, level int) info.PropMap {
	out := info.NewPropMap()
	out.Modify(model.Property_ATK_BASE, data.ATKBase+data.ATKAdd*float64(level-1))
	out.Modify(model.Property_DEF_BASE, data.DEFBase+data.DEFAdd*float64(level-1))
	out.Modify(model.Property_HP_BASE, data.HPBase+data.HPAdd*float64(level-1))
	out.Modify(model.Property_SPD_BASE, data.SPD)
	out.Modify(model.Property_CRIT_CHANCE, data.CritChance)
	out.Modify(model.Property_CRIT_DMG, data.CritDMG)
	out.Modify(model.Property_AGGRO_BASE, data.Aggro)
	return out
}

func processTraces(traces TraceMap, stats info.PropMap, wanted []string, asc int, level int) map[string]bool {
	active := make(map[string]bool)
	for _, id := range wanted {
		if dup := active[id]; dup {
			continue
		}

		trace, ok := traces[id]
		if !ok {
			continue
		}

		if asc < trace.Ascension || level < trace.Level {
			continue
		}

		// mark as an active trace and add to info
		active[id] = true
		if trace.Stat != model.Property_INVALID_PROP {
			stats.Modify(trace.Stat, trace.Amount)
		}
	}
	return active
}
