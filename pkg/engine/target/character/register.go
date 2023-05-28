package character

import (
	"sync"

	"github.com/simimpact/srsim/pkg/key"
)

var (
	mu               sync.Mutex
	characterCatalog = make(map[key.Character]Config)
)

func Register(key key.Character, character Config) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := characterCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}
	characterCatalog[key] = character
}
