package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/colors"
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
	var bgColor colors.Color
	var fgColor colors.Color
	switch level {
	case core.LevelTrace:
		fgColor = colors.FlagColorFgBrightBlack
		bgColor = colors.FlagColorBgBoldBrightBlack
	case core.LevelDebug:
		fgColor = colors.FlagColorFgBlue
		bgColor = colors.FlagColorBgBoldBlue
	case core.LevelInfo:
		fgColor = colors.FlagColorFgGreen
		bgColor = colors.FlagColorBgBoldGreen
	case core.LevelWarn:
		fgColor = colors.FlagColorFgYellow
		bgColor = colors.FlagColorBgBoldYellow
	case core.LevelError:
		fgColor = colors.FlagColorFgRed
		bgColor = colors.FlagColorBgBoldRed
	case core.LevelFatal:
		fgColor = colors.FlagColorFgMagenta
		bgColor = colors.FlagColorBgBoldMagenta
	}

	// Create string
	return bgColor +
		fgColor + "# " + colors.FlagColorFgBlack +
		levelToLabel[level] +
		fgColor + " " +
		colors.FlagReset
}
