package teststub

import (
	"testing"

	"github.com/fatih/color"
)

func LogError(t *testing.T, format string, args ...interface{}) {
	const errString = "[TESTFRAME] [ERROR] "
	errFormat := errString + format
	color.Set(color.FgRed)
	t.Logf(errFormat, args...)
	color.Unset()
}

func LogExpectFalse(t *testing.T, format string, args ...interface{}) {
	const expectString = "[TESTFRAME] [EXPECT] [Expect Checker False] "
	format = expectString + format
	color.Set(color.FgYellow)
	t.Logf(format, args...)
	color.Unset()
}

func LogExpectSuccess(t *testing.T, format string, args ...interface{}) {
	const expectString = "[TESTFRAME] [EXPECT] [Expect Success] "
	format = expectString + format
	color.Set(color.FgGreen)
	t.Logf(format, args...)
	color.Unset()
}
