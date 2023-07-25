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
	dummyModel := testchar.DummyChar()
	dummyModel.LightCone = testcone.QuidProQuo()
	t.Characters.ResetCharacters()
	dan := t.Characters.AddCharacter(testchar.DanHung())
	dummy := t.Characters.AddCharacter(dummyModel)
	t.SetAutoRun(false)
	t.SetAutoContinue(false)
	t.StartSimulation()
	t.NextTurn(dummy)
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(dummy.ID()))
	t.Continue()
	t.Expect(turnend.ExpectFor())
	t.Assert.Equal(0, dummy.Stats().Energy())
	t.Assert.Equal(8, dan.Stats().Energy())
	t.Continue()
}
