package hook

import (
	"sync"

	"github.com/simimpact/srsim/pkg/engine"
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

func StartupHooks() map[string]HookFunc {
	return hooks
}
