package logging

import (
	"fmt"
	"strings"
)

var loggers []Logger

type Logger interface {
	Log(e any)
}

type LogWrapper struct {
	EventName string `json:"name"`
	Data      any    `json:"event"`
}

func InitLoggers(ls ...Logger) {
	loggers = ls
}

func Log(e any) {
	for _, l := range loggers {
		l.Log(e)
	}
}

func Wrap(e any) *LogWrapper {
	return &LogWrapper{
		EventName: strings.TrimPrefix(fmt.Sprintf("%T", e), "event."),
		Data:      e,
	}
}
