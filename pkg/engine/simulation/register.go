package simulation

import (
	"sync"

	"github.com/simimpact/srsim/pkg/engine/system"
	"github.com/simimpact/srsim/pkg/key"
)


var (
	mu        sync.RWMutex
	charMap   = make(map[key.CharacterKey]NewCharacterFunc)
	enemyMap   = make(map[key.EnemyKey]NewEnemyFunc)
)

type NewCharacterFunc func(sys *system.CharacterServices, id key.TargetID) (Target, error)

func RegisterCharFunc(char key.CharacterKey, f NewCharacterFunc) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := charMap[char]; dup {
		panic("combat: RegisterChar called twice for character " + char)
	}
	charMap[char] = f
}

type NewEnemyFunc func(sys *system.EnemyServices, id key.TargetID) (Target, error)

func RegisterEnemyFunc(enemy key.EnemyKey, f NewEnemyFunc) {
	mu.Lock()
	defer mu.Unlock()
	if _, dup := enemyMap[enemy]; dup {
		panic("combat: RegisterEnemy called twice for character " + enemy)
	}
	enemyMap[enemy] = f
}

