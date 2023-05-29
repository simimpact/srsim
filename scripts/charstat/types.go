package main

import "github.com/simimpact/srsim/pkg/model"

type SkillConfigType int

const (
	SkillTypeNone SkillConfigType = iota

	SkillTypeStatBonus
	SkillTypeAbility
	SkillTypeBonusAbility
)

type HashInfo struct {
	Hash int `json:"Hash"`
}

type ValueInfo struct {
	Value float64 `json:"Value"`
}

type StatusAdd struct {
	PropertyType string    `json:"PropertyType"`
	Value        ValueInfo `json:"Value"`
}

type AvatarInfo struct {
	AvatarName        HashInfo  `json:"AvatarName"`
	Rarity            string    `json:"Rarity"`
	UIAvatarModelPath string    `json:"UIAvatarModelPath"`
	DamageType        string    `json:"DamageType"`
	AvatarBaseType    string    `json:"AvatarBaseType"`
	SPNeed            ValueInfo `json:"SPNeed"`
	// ...
}

type SkillConfig struct {
	PointID              int             `json:"PointID"`
	PointType            SkillConfigType `json:"PointType"`
	AvatarID             int             `json:"AvatarID"`
	StatusAddList        []StatusAdd     `json:"StatusAddList"`
	AvatarLevelLimit     *int            `json:"AvatarLevelLimit"`
	AvatarPromotionLimit *int            `json:"AvatarPromotionLimit"`
	// ...
}

type PromotionDataConfig struct {
	AvatarID       int       `json:"AvatarID"`
	MaxLevel       int       `json:"MaxLevel"`
	AttackBase     ValueInfo `json:"AttackBase"`
	AttackAdd      ValueInfo `json:"AttackAdd"`
	DefenceBase    ValueInfo `json:"DefenceBase"`
	DefenceAdd     ValueInfo `json:"DefenceAdd"`
	HPBase         ValueInfo `json:"HPBase"`
	HPAdd          ValueInfo `json:"HPAdd"`
	SpeedBase      ValueInfo `json:"SpeedBase"`
	CriticalChance ValueInfo `json:"CriticalChance"`
	CriticalDamage ValueInfo `json:"CriticalDamage"`
	BaseAggro      ValueInfo `json:"BaseAggro"`
	// ...
}

type SkillTreeConfig map[string]SkillConfig
type PromotionConfig map[string]PromotionDataConfig

var damageTypes = map[string]model.DamageType{
	"Ice":       model.DamageType_ICE,
	"Wind":      model.DamageType_WIND,
	"Fire":      model.DamageType_FIRE,
	"Imaginary": model.DamageType_IMAGINARY,
	"Thunder":   model.DamageType_THUNDER,
	"Quantum":   model.DamageType_QUANTUM,
	"Physical":  model.DamageType_PHYSICAL,
}

func (a AvatarInfo) GetDamageType() model.DamageType {
	t, ok := damageTypes[a.DamageType]
	if ok {
		return t
	}
	return model.DamageType_INVALID_DAMAGE_TYPE
}

var pathTypes = map[string]model.Path{
	"Knight":  model.Path_PRESERVATION,
	"Rogue":   model.Path_HUNT,
	"Mage":    model.Path_ERUDITION,
	"Warlock": model.Path_NIHILITY,
	"Warrior": model.Path_DESTRUCTION,
	"Shaman":  model.Path_HARMONY,
	"Priest":  model.Path_ABUNDANCE,
}

func (a AvatarInfo) GetPath() model.Path {
	t, ok := pathTypes[a.AvatarBaseType]
	if ok {
		return t
	}
	return model.Path_INVALID_PATH
}

var statAdd = map[string]model.Property{
	"HPAddedRatio":          model.Property_HP_PERCENT,
	"AttackAddedRatio":      model.Property_ATK_PERCENT,
	"DefenceAddedRatio":     model.Property_DEF_PERCENT,
	"BaseSpeed":             model.Property_SPD_BASE,
	"CriticalDamageBase":    model.Property_CRIT_DMG,
	"CriticalChanceBase":    model.Property_CRIT_CHANCE,
	"SPRatioBase":           model.Property_ENERGY_REGEN,
	"StatusResistanceBase":  model.Property_EFFECT_HIT_RATE,
	"StatusProbabilityBase": model.Property_EFFECT_RES,
	"HealRatioBase":         model.Property_HEAL_BOOST,
	"HealTakenRatio":        model.Property_HEAL_TAKEN,
	"FireAddedRatio":        model.Property_FIRE_DMG_PERCENT,
	"IceAddedRatio":         model.Property_ICE_DMG_PERCENT,
	"ThunderAddedRatio":     model.Property_THUNDER_DMG_PERCENT,
	"QuantumAddedRatio":     model.Property_QUANTUM_DMG_PERCENT,
	"ImaginaryAddedRatio":   model.Property_IMAGINARY_DMG_PERCENT,
	"WindAddedRatio":        model.Property_WIND_DMG_PERCENT,
	"PhysicalAddedRatio":    model.Property_PHYSICAL_DMG_PERCENT,
}

func (s StatusAdd) GetType() model.Property {
	m, ok := statAdd[s.PropertyType]
	if ok {
		return m
	}
	return model.Property_INVALID_PROP
}
