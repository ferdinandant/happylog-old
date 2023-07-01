package core

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/env"
	"github.com/ferdinandant/happylog/pkg/formatpretty"
	"github.com/ferdinandant/happylog/pkg/types"
)

func Log(level types.Level, msg string, maybeCtx ...interface{}) {
	var ctx *interface{} = nil
	if len(maybeCtx) > 0 {
		ctx = &maybeCtx[0]
	}

	appName := env.EnvAppName
	if level >= types.LevelError {
		appName = "webacd-desktop"
	}

	now := time.Now()
	logOpts := &types.LogOpts{
		Level:   &level,
		AppName: &appName,
		Now:     &now,
		Msg:     &msg,
		Ctx:     ctx,
	}
	formatpretty.FormatLog(logOpts)
}
