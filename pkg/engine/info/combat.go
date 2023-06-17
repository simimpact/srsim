package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type DamageMap map[model.DamageFormula]float64
type HealMap map[model.HealFormula]float64

type Attack struct {
	// List of targets to perform this attack against (will perform 1 hit per target)
	Targets []key.TargetID

	// The source target which is performing this attack
	Source key.TargetID

	// The type of attack (IE: dot, ult, insert, etc)
	AttackType model.AttackType

	// The damage type of this attack (physical, fire, ice, etc)
	DamageType model.DamageType

	// Map of damage formula -> damage percentage. This is for calculating the "Base Damage" of the
	// attack. IE: info.DamageMap{model.BY_ATK: 0.5} = 50% of ATK.
	//
	// If HitRatio is specified, all base damage multiplier will be multiplied by the hit ratio.
	// IE: info.DamageMap{model.BY_ATK: 0.5} w/ 0.5 HitRatio means 25% of ATK.
	BaseDamage DamageMap

	// How much energy will be generated for the source from this attack. This energy generation
	// will scale with ERR.
	//
	// If HitRatio is specified, the energy gained will be multiplied by the hit ratio.
	// IE: 30.0 EnergyGain with a 0.5 HitRatio means only 15.0 energy added (before ERR bonus)
	EnergyGain float64

	// How much stance/toughness damage this attack will deal. This stance damage will scale with
	// Stance DMG% increase.
	//
	// If HitRatio is specified, the stance damage will be multiplied by the hit ratio.
	// IE: 30.0 StanceDamage with a 0.5 HitRatio means only 15 stance dmage will occur (before bonus)
	StanceDamage float64

	// Hit ratio reduces the BaseDamage, EnergyGain, and StanceDamage by the given percentage. This
	// is used for attacks that perform multiple hits. It is expected that the sum of all HitRatio
	// used for all hits in an attack equal 1.0 (IE: 2 attacks w/ HitRatio of 0.45 & 0.55)
	HitRatio float64

	// If true, will use the "pure damage" formula. This removes some variables from the damage
	// equation, such as crit.
	AsPureDamage bool

	// An additional flat damage amount that can be added to the base damage
	DamageValue float64

	// If set to true, will execute this attack in a "snapshot" state. This means that any modifiers
	// that subscribe to hit listeners will not be executed. This is used by break damage dots.
	UseSnapshot bool
}

type Hit struct {
	// The stats of the attacker of this hit. These stats are a snapshot of the target's state and
	// can be modified
	Attacker *Stats

	// The stats of the defender of this hit. These stats are a snapshot of the target's state and
	// can be modified
	Defender *Stats

	// The type of attack (IE: dot, ult, insert, etc)
	AttackType model.AttackType

	// The damage type of this attack (physical, fire, ice, etc)
	DamageType model.DamageType

	// Map of damage formula -> damage percentage. This is for calculating the "Base Damage" of the
	// attack. IE: info.DamageMap{model.BY_ATK: 0.5} = 50% of ATK.
	//
	// If HitRatio is specified, all base damage multiplier will be multiplied by the hit ratio.
	// IE: info.DamageMap{model.BY_ATK: 0.5} w/ 0.5 HitRatio means 25% of ATK.
	BaseDamage DamageMap

	// How much energy will be generated for the source from this attack. This energy generation
	// will scale with ERR.
	//
	// If HitRatio is specified, the energy gained will be multiplied by the hit ratio.
	// IE: 30.0 EnergyGain with a 0.5 HitRatio means only 15.0 energy added (before ERR bonus)
	EnergyGain float64

	// How much stance/toughness damage this attack will deal. This stance damage will scale with
	// Stance DMG% increase.
	//
	// If HitRatio is specified, the stance damage will be multiplied by the hit ratio.
	// IE: 30.0 StanceDamage with a 0.5 HitRatio means only 15 stance dmage will occur (before bonus)
	StanceDamage float64

	// Hit ratio reduces the BaseDamage, EnergyGain, and StanceDamage by the given percentage. This
	// is used for attacks that perform multiple hits. It is expected that the sum of all HitRatio
	// used for all hits in an attack equal 1.0 (IE: 2 attacks w/ HitRatio of 0.45 & 0.55)
	HitRatio float64

	// If true, will use the "pure damage" formula. This removes some variables from the damage
	// equation, such as crit.
	AsPureDamage bool

	// An additional flat damage amount that can be added to the base damage
	DamageValue float64

	// If set to true, will execute this hit in a "snapshot" state. This means that any modifiers
	// that subscribe to hit listeners will not be executed. This is used by break damage dots.
	UseSnapshot bool
}

type Heal struct {
	// The targets that the healer is healing
	Targets []key.TargetID

	// The healer that is performing the heal
	Source key.TargetID

	// Map of HealFormula -> Heal Percentage. This is for calculating the "Base Heal" amount of the
	// heal. IE: info.HealMap{model.BY_HEALER_MAX_HP: 0.5} = 50% of source target's MaxHP.
	BaseHeal HealMap

	// Additional flat healing that can be added to the base heal amount.
	HealValue float64

	// If set to true, will execute this heal in a "snapshot" state. This means that any modifiers
	// that subscribe to heal listeners will not be executed. This is used by phase1 heals.
	UseSnapshot bool
}
