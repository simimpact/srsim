package main

import "github.com/simimpact/srsim/pkg/model"

//nolint:tagliatelle // need to match datamine
type HashInfo struct {
	Hash int `json:"Hash"`
}

//nolint:tagliatelle // need to match datamine
type ValueInfo struct {
	Value float64 `json:"Value"`
}

//nolint:tagliatelle // need to match datamine
type EquipmentConfig struct {
	EquipmentID    int      `json:"EquipmentID"`
	EquipmentName  HashInfo `json:"EquipmentName"`
	Rarity         string   `json:"Rarity"`
	AvatarBaseType string   `json:"AvatarBaseType"`
	// ...
}

//nolint:tagliatelle // need to match datamine
type PromotionDataConfig struct {
	EquipmentID    int       `json:"EquipmentID"`
	Promotion      int       `json:"Promotion"`
	MaxLevel       int       `json:"MaxLevel"`
	BaseAttack     ValueInfo `json:"BaseAttack"`
	BaseAttackAdd  ValueInfo `json:"BaseAttackAdd"`
	BaseDefence    ValueInfo `json:"BaseDefence"`
	BaseDefenceAdd ValueInfo `json:"BaseDefenceAdd"`
	BaseHP         ValueInfo `json:"BaseHP"`
	BaseHPAdd      ValueInfo `json:"BaseHPAdd"`
	// ...
}

type PromotionConfig map[string]PromotionDataConfig

var pathTypes = map[string]model.Path{
	"Knight":  model.Path_PRESERVATION,
	"Rogue":   model.Path_HUNT,
	"Mage":    model.Path_ERUDITION,
	"Warlock": model.Path_NIHILITY,
	"Warrior": model.Path_DESTRUCTION,
	"Shaman":  model.Path_HARMONY,
	"Priest":  model.Path_ABUNDANCE,
}

func (e EquipmentConfig) GetPath() model.Path {
	t, ok := pathTypes[e.AvatarBaseType]
	if ok {
		return t
	}
	return model.Path_INVALID_PATH
}
