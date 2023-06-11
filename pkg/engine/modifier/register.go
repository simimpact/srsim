package modifier

import (
	"sync"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

var (
	mu              sync.Mutex
	modifierCatalog = make(map[key.Modifier]Config)
)

// Registers a new ModifierConfig to the modifier catalog for use in sim
func Register(key key.Modifier, modifier Config) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := modifierCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}
	modifierCatalog[key] = modifier
}

type Config struct {
	// Determines how duplicate modifiers will "stack" when added. Note: not all stacking behaviors can
	// stack modifiers.
	//
	// If unspecified, will default to the UNIQUE StackingBehavior
	Stacking StackingBehavior

	// Listeners to define any custom modifier logic. This allows you to hook onto sim behavior to add
	// conditional logic (IE: modify stats OnBeforeHit to increase damage)
	Listeners Listeners

	// If specified, modifier will only be applied for the given duration number of turns.
	// If unspecified, will default to -1 (can be overwritten by AddModifier)
	//
	// When Duration reaches 0, this modifier instance will be removed. A negative duration means
	// that this modifier will never expire
	Duration int

	// If specified, modifier will have the count number of stacks. In the event that an instance of
	// this modifier already exists, specifying count may replace, add to, or keep the existing stack
	// count depending on what the modifier "stacking" logic is.
	// If unspecified, will default to 1 (can be overwritten by AddModifier)
	//
	// When count reaches 0, the modifier will be removed from the target.
	Count float64

	// If specified, will set the max count allowed on this modifier in the event that it gets
	// reapplied/stacks and this instance is used.
	// If unspecified, will default to 1 (can be overwritten by AddModifier)
	MaxCount float64

	// When Count is unspecified, CountAddWhenStack determines how much to add to count when a new
	// stack is added.
	//
	// If unspecified, will default to 0 (can be overwritten by AddModifier)
	CountAddWhenStack float64

	// When this modifier "ticks" (phase1 or phase2). If unspecified, will default to phase2
	TickMoment TickMoment

	// Any BehaviorFlags that are associated with this modifier. These behavior flags trigger custom
	// behavior depending on the flag specified. IE: STAT_ flags will be used on "chance" modifier
	// applications to determine if the defender has any DebuffRES to use against this modifier.
	BehaviorFlags []model.BehaviorFlag

	// The type of status this modifier is (BUFF, DEBUFF, or OTHER). If unspecified, will default to
	// OTHER
	StatusType model.StatusType

	// Attacks and Heals can execute in a "snapshot" state. In this state, the modifier listeners will
	// not be called by default. Can be overwritten by setting this field to true
	CanModifySnapshot bool
}

// Determines how duplicate modifiers will "stack" when added. Note: not all stacking behaviors can
// stack modifiers.
type StackingBehavior int

const (
	// In the event of duplicates (checked by modifier name), will keep the current instance on
	// the target unmodified. Can never stack modifiers with this behavior.
	Unique StackingBehavior = iota
	// Compares modifiers by name and source (can have multiple instances of same mod). In the
	// event of an existing modifier from same source, will replace that instance with the incoming
	// instance. Can stack (count increases by CountAddWhenStack).
	ReplaceBySource
	// Compares modifiers by name. In the event of an existing modifier, will replace that instance
	// with the new instance. Can stack (count increases by CountAddWhenStack).
	Replace
	// Does no modifier comparisons/duplicate checks. Calling AddModifier will always add a new
	// modifier instance to the target.
	Multiple
	// In the event of duplicates (by name), will reset the Duration of the existing instance, but
	// not replace. This will not invoke OnAdd and instead will be OnDurationExtended
	Refresh
	// In the event of duplicates (by name), will add the incoming duration to the existing instance.
	// The original instance will be kept and OnDurationExtended will be called instead of OnAdd.
	Prolong
	// In the event of duplicates (by name), will keep the original instance and instead add the count
	// and duration. Will reset the current instance's stats and call OnAdd.
	Merge
)

type TickMoment int

const (
	ModifierPhase2End TickMoment = iota
	ModifierPhase1End
)

func (c Config) HasFlag(flag model.BehaviorFlag) bool {
	for _, f := range c.BehaviorFlags {
		if f == flag {
			return true
		}
	}
	return false
}
