package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/types"
)

var levelToLabel = map[types.Level]string{
	10: "TRACE",
	20: "DEBUG",
	30: " INFO",
	40: " WARN",
	50: "ERROR",
	60: "FATAL",
}

func GetFormattedLabelTag(logOpts *types.FormatLogOpts) string {
	level := *logOpts.Level

	// Determine color
	var bgColor Color
	var fgColor Color
	switch level {
	case types.LevelTrace:
		fgColor = FlagColorFgBrightBlack
		bgColor = FlagColorBgBoldBrightBlack
	case types.LevelDebug:
		fgColor = FlagColorFgBlue
		bgColor = FlagColorBgBoldBlue
	case types.LevelInfo:
		fgColor = FlagColorFgGreen
		bgColor = FlagColorBgBoldGreen
	case types.LevelWarn:
		fgColor = FlagColorFgYellow
		bgColor = FlagColorBgBoldYellow
	case types.LevelError:
		fgColor = FlagColorFgRed
		bgColor = FlagColorBgBoldRed
	case types.LevelFatal:
		fgColor = FlagColorFgMagenta
		bgColor = FlagColorBgBoldMagenta
	}

	// Create string
	return string(bgColor) +
		string(fgColor) + "# " + string(FlagColorFgBlack) +
		levelToLabel[level] +
		string(fgColor) + " " +
		string(FlagReset)
}
