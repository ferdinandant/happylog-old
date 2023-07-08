package main

import (
	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/types"
)

func Trace(msg string, ctx ...interface{}) {
	core.Log(types.LevelTrace, msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	core.Log(types.LevelDebug, msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	core.Log(types.LevelInfo, msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	core.Log(types.LevelWarn, msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	core.Log(types.LevelError, msg, ctx...)
}

func Fatal(msg string, ctx ...interface{}) {
	core.Log(types.LevelFatal, msg, ctx...)
}
