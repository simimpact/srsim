package testchar

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func DanHung() *model.Character {
	return &model.Character{
		Key:      key.DanHeng.String(),
		Level:    80,
		MaxLevel: 80,
		Eidols:   0,
		Traces:   nil,
		Talents:  []uint32{1, 1, 1, 1},
		Cone: &model.LightCone{
			Key:        key.OnlySilenceRemains.String(),
			Level:      80,
			MaxLevel:   80,
			Imposition: 1,
		},
		Relics:      nil,
		StartEnergy: 0,
	}
}
