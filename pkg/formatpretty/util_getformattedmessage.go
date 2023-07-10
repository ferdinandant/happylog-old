package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/colors"
	"github.com/ferdinandant/happylog/pkg/core"
)

func GetFormattedMessage(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	msg := *logOpts.Msg

	// Determine color
	var fgColor colors.Color
	switch level {
	case core.LevelTrace:
		fgColor = colors.FlagColorFgBrightBlack
	case core.LevelDebug:
		fgColor = colors.FlagColorFgBlue
	case core.LevelInfo:
		fgColor = colors.FlagColorFgGreen
	case core.LevelWarn:
		fgColor = colors.FlagColorFgYellow
	case core.LevelError:
		fgColor = colors.FlagColorFgRed
	case core.LevelFatal:
		fgColor = colors.FlagColorFgMagenta
	}

	// Create string
	return colors.FormatTextWithColor(fgColor, msg)
}
