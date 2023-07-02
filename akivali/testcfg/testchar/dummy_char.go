package testchar

import (
	"github.com/simimpact/srsim/akivali/testcfg/testcone"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func DummyChar() *model.Character {
	return &model.Character{
		Key:         key.DummyCharacter.String(),
		Level:       80,
		MaxLevel:    80,
		Eidols:      0,
		Traces:      nil,
		Talents:     nil,
		Cone:        testcone.DataBank(),
		Relics:      nil,
		StartEnergy: 0,
	}
}
