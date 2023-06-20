package logging

import (
	"encoding/json"
	"fmt"
	"strings"
)

type defaultLogger struct {
	entries []string
}

type logEventWrapper struct {
	EventName string
	Data      interface{}
}

func InitDefaultLogger() {
	Singleton = &defaultLogger{entries: make([]string, 0, 10)}
}

// Log registers an event to the defaultLogger entry
func (l *defaultLogger) Log(e interface{}) {
	wrappedEv := &logEventWrapper{
		EventName: strings.TrimPrefix(fmt.Sprintf("%T", e), "event."),
		Data:      e,
	}
	res, err := json.Marshal(wrappedEv)
	if err != nil {
		return
	}
	l.entries = append(l.entries, string(res))
}

func (l *defaultLogger) Flush() string {
	var sb strings.Builder
	for i := range l.entries {
		sb.WriteString(l.entries[i])
		sb.WriteByte('\n')
	}
	l.entries = make([]string, 0, 10)
	return sb.String()
}
