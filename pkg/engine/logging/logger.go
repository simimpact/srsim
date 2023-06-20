package logging

var Singleton Logger

type Logger interface {
	Log(e interface{})
	Flush() []byte
}

func init() {
	Singleton = &nilLogger{}
}

func Log(e interface{}) {
	Singleton.Log(e)
}

func Flush() []byte {
	return Singleton.Flush()
}

func PrintToConsole() {
	if l, ok := Singleton.(*defaultLogger); ok {
		l.PrintToConsole()
	}
}

func IsNil() bool {
	_, ok := Singleton.(*nilLogger)
	return ok
}
