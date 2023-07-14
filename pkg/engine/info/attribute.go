package info

import (
	"github.com/simimpact/srsim/pkg/model"
)

type Attributes struct {
	Level         int          `json:"level"`
	BaseStats     PropMap      `json:"base_stats"`
	BaseDebuffRES DebuffRESMap `json:"base_debuff_res"`
	Weakness      WeaknessMap  `json:"weakness"`
	HPRatio       float64      `json:"hp_ratio"`
	Energy        float64      `json:"energy"`
	MaxEnergy     float64      `json:"max_energy"`
	Stance        float64      `json:"stance"`
	MaxStance     float64      `json:"max_stance"`
}

type ModifyHPByRatio struct {
	// The amount of HP ratio to modify the HP by (negative will remove HP)
	Ratio float64
	// What ratio type should be used (should Ratio be based on MaxHP or CurrentHP)
	RatioType model.ModifyHPRatioType
	// The floor for how low HP can go with this modification. IE: Floor = 1 will prevent the HP
	// from reaching 0 in this modification (can reduce up to 1 HP)
	Floor float64
}

func DefaultAttribute() Attributes {
	return Attributes{
		Level:         1,
		BaseStats:     NewPropMap(),
		BaseDebuffRES: NewDebuffRESMap(),
		Weakness:      NewWeaknessMap(),
		HPRatio:       1.0,
		Energy:        0,
		MaxEnergy:     0,
		Stance:        0,
		MaxStance:     0,
	}
}

type TargetState int

const (
	Invalid TargetState = 0
	Dead    TargetState = 1
	Limbo   TargetState = 2
	Alive   TargetState = 3
)
