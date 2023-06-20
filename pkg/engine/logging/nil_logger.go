package logging

type nilLogger struct{}

func (l *nilLogger) Log(_ interface{}) {}
