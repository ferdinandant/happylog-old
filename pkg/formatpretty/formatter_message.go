package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/types"
)

func GetFormattedMessage(logOpts *types.LogOpts) string {
	level := *logOpts.Level
	msg := *logOpts.Msg

	// Determine color
	var fgColor string
	switch level {
	case types.LevelTrace:
		fgColor = FlagColorFgBrightBlack
	case types.LevelDebug:
		fgColor = FlagColorFgBlue
	case types.LevelInfo:
		fgColor = FlagColorFgGreen
	case types.LevelWarn:
		fgColor = FlagColorFgYellow
	case types.LevelError:
		fgColor = FlagColorFgRed
	case types.LevelFatal:
		fgColor = FlagColorFgMagenta
	}

	// Create string
	return fgColor + msg + FlagReset
}
