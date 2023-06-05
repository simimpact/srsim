package info

import "github.com/simimpact/srsim/pkg/key"

type Enemy struct {
	Key       key.Enemy
	Level     int
	MaxStance float64
	Weakness  WeaknessMap
	DebuffRES DebuffRESMap
}
