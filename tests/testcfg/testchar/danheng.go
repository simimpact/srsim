package testchar

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/tests/testcfg/testcone"
)

func DanHung() *model.Character {
	return &model.Character{
		Key:      key.DanHeng.String(),
		Level:    80,
		MaxLevel: 80,
		Eidols:   0,
		Traces:   nil,
		Abilities: &model.Abilities{
			Attack: 1,
			Skill:  1,
			Ult:    1,
			Talent: 1,
		},
		LightCone:   testcone.OnlySilenceRemains(),
		Relics:      nil,
		StartEnergy: 0,
	}
}
