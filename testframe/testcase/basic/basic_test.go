package basic

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/testframe/eventchecker/battlestart"
	"github.com/simimpact/srsim/testframe/eventchecker/termination"
	"github.com/simimpact/srsim/testframe/eventchecker/turnstart"
	"github.com/simimpact/srsim/testframe/testcfg/testchar"
	"github.com/simimpact/srsim/testframe/teststub"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BasicTest struct {
	teststub.Stub
}

func TestBasicTest(t *testing.T) {
	suite.Run(t, new(BasicTest))
}

func (t *BasicTest) Test_Framework() {
	t.StartSimulation()
	t.Expect(battlestart.ExpectFor())
	t.Expect(termination.ExpectFor())
}

func (t *BasicTest) Test_TurnLogic() {
	t.SetAutoRun(false)
	t.Characters.AddCharacter(testchar.DanHung())
	t.StartSimulation()
	dummyId := t.Characters.GetCharacterId(0)
	danId := t.Characters.GetCharacterId(1)
	t.Require().Equal(key.TargetID(1), dummyId)
	t.Require().Equal(key.TargetID(2), danId)
	t.NextTurn(danId)
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(danId))
	t.TerminateRun()
	t.Expect(termination.ExpectFor())
}

func (t *BasicTest) Test_ManualContinue() {
	t.SetAutoContinue(false)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.GetCharacterId(0)))
	t.Continue()
	t.Expect(termination.ExpectFor())
	t.Continue()
}
