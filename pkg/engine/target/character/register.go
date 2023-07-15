package character

import (
	"sync"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
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

	if character.Create == nil {
		panic("Character Create must be defined: " + key)
	}
	if len(character.Promotions) == 0 {
		panic("Character promotions must be defined: " + key)
	}
	if character.Path == model.Path_INVALID_PATH {
		panic("Character path must be defined: " + key)
	}
	if character.Element == model.DamageType_INVALID_DAMAGE_TYPE {
		panic("character element/damage type must be defined: " + key)
	}
	if character.Traces == nil {
		character.Traces = make(TraceMap)
	}

	characterCatalog[key] = character
}

func Retrieve(key key.Character) Config {
	return characterCatalog[key]
}
