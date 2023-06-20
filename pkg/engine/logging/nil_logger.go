package logging

type nilLogger struct{}

func (l *nilLogger) Log(_ interface{}) {}

func (l *nilLogger) Flush() []byte {
	return nil
}
