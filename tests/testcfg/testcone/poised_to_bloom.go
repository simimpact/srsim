package testcone

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func PoisedToBloom() *model.LightCone {
	return &model.LightCone{
		Key:        key.PoisedToBloom.String(),
		Level:      80,
		MaxLevel:   80,
		Imposition: 1,
	}
}
