package info

import (
	"encoding/json"

	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// A snapshot of a targets stats at a point in time
type Stats struct {
	id                 key.TargetID
	props              PropMap
	debuffRES          DebuffRESMap
	weakness           WeaknessMap
	flags              []model.BehaviorFlag
	statusCounts       map[model.StatusType]int
	modifiers          []ModifierChangeSet
	attributes         *Attributes
	propChanges        []propChangeSet
	debuffRESChangeSet []debuffRESChangeSet
}

type propChangeSet struct {
	prop   prop.Property
	key    key.Reason
	amount float64
}

type debuffRESChangeSet struct {
	flag   model.BehaviorFlag
	key    key.Reason
	amount float64
}

func NewStats(id key.TargetID, attributes *Attributes, mods *ModifierState) *Stats {
	mods.Props.AddAll(attributes.BaseStats)
	mods.DebuffRES.AddAll(attributes.BaseDebuffRES)
	mods.Weakness.AddAll(attributes.Weakness)

	return &Stats{
		id:                 id,
		props:              mods.Props,
		debuffRES:          mods.DebuffRES,
		weakness:           mods.Weakness,
		flags:              mods.Flags,
		statusCounts:       mods.Counts,
		modifiers:          mods.Modifiers,
		attributes:         copyAttributes(attributes),
		propChanges:        make([]propChangeSet, 0, 16),
		debuffRESChangeSet: make([]debuffRESChangeSet, 0, 16),
	}
}

func copyAttributes(attributes *Attributes) *Attributes {
	props := make(PropMap, len(attributes.BaseStats))
	props.AddAll(attributes.BaseStats)

	debuffRES := make(DebuffRESMap, len(attributes.BaseDebuffRES))
	debuffRES.AddAll(attributes.BaseDebuffRES)

	weakness := make(WeaknessMap, len(attributes.Weakness))
	weakness.AddAll(attributes.Weakness)

	return &Attributes{
		Level:         attributes.Level,
		BaseStats:     props,
		BaseDebuffRES: debuffRES,
		Weakness:      weakness,
		HPRatio:       attributes.HPRatio,
		Energy:        attributes.Energy,
		MaxEnergy:     attributes.MaxEnergy,
		Stance:        attributes.Stance,
		MaxStance:     attributes.MaxStance,
	}
}

// The targetID for who these stats are for
func (stats *Stats) ID() key.TargetID {
	return stats.id
}

// Adds a property to this Stats snapshot
func (stats *Stats) AddProperty(key key.Reason, p prop.Property, amt float64) {
	stats.propChanges = append(stats.propChanges, propChangeSet{
		key:    key,
		prop:   p,
		amount: amt,
	})
	stats.props.Modify(p, amt)
}

// Adds a debuff RES to this Stats snapshot
func (stats *Stats) AddDebuffRES(key key.Reason, flag model.BehaviorFlag, amt float64) {
	stats.debuffRESChangeSet = append(stats.debuffRESChangeSet, debuffRESChangeSet{
		key:    key,
		flag:   flag,
		amount: amt,
	})
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
func (stats *Stats) Modifiers() []ModifierChangeSet {
	return stats.modifiers
}

// Returns true if this target is weak to the given damage type
func (stats *Stats) IsWeakTo(t model.DamageType) bool {
	return stats.weakness[t]
}

// The current level of the target.
func (stats *Stats) Level() int {
	return stats.attributes.Level
}

// The current HP amount of the target (HPRatio * MaxHP)
func (stats *Stats) CurrentHP() float64 {
	return stats.attributes.HPRatio * stats.MaxHP()
}

// The current HPRatio (value between 0 and 1).
func (stats *Stats) CurrentHPRatio() float64 {
	return stats.attributes.HPRatio
}

// The current Stance/Toughness amount of the target
func (stats *Stats) Stance() float64 {
	return stats.attributes.Stance
}

// The max Stance/Toughness amount of the target
func (stats *Stats) MaxStance() float64 {
	return stats.attributes.MaxStance
}

// The current energy amount of the target
func (stats *Stats) Energy() float64 {
	return stats.attributes.Energy
}

// The max energy amount of the target
func (stats *Stats) MaxEnergy() float64 {
	return stats.attributes.MaxEnergy
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

// DAMAGE_PERCENT = ALL_DAMAGE_PERCENT + DAMAGE_TYPE_PERCENT
func (stats *Stats) DamagePercent(dmg model.DamageType) float64 {
	return stats.props[prop.AllDamagePercent] + stats.props[prop.DamagePercent(dmg)]
}

// DAMAGE_RES = ALL_DAMAGE_RES + DAMAGE_TYPE_RES
func (stats *Stats) DamageRES(dmg model.DamageType) float64 {
	return stats.props[prop.AllDamageRES] + stats.props[prop.DamageRES(dmg)]
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
	Stats        *ComputedStats           `json:"stats"`
	Flags        []model.BehaviorFlag     `json:"flags"`
	StatusCounts map[model.StatusType]int `json:"status_counts"`
	Modifiers    map[key.Modifier]int     `json:"modifiers"`
	Props        []*LoggedProp            `json:"props"`
	DebuffRES    []*LoggedDebuffRES       `json:"debuff_res"`
	Weakness     WeaknessMap              `json:"weakness"`
}

type ComputedStats struct {
	HP                     float64 `json:"hp"`
	ATK                    float64 `json:"atk"`
	DEF                    float64 `json:"def"`
	SPD                    float64 `json:"spd"`
	Aggro                  float64 `json:"aggro"`
	CritChance             float64 `json:"crit_chance"`
	CritDMG                float64 `json:"crit_dmg"`
	HealBoost              float64 `json:"heal_boost"`
	EffectHitRate          float64 `json:"effect_hit_rate"`
	EffectRES              float64 `json:"effect_res"`
	EnergyRegen            float64 `json:"energy_regen"`
	BreakEffect            float64 `json:"break_effect"`
	PhysicalDamagePercent  float64 `json:"physical_damage_percent"`
	FireDamagePercent      float64 `json:"fire_damage_percent"`
	IceDamagePercent       float64 `json:"ice_damage_percent"`
	LightningDamagePercent float64 `json:"lightning_damage_percent"`
	WindDamagePercent      float64 `json:"wind_damage_percent"`
	QuantumDamagePercent   float64 `json:"quantum_damage_percent"`
	ImaginaryDamagePercent float64 `json:"imaginary_damage_percent"`
	PhysicalRES            float64 `json:"physical_res"`
	FireRES                float64 `json:"fire_res"`
	IceRES                 float64 `json:"ice_res"`
	LightningRES           float64 `json:"lightning_res"`
	WindRES                float64 `json:"wind_res"`
	QuantumRES             float64 `json:"quantum_res"`
	ImaginaryRES           float64 `json:"imaginary_res"`
}

type LoggedProp struct {
	Prop    prop.Property    `json:"prop"`
	Total   float64          `json:"total"`
	Sources []FloatChangeSet `json:"sources"`
}

type LoggedDebuffRES struct {
	Flag    model.BehaviorFlag `json:"flag"`
	Total   float64            `json:"total"`
	Sources []FloatChangeSet   `json:"sources"`
}

type FloatChangeSet struct {
	Key    key.Reason `json:"key"`
	Amount float64    `json:"amount"`
}

func (stats *Stats) MarshalJSON() ([]byte, error) {
	modOccurrences := make(map[key.Modifier]int, len(stats.modifiers))
	props := make(map[prop.Property]*LoggedProp, prop.Total())
	debuffRES := make(map[model.BehaviorFlag]*LoggedDebuffRES, len(model.BehaviorFlag_name))

	for p, amt := range stats.props {
		props[p] = &LoggedProp{
			Prop:    p,
			Total:   amt,
			Sources: make([]FloatChangeSet, 0, 8),
		}
	}

	for f, amt := range stats.debuffRES {
		debuffRES[f] = &LoggedDebuffRES{
			Flag:    f,
			Total:   amt,
			Sources: make([]FloatChangeSet, 0, 8),
		}
	}

	for p, amt := range stats.attributes.BaseStats {
		props[p].Sources = append(props[p].Sources, FloatChangeSet{
			Key:    "base",
			Amount: amt,
		})
	}

	for f, amt := range stats.attributes.BaseDebuffRES {
		debuffRES[f].Sources = append(debuffRES[f].Sources, FloatChangeSet{
			Key:    "base",
			Amount: amt,
		})
	}

	for _, m := range stats.modifiers {
		modOccurrences[m.Name] += 1

		for p, amt := range m.Props {
			props[p].Sources = append(props[p].Sources, FloatChangeSet{
				Key:    key.Reason(m.Name),
				Amount: amt,
			})
		}

		for f, amt := range m.DebuffRES {
			debuffRES[f].Sources = append(debuffRES[f].Sources, FloatChangeSet{
				Key:    key.Reason(m.Name),
				Amount: amt,
			})
		}
	}

	for _, change := range stats.propChanges {
		props[change.prop].Sources = append(props[change.prop].Sources, FloatChangeSet{
			Key:    change.key,
			Amount: change.amount,
		})
	}

	for _, change := range stats.debuffRESChangeSet {
		debuffRES[change.flag].Sources = append(debuffRES[change.flag].Sources, FloatChangeSet{
			Key:    change.key,
			Amount: change.amount,
		})
	}

	loggedProps := make([]*LoggedProp, 0, len(props))
	for _, v := range props {
		loggedProps = append(loggedProps, v)
	}

	loggedDebuffRES := make([]*LoggedDebuffRES, 0, len(debuffRES))
	for _, v := range debuffRES {
		loggedDebuffRES = append(loggedDebuffRES, v)
	}

	out := StatsEncoded{
		ID:           stats.ID(),
		HPRatio:      stats.CurrentHPRatio(),
		Energy:       stats.Energy(),
		Stance:       stats.Stance(),
		Props:        loggedProps,
		DebuffRES:    loggedDebuffRES,
		Weakness:     stats.weakness,
		Flags:        stats.flags,
		StatusCounts: stats.statusCounts,
		Modifiers:    modOccurrences,
		Stats: &ComputedStats{
			// base stats
			HP:  stats.HP(),
			ATK: stats.ATK(),
			DEF: stats.DEF(),
			SPD: stats.SPD(),

			// advanced stats
			CritChance:    stats.CritChance(),
			CritDMG:       stats.CritDamage(),
			BreakEffect:   stats.BreakEffect(),
			HealBoost:     stats.HealBoost(),
			EnergyRegen:   stats.EnergyRegen(),
			EffectHitRate: stats.EffectHitRate(),
			EffectRES:     stats.EffectRES(),

			// dmg type
			PhysicalDamagePercent:  stats.DamagePercent(model.DamageType_PHYSICAL),
			FireDamagePercent:      stats.DamagePercent(model.DamageType_FIRE),
			IceDamagePercent:       stats.DamagePercent(model.DamageType_ICE),
			LightningDamagePercent: stats.DamagePercent(model.DamageType_THUNDER),
			WindDamagePercent:      stats.DamagePercent(model.DamageType_WIND),
			QuantumDamagePercent:   stats.DamagePercent(model.DamageType_QUANTUM),
			ImaginaryDamagePercent: stats.DamagePercent(model.DamageType_IMAGINARY),
			PhysicalRES:            stats.DamageRES(model.DamageType_PHYSICAL),
			FireRES:                stats.DamageRES(model.DamageType_FIRE),
			IceRES:                 stats.DamageRES(model.DamageType_ICE),
			LightningRES:           stats.DamageRES(model.DamageType_THUNDER),
			WindRES:                stats.DamageRES(model.DamageType_WIND),
			QuantumRES:             stats.DamageRES(model.DamageType_QUANTUM),
			ImaginaryRES:           stats.DamageRES(model.DamageType_IMAGINARY),

			// hidden stats
			Aggro: stats.Aggro(),
		},
	}
	return json.Marshal(out)
}
