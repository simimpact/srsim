package gen

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/model"
)

var keyRegex = regexp.MustCompile(`\W+`) // for removing spaces
var rarityRegex = regexp.MustCompile(`CombatPowerAvatarRarityType(\d+)`)

type SkillConfigType int

const (
	SkillTypeNone SkillConfigType = iota

	SkillTypeStatBonus
	SkillTypeAbility
	SkillTypeBonusAbility
)

//nolint:tagliatelle // need to match datamine
type AvatarInfo struct {
	AvatarName        HashInfo  `json:"AvatarName"`
	Rarity            string    `json:"Rarity"`
	UIAvatarModelPath string    `json:"UIAvatarModelPath"`
	DamageType        string    `json:"DamageType"`
	AvatarBaseType    string    `json:"AvatarBaseType"`
	SPNeed            ValueInfo `json:"SPNeed"`
	JSONPath          string    `json:"JsonPath"`
	// ...
}

//nolint:tagliatelle // need to match datamine
type AvatarConfig struct {
	SkillList []AvatarSkillMetadata `json:"SkillList"`
	// ...
}

//nolint:tagliatelle // need to match datamine
type AvatarSkillMetadata struct {
	Name       string     `json:"Name"`
	SkillType  string     `json:"SkillType"`
	UseType    string     `json:"UseType"`
	TargetInfo TargetInfo `json:"TargetInfo"`
	// ...
}

//nolint:tagliatelle // need to match datamine
type TraceConfig struct {
	PointID              int             `json:"PointID"`
	PointType            SkillConfigType `json:"PointType"`
	AvatarID             int             `json:"AvatarID"`
	StatusAddList        []StatusAdd     `json:"StatusAddList"`
	AvatarLevelLimit     *int            `json:"AvatarLevelLimit"`
	AvatarPromotionLimit *int            `json:"AvatarPromotionLimit"`
	// ...
}

//nolint:tagliatelle // need to match datamine
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

//nolint:tagliatelle // need to match datamine
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

func GetCharacters() map[string]*CharData {
	if !IsDMAvailable() {
		return nil
	}

	var avatars map[string]AvatarInfo
	var skills map[string]SkillTreeConfig
	var promotions map[string]PromotionConfig
	var avatarSkills map[string]SkillConfig

	ReadDMFile(&avatars, "ExcelOutput", "AvatarConfig.json")
	ReadDMFile(&skills, "ExcelOutput", "AvatarSkillTreeConfig.json")
	ReadDMFile(&promotions, "ExcelOutput", "AvatarPromotionConfig.json")
	ReadDMFile(&avatarSkills, "ExcelOutput", "AvatarSkillConfig.json")

	out := make(map[string]*CharData)
	for id, value := range avatars {
		charName := FromTextMap(value.AvatarName.Hash)
		switch charName {
		case "":
			continue
		case "{NICKNAME}":
			charName = "Trailblazer" + value.DamageType
		}

		var config AvatarConfig
		ReadDMFile(&config, value.JSONPath)

		key := keyRegex.ReplaceAllString(charName, "")
		data := &CharData{
			Key:           key,
			KeyLower:      strings.ToLower(key),
			Rarity:        rarityRegex.FindStringSubmatch(value.Rarity)[1],
			Element:       value.GetDamageType(),
			Path:          value.GetPath(),
			MaxEnergy:     int(value.SPNeed.Value),
			PromotionData: promotions[id].ToData(),
			Traces:        ToTraceMap(FindCharSkills(skills, id)),
			SkillInfo:     FindSkillInfo(avatarSkills, config, key),
		}
		out[data.KeyLower] = data
	}
	return out
}

func FindCharSkills(skills map[string]SkillTreeConfig, key string) []SkillTreeConfig {
	id, _ := strconv.Atoi(key)
	result := make([]SkillTreeConfig, 0)
	for _, value := range skills {
		if value["1"].AvatarID == id {
			result = append(result, value)
		}
	}
	return result
}

func ToTraceMap(skills []SkillTreeConfig) character.TraceMap {
	out := make(character.TraceMap, len(skills))
	for _, config := range skills {
		value := config["1"]
		switch value.PointType {
		case SkillTypeStatBonus:
		case SkillTypeBonusAbility:
		default:
			continue
		}

		var trace character.Trace
		if len(value.StatusAddList) > 0 {
			trace.Stat = value.StatusAddList[0].GetType()
			trace.Amount = value.StatusAddList[0].Value.Value
		}
		if value.AvatarLevelLimit != nil {
			trace.Level = *value.AvatarLevelLimit
		}
		if value.AvatarPromotionLimit != nil {
			trace.Ascension = *value.AvatarPromotionLimit
		}
		key := strconv.Itoa(value.PointID)
		out[key[len(key)-3:]] = trace
	}
	return out
}

func (p PromotionConfig) ToData() []character.PromotionData {
	out := make([]character.PromotionData, len(p))
	for i := 0; i < len(p); i++ {
		val, ok := p[strconv.Itoa(i)]
		if !ok {
			break
		}
		out[i] = character.PromotionData{
			MaxLevel:   val.MaxLevel,
			ATKBase:    val.AttackBase.Value,
			ATKAdd:     val.AttackAdd.Value,
			DEFBase:    val.DefenceBase.Value,
			DEFAdd:     val.DefenceAdd.Value,
			HPBase:     val.HPBase.Value,
			HPAdd:      val.HPAdd.Value,
			SPD:        val.SpeedBase.Value,
			CritChance: val.CriticalChance.Value,
			CritDMG:    val.CriticalDamage.Value,
			Aggro:      val.BaseAggro.Value,
		}
	}
	return out
}

func FindSkillInfo(skills map[string]SkillConfig, config AvatarConfig, key string) character.SkillInfo {
	var info character.SkillInfo
	for k, value := range skills {
		if !strings.HasPrefix(k, key) {
			continue
		}

		var targetType model.TargetType
		for _, s := range config.SkillList {
			if s.Name == value["1"].SkillTriggerKey {
				targetType = s.TargetInfo.GetType()
				break
			}
		}

		bpAdd := int(value["1"].BPAdd.Value)
		if bpAdd < 0 {
			bpAdd = 0
		}

		bpNeed := int(value["1"].BPNeed.Value)
		if bpNeed < 0 {
			bpNeed = 0
		}

		switch value["1"].SkillTriggerKey {
		case "Skill01":
			info.Attack = character.Attack{
				SPAdd:      bpAdd,
				TargetType: targetType,
			}
		case "Skill02":
			info.Skill = character.Skill{
				SPNeed:     bpNeed,
				TargetType: targetType,
			}
		case "Skill03":
			info.Ult = character.Ult{
				TargetType: targetType,
			}
		case "SkillMaze":
			info.Technique = character.Technique{
				TargetType: targetType,
				IsAttack:   value["1"].SkillEffect == "MazeAttack",
			}
		default:
			continue
		}
	}
	return info
}

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
