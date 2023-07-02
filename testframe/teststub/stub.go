package teststub

import (
	"fmt"
	"time"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/simimpact/srsim/testframe/eventchecker"
	"github.com/simimpact/srsim/testframe/eventchecker/battlestart"
	"github.com/simimpact/srsim/testframe/testcfg"
	"github.com/stretchr/testify/suite"
)

type Stub struct {
	suite.Suite
	// AutoContinue determines if events are automatically piped.
	// When disabled, you must call s.Continue() after each Expect call.
	autoContinue  bool
	eventPipe     chan handler.Event
	haltSignaller chan struct{}

	// AutoRun determines if simulation will automatically run.
	// When disabled, you must call s.NextTurn() to queue the next TurnStart event.
	autoRun  bool
	turnPipe chan TurnCommand

	// cfg and eval are used to start a normal run
	cfg       *model.SimConfig
	eval      *eval.Eval
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
	s.Characters = Characters{cfg: s.cfg}
}

func (s *Stub) TearDownTest() {
	fmt.Println("Test Finished")
	logging.Singleton = logging.NewNilLogger()
	select {
	case <-s.eventPipe:
		s.haltSignaller <- struct{}{}
	default:
	}
	close(s.eventPipe)
	close(s.haltSignaller)
	close(s.turnPipe)
}

func (s *Stub) StartSimulation() {
	l := NewTestLogger(s.eventPipe, s.haltSignaller)
	logging.InitLogger(l)
	s.eval = testcfg.GenerateStandardTestEval(s.cfg)
	s.simulator = simulation.NewSimulation(s.cfg, s.eval, 0)
	if !s.autoRun {
		s.simulator.Turn = newMockManager(s.turnPipe)
	}
	s.Characters.characters = s.simulator.Characters()
	s.Characters.attributes = s.simulator.Attr
	go func() {
		itres, err := s.simulator.Run()
		if err != nil {
			s.FailNow("Simulation run error", err)
		}
		fmt.Println(itres)
	}()
	// common start sim logic, fast-forward sim to BattleStart state
	s.Expect(battlestart.ExpectFor())
	s.RefreshCharacters()
	if !s.autoContinue {
		s.Continue()
	}
}

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
			} else {
				if err != nil {
					s.FailNow("Event Checker err", err)
					return
				}
				break
			}
		}
		if !toContinue {
			LogExpectFalse("%#+v", e)
			s.haltSignaller <- struct{}{}
		}
		if toContinue {
			LogExpectSuccess("%#+v", e)
			if s.autoContinue {
				s.haltSignaller <- struct{}{}
			}
			return
		}
	}
}

// Continue resumes the simulation. This must be called after each Expect if AutoContinue is disabled.
func (s *Stub) Continue() {
	s.haltSignaller <- struct{}{}
}

func (s *Stub) SetAutoContinue(cont bool) {
	s.autoContinue = cont
}

func (s *Stub) SetAutoRun(cont bool) {
	s.autoRun = cont
}

// TerminateRun pipes a command with an astronomical AV to immediately exceed the cycle limit, ending the run
func (s *Stub) TerminateRun() {
	go func() {
		s.turnPipe <- TurnCommand{Next: s.Characters.GetCharacterId(0), Av: 100000}
	}()
}

// NextTurn queues the next turn without using up any AV cost
func (s *Stub) NextTurn(id key.TargetID) {
	go func() {
		s.turnPipe <- TurnCommand{Next: id}
	}()
}

func (s *Stub) RefreshCharacters() {
	s.Characters.characters = s.simulator.Characters()
}
