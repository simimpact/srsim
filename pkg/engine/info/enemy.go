package info

import "github.com/simimpact/srsim/pkg/key"

type Enemy struct {
	Key       key.Enemy
	Level     int
	Weakness  WeaknessMap
	DebuffRES DebuffRESMap
}
