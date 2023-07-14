package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type ShieldMap map[model.ShieldFormula]float64

type Shield struct {
	// The shielder that is creating the shield
	Source key.TargetID `json:"source"`

	// The target that the shielder is shielding
	Target key.TargetID `json:"target"`

	// Map of ShieldFormula -> Shield Percentage. This is for calculating the "base shield" amount of
	// the shield. IE: info.ShieldMap{model.ShieldFormula_SHIELD_BY_SHIELDER_ATK: 0.5} = 50% of
	// source target's ATK.
	BaseShield ShieldMap `json:"base_shield"`

	// Additional flat shield hp that can be added to the base heal amount.
	ShieldValue float64 `exhaustruct:"optional" json:"shield_value"`
}
