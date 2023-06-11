package simulation

import (
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

type HookFunc func(engine engine.Engine) error

var (
	mu    sync.Mutex
	hooks = make(map[string]HookFunc)
)

// startup hooks are for adding any custom logic that is not part of the core engine.
// IE: break effects, energy gen on death, etc
func RegisterStartupHook(key string, f HookFunc) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := hooks[key]; dup {
		panic("duplicate registration attempt: " + key)
	}
	hooks[key] = f
}

func (s *simulation) subscribe() {
	s.event.TargetDeath.Subscribe(s.onDeath)
}

func (s *simulation) onDeath(e event.TargetDeathEvent) {
	// remove this target from active arrays (these arrays represent order in battle map)
	switch s.targets[e.Target] {
	case info.ClassCharacter:
		s.characters = remove(s.characters, e.Target)
	case info.ClassEnemy:
		s.enemies = remove(s.enemies, e.Target)
	case info.ClassNeutral:
		s.neutrals = remove(s.neutrals, e.Target)
	}

	// remove this target from the turn order
	s.turnManager.RemoveTarget(e.Target)
}

func remove(arr []key.TargetID, id key.TargetID) []key.TargetID {
	for i, t := range arr {
		if id == t {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
