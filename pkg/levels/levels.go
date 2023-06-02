package levels

type Level int

const (
	Trace Level = 10
	Debug Level = 20
	Info  Level = 30
	Warn  Level = 40
	Error Level = 50
	Fatal Level = 60
)
