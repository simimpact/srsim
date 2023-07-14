package relic

import (
	"testing"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/tests/testcfg/testchar"
	"github.com/simimpact/srsim/tests/testcfg/testrelic"
	"github.com/simimpact/srsim/tests/teststub"
	"github.com/stretchr/testify/suite"
)

type MusketeerTest struct {
	teststub.Stub
}

func TestBasicTest(t *testing.T) {
	suite.Run(t, new(MusketeerTest))
}

// Verify that Musketeer Set adds atk% and spd
func (t *MusketeerTest) Test_Tracemap_Registration() {
	dan := testchar.DanHung()
	dan.Relics = testrelic.MusketeerStatlessSet()
	t.Characters.AddCharacter(dan)
	t.StartSimulation()
	info := t.Characters.GetCharacterInfo(t.Characters.CharacterIdx(key.DanHeng))
	// DH base atk is 1186.8 with OSR and base spd 110
	t.Require().Less(float64(111), info.SPD())
	t.Require().Less(float64(1187), info.ATK())
}
