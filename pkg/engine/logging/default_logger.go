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
	res2, err := json.Marshal(wrappedEv)
	if err != nil {
		return
	}
	l.entries = append(l.entries, string(res2))
}

func (l *defaultLogger) PrintToConsole() {
	if l == nil {
		return
	}
	for i := range l.entries {
		fmt.Println(l.entries[i])
	}
}

func (l *defaultLogger) Flush() []byte {
	res, err := json.Marshal(l.entries)
	if err != nil {
		return nil
	}
	return res
}
