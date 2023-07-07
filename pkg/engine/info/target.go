package info

import "github.com/simimpact/srsim/pkg/key"

type TargetClass int

type Retarget struct {
	// List of all targets to be filtered from
	Targets []key.TargetID

	// Filter function
	Filter func(targets []key.TargetID) []key.TargetID

	// Maximum amount of targets to select from the filtered target list
	Max int

	// Option to include targets in limbo(0 HP)
	IncludeLimbo bool
}

const (
	ClassInvalid TargetClass = iota
	ClassCharacter
	ClassEnemy
	ClassNeutral
)
