package character

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type register interface {
	AddTarget(key key.TargetID, base attribute.BaseStats) error
}

type Manager struct {
	engine    engine.Engine
	attr      register
	instances map[key.TargetID]CharInstance
	info      map[key.TargetID]info.Character
}

func New(engine engine.Engine, attr register) *Manager {
	return &Manager{
		engine:    engine,
		attr:      attr,
		instances: make(map[key.TargetID]CharInstance),
		info:      make(map[key.TargetID]info.Character),
	}
}

func (mgr *Manager) AddCharacter(id key.TargetID, char *model.Character) error {
	config, ok := characterCatalog[key.Character(char.Key)]
	if !ok {
		return fmt.Errorf("invalid character: %v", char.Key)
	}

	baseStats, maxLevel := config.fromPromotionData(int(char.Ascension), int(char.Level))
	err := mgr.attr.AddTarget(id, attribute.BaseStats{
		Stats:     baseStats,
		MaxEnergy: config.MaxEnergy,
	})
	if err != nil {
		return err
	}

	// TODO: lightcone + relic initialization (before or after character init?)

	info := info.Character{
		Key:       key.Character(char.Key),
		Level:     int(char.Level),
		MaxLevel:  maxLevel,
		Ascension: int(char.Ascension),
		Eidolon:   int(char.Eidols),
		Path:      config.Path,
		Element:   config.Element,
		BaseStats: baseStats,
		Traces:    char.Traces,
	}

	mgr.info[id] = info
	mgr.instances[id] = config.Create(mgr.engine, id, info)

	// TODO: emit CharacterAddedEvent
	return nil
}

func (mgr *Manager) Get(id key.TargetID) CharInstance {
	return mgr.instances[id]
}

func (mgr *Manager) Info(id key.TargetID) (info.Character, error) {
	if char, ok := mgr.info[id]; ok {
		return char, nil
	}
	return info.Character{}, fmt.Errorf("target is not a character: %v", id)
}
