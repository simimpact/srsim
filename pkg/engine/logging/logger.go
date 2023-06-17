package logging

type Logger struct {
	entries []string
}

func (l *Logger) Pop() string {
	if l == nil {
		return ""
	}
	if len(l.entries) > 0 {
		res := l.entries[0]
		l.entries = l.entries[1:]
		return res
	}
	return ""
}
