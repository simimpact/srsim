package basic

import (
	"testing"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/tests/eventchecker/termination"
	"github.com/simimpact/srsim/tests/eventchecker/turnstart"
	"github.com/simimpact/srsim/tests/testcfg/testchar"
	"github.com/simimpact/srsim/tests/teststub"
	"github.com/stretchr/testify/suite"
)

type BasicTest struct {
	teststub.Stub
}

func TestBasicTest(t *testing.T) {
	suite.Run(t, new(BasicTest))
}

func (t *BasicTest) Test_Framework() {
	t.StartSimulation()
	t.Expect(termination.ExpectFor())
}

func (t *BasicTest) Test_TurnLogic() {
	t.SetAutoRun(false)
	t.Characters.AddCharacter(testchar.DummyChar())
	t.Characters.AddCharacter(testchar.DanHung())
	t.StartSimulation()
	dummyID := t.Characters.GetCharacterID(0)
	danID := t.Characters.GetCharacterID(1)
	t.Require().Equal(key.TargetID(1), dummyID)
	t.Require().Equal(key.TargetID(2), danID)
	t.NextTurn(danID)
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(danID))
	t.TerminateRun()
	t.Expect(termination.ExpectFor())
}

func (t *BasicTest) Test_ManualContinue() {
	t.SetAutoContinue(false)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.GetCharacterID(0)))
	t.Continue()
	t.Expect(termination.ExpectFor())
	t.Continue()
}
