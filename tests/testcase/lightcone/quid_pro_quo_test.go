package lightcone

import (
	"testing"

	"github.com/simimpact/srsim/tests/eventchecker/turnend"
	"github.com/simimpact/srsim/tests/eventchecker/turnstart"
	"github.com/simimpact/srsim/tests/testcfg/testchar"
	"github.com/simimpact/srsim/tests/testcfg/testcone"
	"github.com/simimpact/srsim/tests/teststub"
	"github.com/stretchr/testify/suite"
)

type QPQTest struct {
	teststub.Stub
}

func TestQPQTest(t *testing.T) {
	suite.Run(t, new(QPQTest))
}

// Basic Testcase for checking that 8 energy gets added to a char when holder's turn starts
func (t *QPQTest) Test_EnergyAdd() {
	dummyChar := testchar.DummyChar()
	dummyChar.LightCone = testcone.QuidProQuo()
	t.Characters.ResetCharacters()
	t.Characters.AddCharacter(testchar.DanHung())
	t.Characters.AddCharacter(dummyChar)
	t.SetAutoRun(false)
	t.SetAutoContinue(false)
	t.StartSimulation()
	dummy := t.Characters.GetCharacterTargetID(1)
	t.NextTurn(dummy)
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(dummy))
	t.Continue()
	t.Expect(turnend.ExpectFor())
	info := t.Characters.GetCharacterInfo(0)
	t.Require().Equal(float64(8), info.Energy())
	t.Continue()
}
