package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/attribute"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/tests/testcfg/testeval"
)

// Characters provide interfaces for
// -- before sim starts: adding new chars to the sim, modifying existing ones
// -- during sim: checking the state of existing chars
type Characters struct {
	cfg             *model.SimConfig
	testChars       []*Character
	attributes      attribute.Manager
	customFunctions []testeval.ActionEval
}

// Character is a test instance of a character which provides various Checker methods and retrievers
type Character struct {
	model *model.Character
	key   key.TargetID
	stat  func(id key.TargetID) *info.Stats
}

func (s *Characters) ResetCharacters() {
	s.cfg.Characters = make([]*model.Character, 0, 4)
	s.testChars = make([]*Character, 0, 4)
}

// AddCharacter adds a char to the list of SimConfig.
// This method is available only BEFORE calling Stub.StartSimulation
// This returns a testchar instance that will be useful for assertions later on
func (s *Characters) AddCharacter(char *model.Character) *Character {
	s.cfg.Characters = append(s.cfg.Characters, char)
	s.testChars = append(s.testChars, newCharacter(char))
	return s.testChars[len(s.testChars)-1]
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

// GetCharacterTargetID fetches the key.TargetID value for the character at idx that the Sim is using.
// This is used for most state checks
func (s *Characters) GetCharacterTargetID(idx int) key.TargetID {
	if idx >= len(s.testChars) {
		LogError("invalid idx %d, insufficient characters", idx)
		panic("Invalid index")
	}
	return s.testChars[idx].key
}

// Character fetches the test instance of the first character with the given key. This is useful for adding clarity
// or todo: cases where the character is added post-combat start
func (s *Characters) Character(key key.Character) *Character {
	for _, v := range s.testChars {
		if v.model.Key == key.String() {
			return v
		}
	}
	LogError("Character Key %s is not in the SimConfig", key.String())
	return nil
}

// GetCharacterInfo fetches the info.Stats value for the character at idx. Useful for verifying energy state etc.
func (s *Characters) GetCharacterInfo(idx int) *info.Stats {
	if idx >= len(s.testChars) {
		LogError("invalid idx %d, insufficient characters", idx)
	}
	return s.testChars[idx].Stats()
}

func (s *Characters) getCharacterEval(idx int) testeval.ActionEval {
	if len(s.customFunctions) > idx {
		return s.customFunctions[idx]
	}
	return nil
}

func (s *Characters) init(characters []key.TargetID) {
	if len(s.testChars) < len(characters) { // eval loaded config, need to populate testchars
		for i := range s.cfg.Characters {
			s.testChars = append(s.testChars, newCharacter(s.cfg.Characters[i]))
		}
	}
	for i := range s.testChars {
		s.testChars[i].key = characters[i]
		s.testChars[i].stat = s.attributes.Stats
	}
}
