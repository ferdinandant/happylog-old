package main

import (
	"github.com/ferdinandant/happylog/pkg/logintf"
	"github.com/ferdinandant/happylog/pkg/types"
)

func main() {
	Trace("Hello ah")
	Debug("Hello ah")
	Info("Hello ah")
	Warn("Hello ah")
	Error("Hello ah")
	Fatal("Hello ah")
}

func Trace(msg string, ctx ...*interface{}) {
	logintf.WriteLog(types.LevelTrace, msg, ctx...)
}

func Debug(msg string, ctx ...*interface{}) {
	logintf.WriteLog(types.LevelDebug, msg, ctx...)
}

func Info(msg string, ctx ...*interface{}) {
	logintf.WriteLog(types.LevelInfo, msg, ctx...)
}

func Warn(msg string, ctx ...*interface{}) {
	logintf.WriteLog(types.LevelWarn, msg, ctx...)
}

func Error(msg string, ctx ...*interface{}) {
	logintf.WriteLog(types.LevelError, msg, ctx...)
}

func Fatal(msg string, ctx ...*interface{}) {
	logintf.WriteLog(types.LevelFatal, msg, ctx...)
}
