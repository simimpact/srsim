package logging

var Singleton Logger

type Logger interface {
	Log(e interface{})
	Flush() string
}

func init() {
	Singleton = &nilLogger{}
}

func Log(e interface{}) {
	Singleton.Log(e)
}

func Flush() string {
	return Singleton.Flush()
}

func IsNil() bool {
	_, ok := Singleton.(*nilLogger)
	return ok
}
