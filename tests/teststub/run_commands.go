package teststub

// Continue resumes the simulation. This must be called after each Expect if AutoContinue is disabled.
// This function does nothing if an Expect is not called prior to this or if AutoContinue is enabled.
func (s *Stub) Continue() {
	if !s.isExpecting {
		return
	}
	s.isExpecting = false
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
		s.turnPipe <- TurnCommand{Next: s.Characters.testChars[0].ID(), Av: 100000}
	}()
}

// NextTurn queues the next turn without using up any AV cost
func (s *Stub) NextTurn(char *Character) {
	go func() {
		s.turnPipe <- TurnCommand{Next: char.ID(), Av: 0}
	}()
}
