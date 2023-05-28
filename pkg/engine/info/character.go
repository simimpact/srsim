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
	Path          model.Path
	Element       model.DamageType
	BaseStats     PropMap
	BaseDebuffRES DebuffRESMap
	LightCone     LightCone
	// TODO: lighcone + relics?
}
