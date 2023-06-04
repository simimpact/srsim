package info

import (
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

type PropMap map[prop.Property]float64
type DebuffRESMap map[model.BehaviorFlag]float64
type WeaknessMap map[model.DamageType]bool

func NewPropMap() PropMap {
	return make(map[prop.Property]float64)
}

func NewDebuffRESMap() DebuffRESMap {
	return make(map[model.BehaviorFlag]float64)
}

func NewWeaknessMap() WeaknessMap {
	return make(map[model.DamageType]bool)
}

// adds a property to the PropMap using the correct equation (additive, multiplicative, or special)
func (m PropMap) Modify(p prop.Property, amt float64) {
	if p == prop.AllDamageReduce || p == prop.Fatigue {
		m[p] = 1 - (1-m[p])*(1-amt)
	} else {
		m[p] += amt
	}
}

// resets the given property to the given value (will overwrite the existing value)
func (m PropMap) Set(p prop.Property, amt float64) {
	m[p] = 0
	m.Modify(p, amt)
}

// Adds a debuff res to the DebuffRESMap for the given behavior flag
// TODO: unknown if this stacks additively or multiplicatively or only takes max
func (m DebuffRESMap) Modify(flag model.BehaviorFlag, amt float64) {
	m[flag] += amt
}

// Gets the current Debuff RES given the set of flags (max of res for associated flags)
func (m DebuffRESMap) GetDebuffRES(flags ...model.BehaviorFlag) float64 {
	out := 0.0
	for _, flag := range flags {
		res := m[flag]
		if res > out {
			out = res
		}
	}
	return out
}

// Add all properties from other PropMap to this PropMap
func (m PropMap) AddAll(other PropMap) {
	for k, v := range other {
		if v != 0 {
			m.Modify(k, v)
		}
	}
}

// Add all debuff res values from the other DebuffRESMap to this DebuffRESMap
func (m DebuffRESMap) AddAll(other DebuffRESMap) {
	for k, v := range other {
		if v != 0 {
			m.Modify(k, v)
		}
	}
}

// Add all weaknesses from the other WeaknessMap to this WeaknessMap
func (m WeaknessMap) AddAll(other WeaknessMap) {
	for k, v := range other {
		if v {
			m[k] = v
		}
	}
}
