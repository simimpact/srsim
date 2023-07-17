package info

import (
	"github.com/simimpact/srsim/pkg/key"
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
	// A unique identifier for this modification
	Key key.Reason `json:"key"`

	// The target of this HP modification
	Target key.TargetID `json:"target"`

	// The source of this HP modification (who caused it)
	Source key.TargetID `json:"source"`

	// The amount of HP ratio to modify the HP by (negative will remove HP)
	Ratio float64 `json:"ratio"`

	// What ratio type should be used (should Ratio be based on MaxHP or CurrentHP)
	RatioType model.ModifyHPRatioType `json:"ratio_type"`

	// The floor for how low HP can go with this modification. IE: Floor = 1 will prevent the HP
	// from reaching 0 in this modification (can reduce up to 1 HP)
	Floor float64 `json:"floor"`
}

type ModifyAttribute struct {
	// A unique identifier for this modification
	Key key.Reason `json:"key"`

	// The target of this modification
	Target key.TargetID `json:"target"`

	// The source of this modification
	Source key.TargetID `json:"source"`

	// The amount that should be modified (added or removed)
	Amount float64 `json:"amount"`
}

type ModifySP struct {
	// A unique identifier for this modification
	Key key.Reason `json:"key"`

	// The source of this modification
	Source key.TargetID `json:"source"`

	// The amount of SP to be added or removed
	Amount int `json:"amount"`
}

type ModifyCurrentGaugeCost struct {
	// A unique identifier for this modification
	Key key.Reason `json:"key"`

	// The source of this modification
	Source key.TargetID `json:"source"`

	// The amount of gauge cost to be changed
	Amount float64 `json:"amount"`
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
