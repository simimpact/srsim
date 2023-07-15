package teststub

import (
	"os"

	"github.com/fatih/color"
)

type logLevel uint8

const (
	Info logLevel = iota
	Error
	Disabled
)

var curLogLevel = Info

func init() {
	switch os.Getenv("LogLevel") {
	case "Info":
		curLogLevel = Info
	case "Error":
		curLogLevel = Error
	case "Disabled":
		curLogLevel = Disabled
	default:
		curLogLevel = Info
	}
}

func LogError(format string, args ...interface{}) {
	if curLogLevel == Disabled {
		return
	}
	const errString = "[TESTFRAME] [ERROR] "
	errFormat := errString + format
	color.Red(errFormat, args...)
}

func LogExpectFalse(format string, args ...interface{}) {
	if curLogLevel > Info {
		return
	}
	const expectString = "[TESTFRAME] [EXPECT] [Expect Checker False] "
	format = expectString + format
	color.Yellow(format, args...)
}

func LogExpectSuccess(format string, args ...interface{}) {
	if curLogLevel > Info {
		return
	}
	const expectString = "[TESTFRAME] [EXPECT] [Expect Success] "
	format = expectString + format
	color.Green(format, args...)
}
