package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Definition of a modifier. Is used to specify and additional stats or behavior associated with
// the target it is attached to (managed by the modifier package)
type Modifier struct {
	// What modifier this instance is (used to get metadata + listeners for the modifier)
	Name key.Modifier
	// TargetID for who created this modifier instance
	Source key.TargetID
	// Custom state that can be used to parameterize modifier logic (listeners can depend on state)
	// Note: State will be JSON serialized for logging purposes, so should be serialization friendly.
	State any
	// If specified, modifier will be applied with a random chance against the resistance
	// 		add prob = 1 - chance * (1 + source ERR) * (1 - target EffectRES) * (1 - target Debuff RES)
	// If unspecified, modifier will always be added (using the defined stacking logic)
	Chance float64
	// If specified, modifier will only be applied for the given duration number of turns
	// If unspecified, will default to the Duration in the ModifierConfig.
	//
	// When Duration reaches 0, this modifier instance will be removed. A negative duration means
	// that this modifier will never expire
	Duration int
	// When duration is > 0, the turn a modifier is added on will not count torwards the duration.
	// If this field is set to true, this will override that behavior and count the application turn
	// against the duration (if application happens before the check).
	TickImmediately bool
	// If specified, modifier will have the count number of stacks. In the event that an instance of
	// this modifier already exists, specifying count may replace, add to, or keep the existing stack
	// count depending on what the modifier "stacking" logic is.
	// If unspecified, will default to the Count in ModifierConfig (which will be 1 if not defined)
	//
	// When count reaches 0, the modifier will be removed from the target.
	Count float64
	// If specified, will set the max count allowed on this modifier in the event that it gets
	// reapplied/stacks and this instance is used.
	// If unspecified, will default to the MaxCount in ModifierConfig (which will be 1 if not defined)
	MaxCount float64
	// When Count is unspecified, CountAddWhenStack determines how much to add to count when a new
	// stack is added. Specifying this field will overrride the default value for this defined in the
	// ModifierConfig (which defaults to 0 if undefined)
	CountAddWhenStack float64

	Stats     PropMap
	DebuffRES DebuffRESMap
}

type ModifierState struct {
	Props     PropMap
	DebuffRES DebuffRESMap
	Flags     []model.BehaviorFlag
	Counts    map[model.StatusType]int
	Modifiers []key.Modifier
}

// TODO: ToProto for logging
