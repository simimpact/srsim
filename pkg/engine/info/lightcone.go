package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type LightCone struct {
	Key       key.LightCone
	Level     int
	Ascension int
	Rank      int
	Path      model.Path
}
