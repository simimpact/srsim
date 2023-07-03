package prop

//go:generate stringer -type=Property

import (
	"encoding/json"

	"github.com/simimpact/srsim/pkg/model"
)

type Property int

const (
	Invalid Property = iota
	// HP = HP_BASE * (1 + HP_PERCENT) + HP_FLAT + HP_CONVERT
	HPBase
	HPPercent
	HPFlat
	HPConvert

	// ATK = ATK_BASE * (1 + ATK_PERCENT) + ATK_FLAT + ATK_CONVERT
	ATKBase
	ATKPercent
	ATKFlat
	ATKConvert

	// DEF = DEF_BASE * (1 + DEF_PERCENT) + DEF_FLAT + DEF_CONVERT
	DEFBase
	DEFPercent
	DEFFlat
	DEFConvert

	// SPD = SPD_BASE * (1 + SPD_PERCENT) + SPD_FLAT + SPD_CONVERT
	SPDBase
	SPDPercent
	SPDFlat
	SPDConvert

	// Crit
	CritChance
	CritDMG

	// Energy Regen
	EnergyRegen
	EnergyRegenConvert

	// Effect Hit Rate
	EffectHitRate
	EffectHitRateConvert

	// Effect RES
	EffectRES
	EffectRESConvert

	// Increases heal strength that are created by target
	HealBoost
	HealBoostConvert

	// Increases heal strength that are applied to target
	HealTaken

	// Increases shield strength that are created by target
	ShieldBoost

	// Increases shield strength that are applied to target
	ShieldTaken

	// AGGRO = AGGRO_BASE * (1 + AGGRO_PERCENT) + AGGRO_FLAT
	AggroBase
	AggroPercent
	AggroFlat

	// Break Effect
	BreakEffect

	// Damage Resistances (RES = ALL_DMG_RES + ELEMENT_DMG_RES)
	AllDamageRES
	PhysicalDamageRES
	FireDamageRES
	IceDamageRES
	ThunderDamageRES
	QuantumDamageRES
	ImaginaryDamageRES
	WindDamageRES

	// Elemental Penetrates
	PhysicalPEN
	FirePEN
	IcePEN
	ThunderPEN
	QuantumPEN
	ImaginaryPEN
	WindPEN

	// Damage Taken Boost (TAKEN = ALL_DMG_TAKEN + ELEMENT_DMG_TAKEN)
	AllDamageTaken
	PhysicalDamageTaken
	FireDamageTaken
	IceDamageTaken
	ThunderDamageTaken
	QuantumDamageTaken
	ImaginaryDamageTaken
	WindDamageTaken

	// DMG% increases (DMG% = ALL_DMG% + ELEMENT_DMG% + DOT_DMG%)
	AllDamagePercent
	DOTDamagePercent
	FireDamagePercent
	IceDamagePercent
	ThunderDamagePercent
	QuantumDamagePercent
	ImaginaryDamagePercent
	WindDamagePercent
	PhysicalDamagePercent

	// Stance DMG% increase (damage to toughness bar, not break effect)
	AllStanceDMGPercent

	// Multiplicative DMG reduction TOTAL_DMG_REDUCE = 1 - (1 - CUR_DMG_REDUCE)*(1 - ADDED_DMG_REDUCE)
	// DMG_REDUCE from target attacked, FATIGUE from attacker
	AllDamageReduce Property = 90
	Fatigue         Property = 91
	MinFatigue      Property = 92
)

func FromProto(p model.Property) Property {
	return Property(p)
}

func (p Property) ToProto() model.Property {
	return model.Property(p)
}

func (p Property) MarshalJSON() ([]byte, error) {
	return json.Marshal(model.Property_name[int32(p)])
}
