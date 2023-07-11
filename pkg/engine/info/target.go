package info

import "github.com/simimpact/srsim/pkg/key"

type TargetClass string

type Retarget struct {
	// List of all targets to be filtered from.
	Targets []key.TargetID

	// Filter function. Used to filter Targets based on some condition(s).
	Filter func(target key.TargetID) bool

	// Maximum amount of targets to select from the filtered target list.
	// Defaults to the size of filtered list
	Max int `exhaustruct:"optional"`

	// Option to include targets in limbo(0 HP).
	IncludeLimbo bool `exhaustruct:"optional"`

	// Option to disable default random retarget behavior.
	DisableRandom bool `exhaustruct:"optional"`
}

const (
	ClassInvalid   TargetClass = "INVALID"
	ClassCharacter TargetClass = "CHARACTER"
	ClassEnemy     TargetClass = "ENEMY"
	ClassNeutral   TargetClass = "NEUTRAL"
)
