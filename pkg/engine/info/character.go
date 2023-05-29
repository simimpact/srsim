package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Character struct {
	Key          key.Character
	Level        int
	Ascension    int
	Eidolon      int
	Traces       map[string]bool
	AbilityLevel AbilityLevel
	Path         model.Path
	Element      model.DamageType
	LightCone    LightCone
	Relics       map[key.Relic]int
}

type AbilityLevel struct {
	Attack int
	Skill  int
	Ult    int
	Talent int
}

type LightCone struct {
	Key       key.LightCone
	Level     int
	Ascension int
	Rank      int
	Path      model.Path
}
