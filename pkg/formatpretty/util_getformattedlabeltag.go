package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/core"
)

var levelToLabel = map[core.Level]string{
	10: "TRACE",
	20: "DEBUG",
	30: " INFO",
	40: " WARN",
	50: "ERROR",
	60: "FATAL",
}

func GetFormattedLabelTag(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level

	// Determine color
	var bgColor core.Color
	var fgColor core.Color
	switch level {
	case core.LevelTrace:
		fgColor = core.FlagColorFgBrightBlack
		bgColor = core.FlagColorBgBoldBrightBlack
	case core.LevelDebug:
		fgColor = core.FlagColorFgBlue
		bgColor = core.FlagColorBgBoldBlue
	case core.LevelInfo:
		fgColor = core.FlagColorFgGreen
		bgColor = core.FlagColorBgBoldGreen
	case core.LevelWarn:
		fgColor = core.FlagColorFgYellow
		bgColor = core.FlagColorBgBoldYellow
	case core.LevelError:
		fgColor = core.FlagColorFgRed
		bgColor = core.FlagColorBgBoldRed
	case core.LevelFatal:
		fgColor = core.FlagColorFgMagenta
		bgColor = core.FlagColorBgBoldMagenta
	}

	// Create string
	return bgColor +
		fgColor + "# " + core.FlagColorFgBlack +
		levelToLabel[level] +
		fgColor + " " +
		core.FlagReset
}
