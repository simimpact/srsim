package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Character struct {
	Key       key.Character
	Level     int
	MaxLevel  int
	Ascension int
	Eidolon   int
	Traces    []string
	Path      model.Path
	Element   model.DamageType
	BaseStats PropMap
	// TODO: lighcone + relics?
}
