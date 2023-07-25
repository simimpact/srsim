package testchar

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/tests/testcfg/testcone"
)

func DummyChar() *model.Character {
	return &model.Character{
		Key:         key.DummyCharacter.String(),
		Level:       80,
		MaxLevel:    80,
		Eidols:      0,
		Traces:      nil,
		Abilities:   nil,
		LightCone:   testcone.DataBank(),
		Relics:      nil,
		StartEnergy: 0,
	}
}
