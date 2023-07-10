package gen

import (
	_ "embed"
	"io"
	"strconv"
	"text/template"

	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/model"
)

//go:embed char_promo.tmpl
var tmplCharStr string
var tmplCharPromo *template.Template

var textMap map[string]string

func init() {
	var err error

	tmplCharPromo, err = template.New("tmplCharPromo").Parse(tmplCharStr)
	if err != nil {
		panic(err)
	}

	if IsDMAvailable() {
		if err := ReadDMFile(&textMap, "TextMap", "TextMapEN.json"); err != nil {
			panic(err)
		}
	}
}

type CharData struct {
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

func GenerateCharPromotions(w io.Writer, data *CharData) error {
	return tmplCharPromo.Execute(w, data)
}

func FromTextMap(hash int) string {
	return textMap[strconv.Itoa(hash)]
}
