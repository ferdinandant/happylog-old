package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/levels"
)

var levelToLabel = map[levels.Level]string{
	10: " TRACE ",
	20: " DEBUG ",
	30: " INFO  ",
	40: " WARN  ",
	50: " ERROR ",
	60: " FATAL ",
}

func GetFormattedLabelTag(level levels.Level) string {
	var bgColor string
	var fgColor string
	switch level {
	case levels.Trace:
		fgColor = FlagColorFgBrightBlack
		bgColor = FlagColorBgBoldBrightBlack
	case levels.Debug:
		fgColor = FlagColorFgBlue
		bgColor = FlagColorBgBoldBlue
	case levels.Info:
		fgColor = FlagColorFgGreen
		bgColor = FlagColorBgBoldGreen
	case levels.Warn:
		fgColor = FlagColorFgYellow
		bgColor = FlagColorBgBoldYellow
	case levels.Error:
		fgColor = FlagColorFgRed
		bgColor = FlagColorBgBoldRed
	case levels.Fatal:
		fgColor = FlagColorFgMagenta
		bgColor = FlagColorBgBoldMagenta
	}

	return bgColor +
		fgColor + "[" + FlagColorFgBlack +
		levelToLabel[level] +
		fgColor + "]" +
		FlagReset
}
