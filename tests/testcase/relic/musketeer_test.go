package relic

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/prop"
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
func (t *MusketeerTest) Test_Musketeer_4p() {
	danModel := testchar.DanHung()
	danModel.Relics = testrelic.MusketeerStatlessSet()
	dan := t.Characters.AddCharacter(danModel)
	t.StartSimulation()
	dan.Equal(prop.SPDPercent, 0.06)
	// OSR +16, Musketeer +12
	dan.Equal(prop.ATKPercent, 0.28)
}
