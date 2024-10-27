package basic

import (
	"testing"

	"github.com/simimpact/srsim/pkg/engine/prop"
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
	dummy := t.Characters.AddCharacter(testchar.DummyChar())
	dan := t.Characters.AddCharacter(testchar.DanHung())
	t.StartSimulation()
	t.QueueTurn(dan)
	t.Require().Equal(key.TargetID(1), dummy.ID())
	t.Require().Equal(key.TargetID(2), dan.ID())
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(dan.ID()))
	t.TerminateRun()
	t.Expect(termination.ExpectFor())
}

// Verify that we can manually continue with AutoContinue disabled
func (t *BasicTest) Test_ManualContinue() {
	t.SetAutoContinue(false)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.Character(key.DummyCharacter).ID()))
	t.Continue()
	t.Expect(termination.ExpectFor())
	t.Continue()
}

// Verify that LoadYamlCfg properly loads the correct cfg settings
func (t *BasicTest) Test_LoadYaml() {
	t.LoadYamlCfg(testyaml.SampleTestCfg)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.Character(key.DanHeng).ID()))
	t.Expect(termination.ExpectFor())
	t.Require().Equal(key.TargetID(1), t.Characters.Character(key.DanHeng).ID())
}

// Verify that LoadYamlCfg loads properly without a gcsl eval
func (t *BasicTest) Test_LoadYamlWithoutEval() {
	t.LoadYamlCfg(testyaml.SampleTestCfgNoEval)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.Character(key.DanHeng).ID()))
	t.Expect(termination.ExpectFor())
	t.Require().Equal(key.TargetID(1), t.Characters.Character(key.DanHeng).ID())
}

// Verify that t.Continue does not deadlock if AutoContinue is disabled, or if it is called twice with AutoContinue enabled
func (t *BasicTest) Test_AutoContinue_Deadlock() {
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(t.Characters.Character(key.DummyCharacter).ID()))
	t.Continue()
	t.SetAutoContinue(false)
	t.Expect(termination.ExpectFor())
	t.Continue()
	t.Continue()
}

// Verify that MaxTraces creates a proper TraceMap
func (t *BasicTest) Test_TraceMap_Registration() {
	danModel := testchar.DanHung()
	danModel.Traces = testcfg.MaxTraces()
	dan := t.Characters.AddCharacter(danModel)
	t.StartSimulation()
	t.Expect(turnstart.ExpectFor(), turnstart.CurrentTurnIs(dan.ID()))
	t.Expect(termination.ExpectFor())
	// 16% from OSR, 18% from Traces
	dan.Equal(prop.ATKPercent, 0.34)
}
