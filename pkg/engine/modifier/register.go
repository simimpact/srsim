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

	// configure the default values
	if modifier.Duration <= 0 {
		modifier.Duration = -1
	}
	if modifier.Count <= 0 {
		modifier.Count = 1
	}
	if modifier.MaxCount <= 0 {
		modifier.MaxCount = 1
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
	// TODO: WorkingTurn
}

type StackingBehavior int

// TODO: right default?
const (
	ReplaceBySource StackingBehavior = iota
	Replace
	Multiple
	Refresh
	Prolong
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
