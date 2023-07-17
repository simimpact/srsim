package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddCharacter(id key.TargetID, char *model.Character) error {
	config, ok := characterCatalog[key.Character(char.Key)]
	if !ok {
		return fmt.Errorf("invalid character: %v", char.Key)
	}

	lvl := int(char.Level)
	if lvl <= 0 {
		lvl = 1
	}
	asc := config.ascension(int(char.MaxLevel), lvl)

	// add char base stats from curve + traces
	baseStats := newBaseStats(config.Promotions[asc], lvl)
	traces := processTraces(config.Traces, baseStats, char.Traces, asc, lvl)

	// add lightcone base stats
	lcLvl := int(char.LightCone.Level)
	if lcLvl <= 0 {
		lcLvl = 1
	}

	lcConfig, err := lightcone.Get(key.LightCone(char.LightCone.Key))
	if err != nil {
		return err
	}
	lcAsc := lcConfig.Ascension(int(char.LightCone.MaxLevel), lcLvl)
	lightcone.AddBaseStats(baseStats, lcConfig.Promotions[lcAsc], lcLvl)

	// add relic stats from sim config
	relics := make(map[key.Relic]int)
	for _, r := range char.Relics {
		relics[key.Relic(r.Key)] += 1
		baseStats.Modify(prop.FromProto(r.MainStat.Stat), r.MainStat.Amount)
		for _, sub := range r.SubStats {
			baseStats.Modify(prop.FromProto(sub.Stat), sub.Amount)
		}
	}

	// add relic stats from relic config + get list of callbacks to call later
	var relicCBs []relic.CreateEffectFunc
	for r, count := range relics {
		config, err := relic.Get(r)
		if err != nil {
			return err
		}

		for _, effect := range config.Effects {
			if count < effect.MinCount {
				continue
			}
			baseStats.AddAll(effect.Stats)
			if effect.CreateEffect != nil {
				relicCBs = append(relicCBs, effect.CreateEffect)
			}
		}
	}

	hp := char.StartHp
	if hp <= 0 || hp > 1.0 {
		hp = 1.0
	}

	// Give the base stats to the attribute manager so Stats calls can work as expected
	err = mgr.attr.AddTarget(id, info.Attributes{
		Level:         lvl,
		BaseStats:     baseStats,
		MaxEnergy:     config.MaxEnergy,
		Energy:        char.StartEnergy,
		BaseDebuffRES: info.NewDebuffRESMap(),
		Weakness:      info.NewWeaknessMap(),
		HPRatio:       hp,
		Stance:        0,
		MaxStance:     0,
	})
	if err != nil {
		return err
	}

	info := info.Character{
		Key:          key.Character(char.Key),
		Level:        lvl,
		Ascension:    asc,
		Eidolon:      int(char.Eidols),
		Path:         config.Path,
		Element:      config.Element,
		Traces:       traces,
		AbilityLevel: abilityLevels(char.Abilities),
		LightCone: info.LightCone{
			Key:        key.LightCone(char.LightCone.Key),
			Level:      lcLvl,
			Ascension:  lcAsc,
			Imposition: int(char.LightCone.Imposition),
			Path:       lcConfig.Path,
		},
		Relics: relics,
	}

	mgr.info[id] = info
	mgr.instances[id] = config.Create(mgr.engine, id, info)

	// only create lightcone passive iff paths match
	if config.Path == lcConfig.Path {
		lcConfig.CreatePassive(mgr.engine, id, info.LightCone)
	}

	// Call each relic CB
	for _, f := range relicCBs {
		f(mgr.engine, id)
	}
	return nil
}

func newBaseStats(data PromotionData, level int) info.PropMap {
	out := info.NewPropMap()
	out.Modify(prop.ATKBase, data.ATKBase+data.ATKAdd*float64(level-1))
	out.Modify(prop.DEFBase, data.DEFBase+data.DEFAdd*float64(level-1))
	out.Modify(prop.HPBase, data.HPBase+data.HPAdd*float64(level-1))
	out.Modify(prop.SPDBase, data.SPD)
	out.Modify(prop.AggroBase, data.Aggro)
	out.Modify(prop.CritChance, data.CritChance)
	out.Modify(prop.CritDMG, data.CritDMG)
	return out
}

func processTraces(traces TraceMap, stats info.PropMap, wanted []string, asc, level int) map[string]bool {
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
		if trace.Stat != prop.Invalid {
			stats.Modify(trace.Stat, trace.Amount)
		}
	}
	return active
}

func abilityLevels(levels *model.Abilities) info.AbilityLevel {
	if levels == nil {
		return info.AbilityLevel{Attack: 1, Skill: 1, Ult: 1, Talent: 1}
	}

	out := info.AbilityLevel{
		Attack: limitAbilityLevel(int(levels.Attack), 9),
		Skill:  limitAbilityLevel(int(levels.Skill), 15),
		Ult:    limitAbilityLevel(int(levels.Ult), 15),
		Talent: limitAbilityLevel(int(levels.Talent), 15),
	}
	return out
}

func limitAbilityLevel(lvl, max int) int {
	if lvl < 1 {
		return 1
	} else if lvl > max {
		return max
	}
	return lvl
}
