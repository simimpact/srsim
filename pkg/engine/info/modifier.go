package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// An instance of a modifier. Is used to specify and additional stats or behavior associated with
// the target it is attached to (managed by the modifier package)
type ModifierInstance struct {
	// What modifier this instance is (used to get metadata + listeners for the modifier)
	Name key.Modifier
	// TargetID for who created this modifier instance
	Source key.TargetID
	// Any custom params to be defined that are used by the underlying modifier logic
	Params map[string]float64
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
	Count int
	// If specified, will set the max count allowed on this modifier in the event that it gets
	// reapplied/stacks and this instance is used.
	// If unspecified, will default to the MaxCount in ModifierConfig (which will be 1 if not defined)
	MaxCount int
	// When Count is unspecified, CountAddWhenStack determines how much to add to count when a new
	// stack is added. Specifying this field will overrride the default value for this defined in the
	// ModifierConfig (which defaults to 0 if undefined)
	CountAddWhenStack int

	// internal fields
	stats     PropMap
	debuffRES DebuffRESMap
	renewTurn int
}

type ModifierState struct {
	Props     PropMap
	DebuffRES DebuffRESMap
	Flags     []model.BehaviorFlag
	Counts    map[model.StatusType]int
	Modifiers []key.Modifier
}

// Resets the modifier instance (will remove any stats, debuffRES, and set a new renew turn)
func (mi *ModifierInstance) Reset(turnCount int) {
	mi.stats = NewPropMap()
	mi.debuffRES = make(map[model.BehaviorFlag]float64)
	mi.renewTurn = turnCount
}

// Returns the renew turn for this modifier instance
func (mi *ModifierInstance) RenewTurn() int {
	return mi.renewTurn
}

// Get the current value of a property set by this modifier instance
func (mi *ModifierInstance) GetProperty(prop model.Property) float64 {
	return mi.stats[prop]
}

// Get the current value of a debuff res set by this modifier instance
func (mi *ModifierInstance) GetDebuffRES(flags ...model.BehaviorFlag) float64 {
	return mi.debuffRES.GetDebuffRES(flags...)
}

// Add a property to this modifier instance
func (mi *ModifierInstance) AddProperty(prop model.Property, amt float64) {
	mi.stats.Modify(prop, amt)
}

// Add a new debuffRES for the given behavior flag
func (mi *ModifierInstance) AddDebuffRES(flag model.BehaviorFlag, amt float64) {
	mi.debuffRES.Modify(flag, amt)
}

// TODO: ToProto for logging
