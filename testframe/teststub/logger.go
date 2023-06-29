package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"time"
)

type TestLogger struct {
	eventPipe     chan handler.Event
	haltSignaller chan struct{}
}

func (l *TestLogger) Log(e interface{}) {
	l.eventPipe <- e
	select {
	case <-l.haltSignaller:
	case <-time.After(1 * time.Second):
		panic("Test Halted")
	}
}

func NewTestLogger(pipe chan handler.Event, signal chan struct{}) *TestLogger {
	return &TestLogger{
		eventPipe:     pipe,
		haltSignaller: signal,
	}
}
