package main

import (
	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/logger"
)

func Trace(msg string, ctx ...interface{}) {
	logger.Log(core.LevelTrace, msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	logger.Log(core.LevelDebug, msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	logger.Log(core.LevelInfo, msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	logger.Log(core.LevelWarn, msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	logger.Log(core.LevelError, msg, ctx...)
}

func Fatal(msg string, ctx ...interface{}) {
	logger.Log(core.LevelFatal, msg, ctx...)
}
