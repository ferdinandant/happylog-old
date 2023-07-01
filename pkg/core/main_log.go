package core

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/env"
	"github.com/ferdinandant/happylog/pkg/formatpretty"
	"github.com/ferdinandant/happylog/pkg/types"
)

func Log(level types.Level, msg string, maybeCtx ...interface{}) {
	var ctxPtr *interface{} = nil
	if len(maybeCtx) > 0 {
		ctxPtr = &maybeCtx[0]
	}

	appName := env.EnvAppName
	if level >= types.LevelError {
		appName = "webacd-desktop"
	}

	now := time.Now()
	formatLogOpts := &types.FormatLogOpts{
		Level:   &level,
		AppName: &appName,
		Now:     &now,
		Msg:     &msg,
		CtxPtr:  ctxPtr,
	}
	formatpretty.FormatLog(formatLogOpts)
}
