package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/core"
)

func GetFormattedMessage(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	msg := *logOpts.Msg

	// Determine color
	var fgColor core.Color
	switch level {
	case core.LevelTrace:
		fgColor = core.FlagColorFgBrightBlack
	case core.LevelDebug:
		fgColor = core.FlagColorFgBlue
	case core.LevelInfo:
		fgColor = core.FlagColorFgGreen
	case core.LevelWarn:
		fgColor = core.FlagColorFgYellow
	case core.LevelError:
		fgColor = core.FlagColorFgRed
	case core.LevelFatal:
		fgColor = core.FlagColorFgMagenta
	}

	// Create string
	return FormatTextWithColor(fgColor, msg)
}
