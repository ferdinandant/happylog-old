package main

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/types"
)

type StructB = struct {
	ValueA bool
	ValueB string
}

type StructA = struct {
	ValueA   int
	ValueB   string
	ValueC   *int
	ValueD   *string
	StructB1 StructB
	StructB2 *StructB
}

func main() {
	aa := StructA{}
	bb := &StructB{}
	Trace("Hello ah", "Hello!")
	Debug("Hello ah", 123)
	Info("Hello ah", aa)
	Warn("Hello ah", time.Now())
	Error("Hello ah", bb)
	Fatal("Hello ah")
}

func Trace(msg string, ctx ...interface{}) {
	core.WriteLog(types.LevelTrace, msg, ctx...)
}

func Debug(msg string, ctx ...interface{}) {
	core.WriteLog(types.LevelDebug, msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	core.WriteLog(types.LevelInfo, msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	core.WriteLog(types.LevelWarn, msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	core.WriteLog(types.LevelError, msg, ctx...)
}

func Fatal(msg string, ctx ...interface{}) {
	core.WriteLog(types.LevelFatal, msg, ctx...)
}
