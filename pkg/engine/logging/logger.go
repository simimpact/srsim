package logging

var Singleton Logger

type Logger interface {
	Log(e interface{})
}

func init() {
	Singleton = &nilLogger{}
}

func InitLogger(l Logger) {
	Singleton = l
}

func Log(e interface{}) {
	Singleton.Log(e)
}
