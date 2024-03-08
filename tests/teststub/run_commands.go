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
	s.Turn.queue(-1)
}

// QueueTurn queues the next turn without using up any AV cost
func (s *Stub) QueueTurn(char *Character) {
	s.Turn.queue(char.ID())
}
