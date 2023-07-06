package logging

import (
	"encoding/json"
	"strings"
)

type DefaultLogger struct {
	entries []string
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{entries: make([]string, 0, 10)}
}

// Log registers an event to the DefaultLogger entry
func (l *DefaultLogger) Log(e interface{}) {
	wrappedEv := Wrap(e)
	res, err := json.Marshal(wrappedEv)
	if err != nil {
		return
	}
	l.entries = append(l.entries, string(res))
}

func (l *DefaultLogger) Flush() string {
	var sb strings.Builder
	for i := range l.entries {
		sb.WriteString(l.entries[i])
		sb.WriteByte('\n')
	}
	l.entries = make([]string, 0, 10)
	return sb.String()
}
