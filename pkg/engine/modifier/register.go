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
	Stacking          StackingBehavior
	Listeners         Listeners
	Duration          int
	Count             int
	MaxCount          int
	CountAddWhenStack int
	TickMoment        BattlePhase
	BehaviorFlags     []model.BehaviorFlag
	StatusType        model.StatusType
	// TODO: WorkingTurn?
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

type BattlePhase int

const (
	ModifierPhase2End BattlePhase = iota
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
