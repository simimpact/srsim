package teststub

import (
	"github.com/fatih/color"
)

func LogError(format string, args ...interface{}) {
	const errString = "[TESTFRAME] [ERROR] "
	errFormat := errString + format
	color.Red(errFormat, args...)
}

func LogExpectFalse(format string, args ...interface{}) {
	const expectString = "[TESTFRAME] [EXPECT] [Expect Checker False] "
	format = expectString + format
	color.Yellow(format, args...)
}

func LogExpectSuccess(format string, args ...interface{}) {
	const expectString = "[TESTFRAME] [EXPECT] [Expect Success] "
	format = expectString + format
	color.Green(format, args...)
}
