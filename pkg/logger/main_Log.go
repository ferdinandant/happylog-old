package logger

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/env"
	"github.com/ferdinandant/happylog/pkg/formatpretty"
)

func Log(level core.Level, msg string, maybeCtx ...interface{}) {
	var ctxPtr *interface{} = nil
	if len(maybeCtx) > 0 {
		ctxPtr = &maybeCtx[0]
	}

	appName := env.EnvAppName
	if level >= core.LevelError {
		appName = "webacd-desktop"
	}

	now := time.Now()
	formatLogOpts := &core.FormatLogOpts{
		Level:   &level,
		AppName: &appName,
		Now:     &now,
		Msg:     &msg,
		CtxPtr:  ctxPtr,
	}
	formatpretty.FormatLog(formatLogOpts)
}
