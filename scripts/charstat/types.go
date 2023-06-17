package main

import (
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

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

type TargetInfo struct {
	TargetType string `json:"TargetType"`
}

type AvatarInfo struct {
	AvatarName        HashInfo  `json:"AvatarName"`
	Rarity            string    `json:"Rarity"`
	UIAvatarModelPath string    `json:"UIAvatarModelPath"`
	DamageType        string    `json:"DamageType"`
	AvatarBaseType    string    `json:"AvatarBaseType"`
	SPNeed            ValueInfo `json:"SPNeed"`
	JsonPath          string    `json:"JsonPath"`
	// ...
}

type AvatarConfig struct {
	SkillList []AvatarSkillMetadata `json:"SkillList"`
	// ...
}

type AvatarSkillMetadata struct {
	Name       string     `json:"Name"`
	SkillType  string     `json:"SkillType"`
	UseType    string     `json:"UseType"`
	TargetInfo TargetInfo `json:"TargetInfo"`
	// ...
}

type TraceConfig struct {
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

type AvatarSkillConfig struct {
	BPNeed          ValueInfo `json:"BPNeed"`
	BPAdd           ValueInfo `json:"BPAdd"`
	SkillEffect     string    `json:"SkillEffect"`
	AttackType      string    `json:"AttackType"`
	SkillTriggerKey string    `json:"SkillTriggerKey"`
	// ...
}

type SkillTreeConfig map[string]TraceConfig
type PromotionConfig map[string]PromotionDataConfig
type SkillConfig map[string]AvatarSkillConfig

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

func (s StatusAdd) GetType() prop.Property {
	m, ok := prop.GameToProp[s.PropertyType]
	if ok {
		return m
	}
	return prop.Invalid
}

var targetTypes = map[string]model.TargetType{
	"Caster":        model.TargetType_SELF,
	"FriendSelect":  model.TargetType_ALLIES,
	"EnemySelect":   model.TargetType_ENEMIES,
	"AllTeamMember": model.TargetType_ALLIES,
	"AllEnemy":      model.TargetType_ENEMIES,
}

func (t TargetInfo) GetType() model.TargetType {
	m, ok := targetTypes[t.TargetType]
	if ok {
		return m
	}
	return model.TargetType_INVALID_TARGET_TYPE
}
