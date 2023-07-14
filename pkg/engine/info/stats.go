package info

import (
	"encoding/json"

	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A snapshot of a targets stats at a point in time
type Stats struct {
	id           key.TargetID
	level        int
	currentHP    float64
	energy       float64
	maxEnergy    float64
	stance       float64
	maxStance    float64
	props        PropMap
	debuffRES    DebuffRESMap
	weakness     WeaknessMap
	flags        []model.BehaviorFlag
	statusCounts map[model.StatusType]int
	modifiers    []key.Modifier
}

// TODO: ToProto method for logging

func NewStats(id key.TargetID, attributes *Attributes, mods ModifierState) *Stats {
	mods.Props.AddAll(attributes.BaseStats)
	mods.DebuffRES.AddAll(attributes.BaseDebuffRES)
	mods.Weakness.AddAll(attributes.Weakness)
	// TODO: merge weaknesses between attributes + mods for cases of weakness app like Silver Wolf
	return &Stats{
		id:           id,
		level:        attributes.Level,
		currentHP:    attributes.HPRatio,
		energy:       attributes.Energy,
		maxEnergy:    attributes.MaxEnergy,
		stance:       attributes.Stance,
		maxStance:    attributes.MaxStance,
		weakness:     attributes.Weakness,
		props:        mods.Props,
		debuffRES:    mods.DebuffRES,
		flags:        mods.Flags,
		statusCounts: mods.Counts,
		modifiers:    mods.Modifiers,
	}
}

// The targetID for who these stats are for
func (stats *Stats) ID() key.TargetID {
	return stats.id
}

// Adds a property to this Stats snapshot
func (stats *Stats) AddProperty(p prop.Property, amt float64) {
	stats.props.Modify(p, amt)
}

// Adds a debuff RES to this Stats snapshot
func (stats *Stats) AddDebuffRES(flag model.BehaviorFlag, amt float64) {
	stats.debuffRES.Modify(flag, amt)
}

// Get the current value of a property within this Stats snapshot
func (stats *Stats) GetProperty(p prop.Property) float64 {
	return stats.props[p]
}

// Get the debuff RES for a given set of behavior flagss
func (stats *Stats) GetDebuffRES(flags ...model.BehaviorFlag) float64 {
	return stats.debuffRES.GetDebuffRES(flags...)
}

// Checks if this Stats snapshot has at least one of the given behavior flags associated with it
func (stats *Stats) HasBehaviorFlag(flags ...model.BehaviorFlag) bool {
	for _, flag := range flags {
		for _, sf := range stats.flags {
			if sf == flag {
				return true
			}
		}
	}
	return false
}

// The number of statuses that are associated with this Stats snapshot (Buffs/Debuffs count)
func (stats *Stats) StatusCount(status model.StatusType) int {
	return stats.statusCounts[status]
}

// A list of all modifiers that were used to populate this Stats
func (stats *Stats) Modifiers() []key.Modifier {
	return stats.modifiers
}

// Returns true if this target is weak to the given damage type
func (stats *Stats) IsWeakTo(t model.DamageType) bool {
	return stats.weakness[t]
}

// The current level of the target.
func (stats *Stats) Level() int {
	return stats.level
}

// The current HP amount of the target (HPRatio * MaxHP)
func (stats *Stats) CurrentHP() float64 {
	return stats.currentHP * stats.MaxHP()
}

// The current HPRatio (value between 0 and 1).
func (stats *Stats) CurrentHPRatio() float64 {
	return stats.currentHP
}

// The current Stance/Toughness amount of the target
func (stats *Stats) Stance() float64 {
	return stats.stance
}

// The max Stance/Toughness amount of the target
func (stats *Stats) MaxStance() float64 {
	return stats.maxStance
}

// The current energy amount of the target
func (stats *Stats) Energy() float64 {
	return stats.energy
}

// The max energy amount of the target
func (stats *Stats) MaxEnergy() float64 {
	return stats.maxEnergy
}

// HP = HP_BASE * (1 + HP_PERCENT) + HP_FLAT + HP_CONVERT
func (stats *Stats) MaxHP() float64 {
	return statCalc(
		stats.props[prop.HPBase],
		stats.props[prop.HPPercent],
		stats.props[prop.HPFlat]+stats.props[prop.HPConvert])
}

// HP = HP_BASE * (1 + HP_PERCENT) + HP_FLAT + HP_CONVERT
// alias for MaxHP()
func (stats *Stats) HP() float64 {
	return stats.MaxHP()
}

// ATK = ATK_BASE * (1 + ATK_PERCENT) + ATK_FLAT + ATK_CONVERT
func (stats *Stats) ATK() float64 {
	return statCalc(
		stats.props[prop.ATKBase],
		stats.props[prop.ATKPercent],
		stats.props[prop.ATKFlat]+stats.props[prop.ATKConvert])
}

// DEF = DEF_BASE * (1 + DEF_PERCENT) + DEF_FLAT + DEF_CONVERT
func (stats *Stats) DEF() float64 {
	return statCalc(
		stats.props[prop.DEFBase],
		stats.props[prop.DEFPercent],
		stats.props[prop.DEFFlat]+stats.props[prop.DEFConvert])
}

// SPD = SPD_BASE * (1 + SPD_PERCENT) + SPD_FLAT + SPD_CONVERT
func (stats *Stats) SPD() float64 {
	return statCalc(
		stats.props[prop.SPDBase],
		stats.props[prop.SPDPercent],
		stats.props[prop.SPDFlat]+stats.props[prop.SPDConvert])
}

// AGGRO = AGGRO_BASE * (1 + AGGRO_PERCENT) + AGGRO_FLAT
func (stats *Stats) Aggro() float64 {
	return statCalc(
		stats.props[prop.AggroBase],
		stats.props[prop.AggroPercent],
		stats.props[prop.AggroFlat])
}

// CRIT CHANCE = CRIT_CHANCE + CRIT_CHANCE_CONVERT
func (stats *Stats) CritChance() float64 {
	return stats.props[prop.CritChance]
}

// CRIT DAMAGE = CRIT_DMG + CRIT_DMG_CONVERT
func (stats *Stats) CritDamage() float64 {
	return stats.props[prop.CritDMG]
}

// HEAL BOOST = HEAL_BOOST + HEAL_BOOST_CONVERT
func (stats *Stats) HealBoost() float64 {
	return stats.props[prop.HealBoost] + stats.props[prop.HealBoostConvert]
}

// EHR = EFFECT_HIT_RATE + EFFECT_HIT_RATE_CONVERT
func (stats *Stats) EffectHitRate() float64 {
	return stats.props[prop.EffectHitRate] + stats.props[prop.EffectHitRateConvert]
}

// EFFECT RES = EFFECT_RES + EFFECT_RES_CONVERT
func (stats *Stats) EffectRES() float64 {
	return stats.props[prop.EffectRES] + stats.props[prop.EffectRESConvert]
}

// REGEN = ENERGY_REGEN + ENERGY_REGEN_CONVERT
func (stats *Stats) EnergyRegen() float64 {
	return stats.props[prop.EnergyRegen] + stats.props[prop.EnergyRegenConvert]
}

// BREAK EFFECT = BREAK_EFFECT + BREAK_EFFECT_CONVERT
func (stats *Stats) BreakEffect() float64 {
	return stats.props[prop.BreakEffect]
}

func statCalc(base, percent, flat float64) float64 {
	out := base*(1+percent) + flat
	if out < 0 {
		return 0
	}
	return out
}

type StatsEncoded struct {
	ID           key.TargetID             `json:"id"`
	HPRatio      float64                  `json:"hp_ratio"`
	Energy       float64                  `json:"energy"`
	Stance       float64                  `json:"stance"`
	Props        PropMap                  `json:"props"`
	DebuffRES    DebuffRESMap             `json:"debuff_res"`
	Weakness     WeaknessMap              `json:"weakness"`
	Flags        []model.BehaviorFlag     `json:"flags"`
	StatusCounts map[model.StatusType]int `json:"status_counts"`
	Modifiers    []key.Modifier           `json:"modifiers"`
	Stats        *ComputedStats           `json:"stats"`
}

type ComputedStats struct {
	HP            float64 `json:"hp"`
	ATK           float64 `json:"atk"`
	DEF           float64 `json:"def"`
	SPD           float64 `json:"spd"`
	Aggro         float64 `json:"aggro"`
	CritChance    float64 `json:"crit_chance"`
	CritDMG       float64 `json:"crit_dmg"`
	HealBoost     float64 `json:"heal_boost"`
	EffectHitRate float64 `json:"effect_hit_rate"`
	EffectRES     float64 `json:"effect_res"`
	EnergyRegen   float64 `json:"energy_regen"`
	BreakEffect   float64 `json:"break_effect"`
}

func (stats *Stats) MarshalJSON() ([]byte, error) {
	out := StatsEncoded{
		ID:           stats.ID(),
		HPRatio:      stats.CurrentHPRatio(),
		Energy:       stats.Energy(),
		Stance:       stats.Stance(),
		DebuffRES:    stats.debuffRES,
		Weakness:     stats.weakness,
		Flags:        stats.flags,
		StatusCounts: stats.statusCounts,
		Modifiers:    stats.Modifiers(),
		Props:        stats.props,
		Stats: &ComputedStats{
			HP:            stats.HP(),
			ATK:           stats.ATK(),
			DEF:           stats.DEF(),
			SPD:           stats.SPD(),
			CritChance:    stats.CritChance(),
			CritDMG:       stats.CritDamage(),
			HealBoost:     stats.HealBoost(),
			EffectHitRate: stats.EffectHitRate(),
			EffectRES:     stats.EffectRES(),
			EnergyRegen:   stats.EnergyRegen(),
			BreakEffect:   stats.BreakEffect(),
			Aggro:         stats.Aggro(),
		},
	}
	return json.Marshal(out)
}
