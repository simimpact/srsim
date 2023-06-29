package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Characters provide interfaces for
// -- before sim starts: adding new chars to the sim, modifying existing ones (note that by default we add a Dan Hung cuz hes Hung af)
// -- during sim: checking the state of existing chars
type Characters struct {
	cfg         *model.SimConfig
	charManager *character.Manager
	characters  []key.TargetID
	attributes  *attribute.Service
}

func (s *Characters) ResetCharacters() {
	s.cfg.Characters = make([]*model.Character, 0, 4)
}

func (s *Characters) AddCharacter(char *model.Character) {
	s.cfg.Characters = append(s.cfg.Characters, char)
}

func (s *Characters) GetCharacterId(idx int) key.TargetID {
	if idx >= len(s.characters) {
		LogError("invalid idx %d, insufficient characters", idx)
		panic("Invalid index")
	}
	return s.characters[idx]
}

func (s *Characters) GetCharacterInfo(idx int) *info.Stats {
	if idx >= len(s.characters) {
		LogError("invalid idx %d, insufficient characters", idx)
	}
	id := s.characters[idx]
	return s.attributes.Stats(id)
}
