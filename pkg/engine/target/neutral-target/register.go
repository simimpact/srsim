package neutraltarget

import (
	"sync"

	"github.com/simimpact/srsim/pkg/key"
)

var (
	mu             sync.Mutex
	neutralCatalog = make(map[key.NeutralTarget]Config)
)

func Register(key key.NeutralTarget, neutral Config) {
	mu.Lock()
	defer mu.Unlock()

	if _, dup := neutralCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}

	if neutral.Create == nil {
		panic("Neutral create function must be defined: " + key)
	}

	neutralCatalog[key] = neutral
}
