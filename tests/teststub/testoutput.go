package teststub

import (
	"testing"
)

func LogError(t *testing.T, format string, args ...interface{}) {
	const errString = "[TESTFRAME] [ERROR] "
	errFormat := errString + format
	t.Logf(errFormat, args...)
}

func LogExpectFalse(t *testing.T, format string, args ...interface{}) {
	const expectString = "[TESTFRAME] [EXPECT] [Expect Checker False] "
	format = expectString + format
	t.Logf(format, args...)
}

func LogExpectSuccess(t *testing.T, format string, args ...interface{}) {
	const expectString = "[TESTFRAME] [EXPECT] [Expect Success] "
	format = expectString + format
	t.Logf(format, args...)
}
