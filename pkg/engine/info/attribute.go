package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Attributes struct {
	BaseStats     PropMap
	BaseDebuffRES DebuffRESMap
	HPRatio       float64
	Energy        float64
	MaxEnergy     float64
	Stance        float64
	MaxStance     float64
}

type ModifyHPByRatio struct {
	// The amount of HP ratio to modify the HP by (negative will remove HP)
	Ratio float64
	// What ratio type should be used (should Ratio be based on MaxHP or CurrentHP)
	RatioType model.ModifyHPRatioType
	// The floor for how low HP can go with this modification. IE: Floor = 1 will prevent the HP
	// from reaching 0 in this modification (can reduce up to 1 HP)
	Floor float64
}

// A snapshot of a targets stats at a point in time
type Stats struct {
	id           key.TargetID
	currentHP    float64
	energy       float64
	maxEnergy    float64
	stance       float64
	maxStance    float64
	props        PropMap
	debuffRES    DebuffRESMap
	flags        []model.BehaviorFlag
	statusCounts map[model.StatusType]int
	modifiers    []key.Modifier
}

// TODO: ToProto method for logging

func NewStats(id key.TargetID, attributes *Attributes, mods ModifierState) *Stats {
	mods.Props.AddAll(attributes.BaseStats)
	mods.DebuffRES.AddAll(attributes.BaseDebuffRES)
	return &Stats{
		id:           id,
		currentHP:    attributes.HPRatio,
		energy:       attributes.Energy,
		maxEnergy:    attributes.MaxEnergy,
		stance:       attributes.Stance,
		maxStance:    attributes.MaxStance,
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
func (stats *Stats) AddProperty(prop model.Property, amt float64) {
	stats.props.Modify(prop, amt)
}

// Adds a debuff RES to this Stats snapshot
func (stats *Stats) AddDebuffRES(flag model.BehaviorFlag, amt float64) {
	stats.debuffRES.Modify(flag, amt)
}

// Get the current value of a property within this Stats snapshot
func (stats *Stats) GetProperty(prop model.Property) float64 {
	return stats.props[prop]
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

func (stats *Stats) CurrentHP() float64 {
	return stats.currentHP * stats.MaxHP()
}

func (stats *Stats) CurrentHPRatio() float64 {
	return stats.currentHP
}

func (stats *Stats) Stance() float64 {
	return stats.stance
}

func (stats *Stats) MaxStance() float64 {
	return stats.maxStance
}

func (stats *Stats) Energy() float64 {
	return stats.energy
}

func (stats *Stats) MaxEnergy() float64 {
	return stats.maxEnergy
}

// HP = HP_BASE * (1 + HP_PERCENT) + HP_FLAT + HP_CONVERT
func (stats *Stats) MaxHP() float64 {
	return statCalc(
		stats.props[model.Property_HP_BASE],
		stats.props[model.Property_HP_PERCENT],
		stats.props[model.Property_HP_FLAT]+stats.props[model.Property_HP_CONVERT])
}

// HP = HP_BASE * (1 + HP_PERCENT) + HP_FLAT + HP_CONVERT
// alias for MaxHP()
func (stats *Stats) HP() float64 {
	return stats.MaxHP()
}

// ATK = ATK_BASE * (1 + ATK_PERCENT) + ATK_FLAT + ATK_CONVERT
func (stats *Stats) ATK() float64 {
	return statCalc(
		stats.props[model.Property_ATK_BASE],
		stats.props[model.Property_ATK_PERCENT],
		stats.props[model.Property_ATK_FLAT]+stats.props[model.Property_ATK_CONVERT])
}

// DEF = DEF_BASE * (1 + DEF_PERCENT) + DEF_FLAT + DEF_CONVERT
func (stats *Stats) DEF() float64 {
	return statCalc(
		stats.props[model.Property_DEF_BASE],
		stats.props[model.Property_DEF_PERCENT],
		stats.props[model.Property_DEF_FLAT]+stats.props[model.Property_DEF_CONVERT])
}

// SPD = SPD_BASE * (1 + SPD_PERCENT) + SPD_FLAT + SPD_CONVERT
func (stats *Stats) SPD() float64 {
	return statCalc(
		stats.props[model.Property_SPD_BASE],
		stats.props[model.Property_SPD_PERCENT],
		stats.props[model.Property_SPD_FLAT]+stats.props[model.Property_SPD_CONVERT])
}

// AGGRO = AGGRO_BASE * (1 + AGGRO_PERCENT) + AGGRO_FLAT
func (stats *Stats) Aggro() float64 {
	return statCalc(
		stats.props[model.Property_AGGRO_BASE],
		stats.props[model.Property_AGGRO_PERCENT],
		stats.props[model.Property_AGGRO_FLAT])
}

// CRIT CHANCE = CRIT_CHANCE + CRIT_CHANCE_CONVERT
func (stats *Stats) CritChance() float64 {
	return stats.props[model.Property_CRIT_CHANCE]
}

// CRIT DAMAGE = CRIT_DMG + CRIT_DMG_CONVERT
func (stats *Stats) CritDamage() float64 {
	return stats.props[model.Property_CRIT_DMG]
}

// HEAL BOOST = HEAL_BOOST + HEAL_BOOST_CONVERT
func (stats *Stats) HealBoost() float64 {
	return stats.props[model.Property_HEAL_BOOST] + stats.props[model.Property_HEAL_BOOST_CONVERT]
}

// EHR = EFFECT_HIT_RATE + EFFECT_HIT_RATE_CONVERT
func (stats *Stats) EffectHitRate() float64 {
	return stats.props[model.Property_EFFECT_HIT_RATE] +
		stats.props[model.Property_EFFECT_HIT_RATE_CONVERT]
}

// EFFECT RES = EFFECT_RES + EFFECT_RES_CONVERT
func (stats *Stats) EffectRES() float64 {
	return stats.props[model.Property_EFFECT_RES] +
		stats.props[model.Property_EFFECT_RES_CONVERT]
}

// REGEN = ENERGY_REGEN + ENERGY_REGEN_CONVERT
func (stats *Stats) EnergyRegen() float64 {
	return stats.props[model.Property_ENERGY_REGEN] + stats.props[model.Property_ENERGY_REGEN_CONVERT]
}

// BREAK EFFECT = BREAK_EFFECT + BREAK_EFFECT_CONVERT
func (stats *Stats) BreakEffect() float64 {
	return stats.props[model.Property_BREAK_EFFECT]
}

func statCalc(base, percent, flat float64) float64 {
	out := base*(1+percent) + flat
	if out < 0 {
		return 0
	}
	return out
}
