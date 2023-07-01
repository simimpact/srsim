package teststub

import (
	"fmt"
	"time"

	"github.com/simimpact/srsim/pkg/engine/event/handler"
)

type TestLogger struct {
	eventPipe     chan handler.Event
	haltSignaller chan struct{}
}

// type logEventWrapper struct {
// 	EventName string
// 	Data      interface{}
// }

func (l *TestLogger) Log(e interface{}) {
	fmt.Printf("Event Received: %+#v\n", e)
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
