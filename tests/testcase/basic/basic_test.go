package basic

import (
	"testing"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/tests/eventchecker/termination"
	"github.com/simimpact/srsim/tests/eventchecker/turnstart"
	"github.com/simimpact/srsim/tests/testcfg"
	"github.com/simimpact/srsim/tests/testcfg/testchar"
	"github.com/simimpact/srsim/tests/testcfg/testyaml"
	"github.com/simimpact/srsim/tests/teststub"
	"github.com/stretchr/testify/suite"
)

type BasicTest struct {
	teststub.Stub
}

func TestBasicTest(t *testing.T) {
	suite.Run(t, new(BasicTest))
}

// Sanity check
func (t *BasicTest) Test_Framework() {
	t.StartSimulation()
	t.Expect(termination.ExpectFor())
}

// Verify that we can manually dictate turns
func (t *BasicTest) Test_TurnLogic() {
	t.SetAutoRun(false)
	t.Characters.AddCharacter(testchar.DummyChar())
	t.Characters.AddCharacter(testchar.DanHung())
	t.StartSimulation()
	dummyID := t.Characters.GetCharacterTargetID(0)
	danID := t.Characters.GetCharacterTargetID(1)
	t.Require().Equal(key.TargetID(1), dummyID)
	t.Require().Equal(key.TargetID(2), danID)
	t.NextTurn(danID)
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(danID))
	t.TerminateRun()
	t.Expect(termination.ExpectFor())
}

// Verify that we can manually continue with AutoContinue disabled
func (t *BasicTest) Test_ManualContinue() {
	t.SetAutoContinue(false)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.GetCharacterTargetID(0)))
	t.Continue()
	t.Expect(termination.ExpectFor())
	t.Continue()
}

// Verify that LoadYamlCfg properly loads the correct cfg settings
func (t *BasicTest) Test_LoadYaml() {
	t.LoadYamlCfg(testyaml.SampleTestCfg)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.GetCharacterTargetID(0)))
	t.Expect(termination.ExpectFor())
	t.Require().Equal(0, t.Characters.CharacterIdx(key.DanHeng))
}

// Verify that t.Continue does not deadlock if AutoContinue is disabled, or if it is called twice with AutoContinue enabled
func (t *BasicTest) Test_AutoContinue_Deadlock() {
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.GetCharacterTargetID(0)))
	t.Continue()
	t.SetAutoContinue(false)
	t.Expect(termination.ExpectFor())
	t.Continue()
	t.Continue()
}

// Verify that MaxTraces creates a proper TraceMap
func (t *BasicTest) Test_TraceMap_Registration() {
	dan := testchar.DanHung()
	dan.Traces = testcfg.MaxTraces()
	t.Characters.AddCharacter(dan)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.GetCharacterTargetID(0)))
	t.Expect(termination.ExpectFor())
	info := t.Characters.GetCharacterInfo(t.Characters.CharacterIdx(key.DanHeng))
	// DH base atk is 1186.8 with OSR, here we test that minor atk traces were activated
	t.Require().Less(float64(1187), info.ATK())
}
