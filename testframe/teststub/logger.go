package teststub

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
)

type TestLogger struct {
	eventPipe chan handler.Event
}

func (l *TestLogger) Log(e interface{}) {
	l.eventPipe <- e
}

func NewTestLogger(pipe chan handler.Event) *TestLogger {
	return &TestLogger{
		eventPipe: pipe,
	}
}
