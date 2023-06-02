package logintf

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/formatpretty"
	"github.com/ferdinandant/happylog/pkg/levels"
)

func WriteLog(level levels.Level, msg string, ctx ...interface{}) {
	now := time.Now()
	formatpretty.Log(level, now, msg, ctx...)
}
