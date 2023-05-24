package info

import "github.com/simimpact/srsim/pkg/model"

type PropMap map[model.Property]float64
type DebuffRESMap map[model.BehaviorFlag]float64

func NewPropMap() PropMap {
	return make(map[model.Property]float64)
}

func NewDebuffRESMap() DebuffRESMap {
	return make(map[model.BehaviorFlag]float64)
}

// adds a property to the PropMap using the correct equation (additive, multiplicative, or special)
func (m PropMap) Modify(prop model.Property, amt float64) {
	if prop == model.Property_ALL_DMG_REDUCE || prop == model.Property_FATIGUE {
		m[prop] = 1 - (1-m[prop])*(1-amt)
	} else {
		m[prop] += amt
	}
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
		m.Modify(k, v)
	}
}

// Add all debuff res values from the other DebuffRESMap to this DebuffRESMap
func (m DebuffRESMap) AddAll(other DebuffRESMap) {
	for k, v := range other {
		m.Modify(k, v)
	}
}
