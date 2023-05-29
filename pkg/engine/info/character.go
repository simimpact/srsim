package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Character struct {
	Key           key.Character
	Level         int
	Ascension     int
	Eidolon       int
	Traces        map[string]bool
	AbilityLevels AbilityLevels
	Path          model.Path
	Element       model.DamageType
	LightCone     LightCone
	Relics        map[key.Relic]int
	BaseStats     PropMap
	BaseDebuffRES DebuffRESMap
}

type AbilityLevels struct {
	Attack int
	Skill  int
	Ult    int
	Talent int
}
