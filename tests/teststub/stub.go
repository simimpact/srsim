package teststub

import (
	"fmt"
	"time"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/logic"
	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/simimpact/srsim/tests/eventchecker"
	"github.com/simimpact/srsim/tests/eventchecker/battlestart"
	"github.com/simimpact/srsim/tests/testcfg"
	"github.com/simimpact/srsim/tests/testcfg/testchar"
	"github.com/simimpact/srsim/tests/testcfg/testeval"
	"github.com/simimpact/srsim/tests/testcfg/testyaml"
	"github.com/stretchr/testify/suite"
)

type Stub struct {
	suite.Suite
	// AutoContinue determines if events are automatically piped.
	// When disabled, you must call s.Continue() after each Expect call.
	autoContinue  bool
	eventPipe     chan handler.Event
	haltSignaller chan struct{}
	isExpecting   bool // deadlock prevention for s.Continue

	// AutoRun determines if simulation will automatically run.
	// When disabled, you must call s.NextTurn() to queue the next TurnStart event.
	autoRun  bool
	turnPipe chan TurnCommand

	// cfg and evaluator are used to start a normal run
	cfg       *model.SimConfig
	cfgEval   *eval.Eval
	evaluator *evaluator
	simulator *simulation.Simulation

	// Characters gives access to various character-related testing actions
	Characters Characters
}

func (s *Stub) SetupTest() {
	s.eventPipe = make(chan handler.Event)
	s.haltSignaller = make(chan struct{})
	s.turnPipe = make(chan TurnCommand)
	s.cfg = testcfg.TestConfigTwoElites()
	s.autoContinue = true
	s.autoRun = true
	s.Characters = Characters{
		cfg:             s.cfg,
		characters:      nil,
		attributes:      nil,
		customFunctions: nil,
	}
	s.evaluator = newEvaluator()
}

func (s *Stub) TearDownTest() {
	fmt.Println("Test Finished")
	logging.InitLoggers()
	select {
	case <-s.eventPipe:
		s.haltSignaller <- struct{}{}
	default:
	}
	close(s.eventPipe)
	close(s.haltSignaller)
	close(s.turnPipe)
}

// StartSimulation handles the setup for starting the asynchronous sim run.
// Call this once you finish setting up test parameters.
func (s *Stub) StartSimulation() {
	logging.InitLoggers(NewTestLogger(s.eventPipe, s.haltSignaller))
	// if no chars are provided, we will add a dummy char
	if len(s.cfg.Characters) == 0 {
		s.Characters.AddCharacter(testchar.DummyChar())
	}
	var evalToUse logic.Eval
	if s.cfgEval != nil {
		evalToUse = s.cfgEval
	} else {
		evalToUse = s.evaluator
	}
	s.simulator = simulation.NewSimulation(s.cfg, evalToUse, 0)
	if !s.autoRun {
		s.simulator.Turn = newMockManager(s.turnPipe)
	}
	s.Characters.attributes = s.simulator.Attr
	go func() {
		itres, err := s.simulator.Run()
		if err != nil {
			s.FailNow("Simulation run error", err)
		}
		fmt.Println(itres)
	}()
	// start sim logic, fast-forward sim to BattleStart state, so we can initialize the remaining helper stuff
	s.Expect(battlestart.ExpectFor())
	// initialize the evaluator and Character based on current state
	s.Characters.characters = s.simulator.Characters()
	s.initEval()

	if !s.autoContinue {
		s.isExpecting = true
		s.Continue()
	}
}

// Expect handles all sorts of checks against events. Refer to eventchecker.EventChecker for more details.
func (s *Stub) Expect(checkers ...eventchecker.EventChecker) {
	for {
		var e handler.Event
		select {
		case e = <-s.eventPipe:
		case <-time.After(1 * time.Second):
			s.FailNow("Event not intercepted")
		}
		var toContinue bool
		var err error
		for i := range checkers {
			toContinue, err = checkers[i](e)
			if toContinue {
				continue
			}
			if err != nil {
				s.FailNow("Event Checker err", err)
				return
			}
			break
		}
		if toContinue {
			LogExpectSuccess("%#+v", e)
			if s.autoContinue {
				s.haltSignaller <- struct{}{}
			} else {
				s.isExpecting = true
			}
			return
		}
		LogExpectFalse("%#+v", e)
		s.haltSignaller <- struct{}{}
	}
}

func (s *Stub) initEval() {
	_ = s.evaluator.Init(s.simulator)
	for i := range s.cfg.Characters {
		eval := s.Characters.getCharacterEval(i)
		if eval == nil {
			eval = testeval.Default()
		}
		s.evaluator.registerAction(s.Characters.GetCharacterTargetID(i), eval)
	}
}

func (s *Stub) LoadYamlCfg(filepath string) {
	var err error
	s.cfg, s.cfgEval, err = testyaml.ParseConfig(filepath)
	s.Characters.cfg = s.cfg
	if err != nil {
		s.FailNow("Yaml unmarshal fail", err)
	}
}
