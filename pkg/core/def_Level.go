package core

type Level int

const (
	LevelTrace Level = 10
	LevelDebug Level = 20
	LevelInfo  Level = 30
	LevelWarn  Level = 40
	LevelError Level = 50
	LevelFatal Level = 60
)
