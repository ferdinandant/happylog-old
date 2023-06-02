package main

import (
	"github.com/ferdinandant/happylog/pkg/levels"
	"github.com/ferdinandant/happylog/pkg/logintf"
)

func main() {
	Trace("Hello ah")
	Debug("Hello ah")
	Info("Hello ah")
	Warn("Hello ah")
	Error("Hello ah")
	Fatal("Hello ah")
}

func Trace(msg string, ctx ...interface{}) {
	logintf.WriteLog(levels.Trace, msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	logintf.WriteLog(levels.Debug, msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	logintf.WriteLog(levels.Info, msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	logintf.WriteLog(levels.Warn, msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	logintf.WriteLog(levels.Error, msg, ctx...)
}

func Fatal(msg string, ctx ...interface{}) {
	logintf.WriteLog(levels.Fatal, msg, ctx...)
}
