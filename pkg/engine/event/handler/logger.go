package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Singleton *logger

type logger struct {
	entries []string
}

func (l *logger) Pop() string {
	if l == nil {
		return ""
	}
	if len(l.entries) > 0 {
		res := l.entries[0]
		l.entries = l.entries[1:]
		return res
	}
	return ""
}

func InitLogger() {
	Singleton = &logger{entries: make([]string, 0, 10)}
}

// Log registers an event to the logger entry
func (l *logger) Log(e interface{}) {
	// TODO iCarus: benchmark performance to see if we should skip nil checks
	if l == nil {
		return
	}
	fmt.Printf("Logging event %T%v\n", e, e)
	res, err := json.Marshal(e)
	if err != nil {
		return
	}
	l.entries = append(l.entries, string(res))
}

func (l *logger) PrintToConsole() {
	if l == nil {
		return
	}
	for i := range l.entries {
		fmt.Println(l.entries[i])
	}
}

func (l *logger) FlushAsString() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for cur := l.Pop(); cur != ""; {
		buf.WriteString(cur)
	}
	buf.WriteByte('}')
	return buf.String()
}
