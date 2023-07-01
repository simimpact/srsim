//go:generate go run generate.go types.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/model"
)

// data for templates
type dataTmpl struct {
	Key           string
	KeyLower      string
	Rarity        string
	Element       model.DamageType
	Path          model.Path
	MaxEnergy     int
	PromotionData []character.PromotionData
	Traces        character.TraceMap
	SkillInfo     character.SkillInfo
}

var keyRegex = regexp.MustCompile(`\W+`) // for removing spaces
var rarityRegex = regexp.MustCompile(`CombatPowerAvatarRarityType(\d+)`)

func main() {
	dmPath := os.Getenv("DM_PATH")
	if dmPath == "" {
		fmt.Println("Please provide the path to StarRailData (environment variable DM_PATH).")
		return
	}

	var avatars map[string]AvatarInfo
	var skills map[string]SkillTreeConfig
	var promotions map[string]PromotionConfig
	var textMap map[string]string
	var avatarSkills map[string]SkillConfig

	err := OpenConfig(&avatars, dmPath, "ExcelOutput", "AvatarConfig.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = OpenConfig(&skills, dmPath, "ExcelOutput", "AvatarSkillTreeConfig.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = OpenConfig(&promotions, dmPath, "ExcelOutput", "AvatarPromotionConfig.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = OpenConfig(&textMap, dmPath, "TextMap", "TextMapEN.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = OpenConfig(&avatarSkills, dmPath, "ExcelOutput", "AvatarSkillConfig.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, value := range avatars {
		id, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		charName := GetCharacterName(textMap, value.AvatarName.Hash)
		switch charName {
		case "":
			continue
		case "{NICKNAME}":
			charName = "Trailblazer" + value.DamageType
		}

		var avatarConfig AvatarConfig
		err = OpenConfig(&avatarConfig, dmPath, value.JSONPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		info := FindSkillInfo(avatarSkills, avatarConfig, key)
		ProcessCharacter(charName, value, FindCharSkills(skills, id), info, promotions[key])
	}
}

func OpenConfig(result interface{}, path ...string) error {
	jsonFile := filepath.Join(path...)
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	return nil
}

func FindCharSkills(skills map[string]SkillTreeConfig, id int) []SkillTreeConfig {
	result := make([]SkillTreeConfig, 0)
	for _, value := range skills {
		if value["1"].AvatarID == id {
			result = append(result, value)
		}
	}
	return result
}

func FindSkillInfo(skills map[string]SkillConfig, config AvatarConfig, key string) character.SkillInfo {
	info := character.SkillInfo{}
	for k, value := range skills {
		if strings.HasPrefix(k, key) {
			var targetType model.TargetType
			for _, s := range config.SkillList {
				if s.Name == value["1"].SkillTriggerKey {
					targetType = s.TargetInfo.GetType()
					break
				}
			}

			BPAdd := int(value["1"].BPAdd.Value)
			if BPAdd < 0 {
				BPAdd = 0
			}

			BPNeed := int(value["1"].BPNeed.Value)
			if BPNeed < 0 {
				BPNeed = 0
			}

			switch value["1"].SkillTriggerKey {
			case "Skill01":
				info.Attack = character.Attack{
					SPAdd:      BPAdd,
					TargetType: targetType,
				}
			case "Skill02":
				info.Skill = character.Skill{
					SPNeed:     BPNeed,
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
	}
	return info
}

func GetCharacterName(textMap map[string]string, hash int) string {
	return textMap[strconv.Itoa(hash)]
}

func ProcessCharacter(
	name string, avatar AvatarInfo, skills []SkillTreeConfig, info character.SkillInfo, promotions PromotionConfig) {
	data := dataTmpl{}
	data.Key = keyRegex.ReplaceAllString(name, "")
	data.KeyLower = strings.ToLower(data.Key)
	data.Rarity = rarityRegex.FindStringSubmatch(avatar.Rarity)[1]
	data.Element = avatar.GetDamageType()
	data.Path = avatar.GetPath()
	data.MaxEnergy = int(avatar.SPNeed.Value)
	data.SkillInfo = info

	data.PromotionData = make([]character.PromotionData, len(promotions))
	for i := 0; i < len(promotions); i++ {
		val, ok := promotions[strconv.Itoa(i)]
		if !ok {
			break
		}
		data.PromotionData[i] = character.PromotionData{
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

	data.Traces = make(character.TraceMap)
	for _, config := range skills {
		value := config["1"]
		switch value.PointType {
		case SkillTypeStatBonus:
		case SkillTypeBonusAbility:
		default:
			continue
		}

		trace := character.Trace{}
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
		data.Traces[strconv.Itoa(value.PointID)] = trace
	}

	// save .go files
	path := filepath.Join(".", "result", data.KeyLower)
	os.MkdirAll(path, os.ModePerm)

	fchar, err := os.Create(filepath.Join(path, data.KeyLower+".go"))
	if err != nil {
		log.Fatal(err)
	}
	defer fchar.Close()
	tchar, err := template.New("outchar").Parse(tmplChar)
	if err != nil {
		log.Fatal(err)
	}
	if err := tchar.Execute(fchar, data); err != nil {
		log.Fatal(err)
	}

	fdata, err := os.Create(filepath.Join(path, "data.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer fdata.Close()
	tdata, err := template.New("outdata").Parse(tmplData)
	if err != nil {
		log.Fatal(err)
	}
	if err := tdata.Execute(fdata, data); err != nil {
		log.Fatal(err)
	}

	fatk, err := os.Create(filepath.Join(path, "attack.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer fatk.Close()
	tatk, err := template.New("outattack").Parse(tmplAtk)
	if err != nil {
		log.Fatal(err)
	}
	if err := tatk.Execute(fatk, data); err != nil {
		log.Fatal(err)
	}

	fskill, err := os.Create(filepath.Join(path, "skill.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer fskill.Close()
	tskill, err := template.New("outskill").Parse(tmplSkill)
	if err != nil {
		log.Fatal(err)
	}
	if err := tskill.Execute(fskill, data); err != nil {
		log.Fatal(err)
	}

	fult, err := os.Create(filepath.Join(path, "ult.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer fult.Close()
	tult, err := template.New("outult").Parse(tmplUlt)
	if err != nil {
		log.Fatal(err)
	}
	if err := tult.Execute(fult, data); err != nil {
		log.Fatal(err)
	}

	ftech, err := os.Create(filepath.Join(path, "technique.go"))
	if err != nil {
		log.Fatal(err)
	}
	defer ftech.Close()
	ttech, err := template.New("outtech").Parse(tmplTechnique)
	if err != nil {
		log.Fatal(err)
	}
	if err := ttech.Execute(ftech, data); err != nil {
		log.Fatal(err)
	}
}

var tmplChar = `package {{.KeyLower}}

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	character.Register(key.{{.Key}}, character.Config{
		Create:     NewInstance,
		Rarity:     {{.Rarity}},
		Element:    model.DamageType_{{.Element}},
		Path:       model.Path_{{.Path}},
		MaxEnergy:  {{.MaxEnergy}},
		Promotions: promotions,
		Traces:     traces,
		SkillInfo: character.SkillInfo{
			Attack: character.Attack{
				SPAdd:      {{.SkillInfo.Attack.SPAdd}},
				TargetType: model.TargetType_{{.SkillInfo.Attack.TargetType}},
			},
			Skill: character.Skill{
				SPNeed:     {{.SkillInfo.Skill.SPNeed}},
				TargetType: model.TargetType_{{.SkillInfo.Skill.TargetType}},
			},
			Ult: character.Ult{
				TargetType: model.TargetType_{{.SkillInfo.Ult.TargetType}},
			},
			Technique: character.Technique{
				TargetType: model.TargetType_{{.SkillInfo.Technique.TargetType}},
				IsAttack:   {{.SkillInfo.Technique.IsAttack}},
			},
		},
	})
}

type char struct {
	engine engine.Engine
	id     key.TargetID
	info   info.Character
}

func NewInstance(engine engine.Engine, id key.TargetID, charInfo info.Character) info.CharInstance {
	c := &char{
		engine: engine,
		id:     id,
		info:   charInfo,
	}

	return c
}
`

var tmplData = `// Code generated by "charstat"; DO NOT EDIT.

package {{.KeyLower}}

import (
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/engine/target/character"
)

var promotions = []character.PromotionData{
	{{- range $e := $.PromotionData}}
	{
		MaxLevel:   {{$e.MaxLevel}},
		ATKBase:    {{$e.ATKBase}},
		ATKAdd:     {{$e.ATKAdd}},
		DEFBase:    {{$e.DEFBase}},
		DEFAdd:     {{$e.DEFAdd}},
		HPBase:     {{$e.HPBase}},
		HPAdd:      {{$e.HPAdd}},
		SPD:        {{$e.SPD}},
		CritChance: {{$e.CritChance}},
		CritDMG:    {{$e.CritDMG}},
		Aggro:      {{$e.Aggro}},
	},
	{{- end}}
}

var traces = character.TraceMap{
	{{- range $key, $value := .Traces}}
	"{{$key}}": {
		{{- if $value.Stat }}
		Stat: prop.{{$value.Stat}},
		{{- end}}
		{{- if $value.Amount }}
		Amount: {{$value.Amount}},
		{{- end}}
		{{- if $value.Ascension }}
		Ascension: {{$value.Ascension}},
		{{- end}}
		{{- if $value.Level }}
		Level: {{$value.Level}},
		{{- end}}
	},
	{{- end}}
}`

var tmplAtk = `package {{.KeyLower}}

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {

}
`

var tmplSkill = `package {{.KeyLower}}

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	
}
`

var tmplUlt = `package {{.KeyLower}}

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	
}
`

var tmplTechnique = `package {{.KeyLower}}

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	
}
`
