package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
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

	// add char base stats from curve + traces
	baseDebuffRES := info.NewDebuffRESMap()
	baseStats := newBaseStats(config.Promotions[asc], lvl)
	traces := processTraces(config.Traces, baseStats, char.Traces, asc, lvl)

	// add lightcone base stats
	lcLvl := int(char.Cone.Level)
	lcConfig, err := lightcone.Get(key.LightCone(char.Cone.Key))
	if err != nil {
		return err
	}
	lcAsc := lcConfig.Ascension(int(char.Cone.MaxLevel))
	lightcone.AddBaseStats(baseStats, lcConfig.Promotions[lcAsc], lcLvl)

	// TODO: relic base stats

	err = mgr.attr.AddTarget(id, attribute.BaseStats{
		Stats:     baseStats,
		DebuffRES: baseDebuffRES,
		MaxEnergy: config.MaxEnergy,
	})
	if err != nil {
		return err
	}

	// TODO: relic initialization

	info := info.Character{
		Key:       key.Character(char.Key),
		Level:     lvl,
		Ascension: asc,
		Eidolon:   int(char.Eidols),
		Path:      config.Path,
		Element:   config.Element,
		BaseStats: baseStats,
		Traces:    traces,
		LightCone: info.LightCone{
			Key:       key.LightCone(char.Cone.Key),
			Level:     lcLvl,
			Ascension: lcAsc,
			Rank:      int(char.Cone.Imposition),
			Path:      lcConfig.Path,
		},
	}

	mgr.info[id] = info
	mgr.instances[id] = config.Create(mgr.engine, id, info)

	// only create lightcone passive iff paths match
	if config.Path == lcConfig.Path {
		lcConfig.CreatePassive(mgr.engine, id, info.LightCone)
	}

	// TODO: relic create call

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
