package logging

type nilLogger struct{}

func (l *nilLogger) Log(_ interface{}) {}

func NewNilLogger() Logger {
	return new(nilLogger)
}
