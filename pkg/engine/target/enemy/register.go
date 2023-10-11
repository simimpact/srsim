package enemy

import (
	"sync"

	"github.com/simimpact/srsim/pkg/key"
)

var (
	mu           sync.Mutex
	enemyCatalog = make(map[key.Enemy]Config)
)

func Register(key key.Enemy, enemy Config) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := enemyCatalog[key]; dup {
		panic("duplicate registration attempt: " + key)
	}

	if enemy.Create == nil {
		panic("Enemy Create must be defined: " + key)
	}

	enemyCatalog[key] = enemy
}
