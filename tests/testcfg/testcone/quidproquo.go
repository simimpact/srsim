package testcone

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func QuidProQuo() *model.LightCone {
	return &model.LightCone{
		Key:        key.QuidProQuo.String(),
		Level:      80,
		MaxLevel:   80,
		Imposition: 1,
	}
}
