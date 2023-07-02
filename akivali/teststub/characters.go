package teststub

import (
	"github.com/simimpact/srsim/akivali/testcfg/testeval"
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Characters provide interfaces for
// -- before sim starts: adding new chars to the sim, modifying existing ones
// -- during sim: checking the state of existing chars
type Characters struct {
	cfg             *model.SimConfig
	characters      []key.TargetID
	attributes      *attribute.Service
	customFunctions []testeval.ActionEval
}

func (s *Characters) ResetCharacters() {
	s.cfg.Characters = make([]*model.Character, 0, 4)
}

// AddCharacter adds a char to the list of SimConfig.
// This method is available only BEFORE calling Stub.StartSimulation
func (s *Characters) AddCharacter(char *model.Character) {
	s.cfg.Characters = append(s.cfg.Characters, char)
}

// AddCharacterEval adds a custom eval for the character at idx (based on AddCharacter order).
// This method is available only BEFORE calling Stub.StartSimulation
// If no Eval is provided, a basic ActionEval is generated for each missing one which will always Basic Attack.
func (s *Characters) AddCharacterEval(eval testeval.ActionEval, idx int) {
	for len(s.customFunctions) <= idx {
		s.customFunctions = append(s.customFunctions, nil)
	}
	s.customFunctions[idx] = eval
}

func (s *Characters) GetCharacterID(idx int) key.TargetID {
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

func (s *Characters) getCharacterEval(idx int) testeval.ActionEval {
	if len(s.customFunctions) > idx {
		return s.customFunctions[idx]
	}
	return nil
}
