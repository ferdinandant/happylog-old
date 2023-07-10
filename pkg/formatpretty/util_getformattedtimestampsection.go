package formatpretty

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/colors"
	"github.com/ferdinandant/happylog/pkg/core"
)

func GetFormattedTimestampSection(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	now := logOpts.Now
	appName := *logOpts.AppName

	// Determine color
	var fgColor colors.Color
	switch level {
	case core.LevelTrace:
		fgColor = colors.FlagColorFgBoldBrightBlack
	case core.LevelDebug:
		fgColor = colors.FlagColorFgBoldBlue
	case core.LevelInfo:
		fgColor = colors.FlagColorFgBoldGreen
	case core.LevelWarn:
		fgColor = colors.FlagColorFgBoldYellow
	case core.LevelError:
		fgColor = colors.FlagColorFgBoldRed
	case core.LevelFatal:
		fgColor = colors.FlagColorFgBoldMagenta
	}

	// Create string
	timestampSectionStr := formatTime(now)
	if appName != "" {
		timestampSectionStr += " [" + appName + "]"
	}

	// Return
	return colors.FormatTextWithColor(fgColor, timestampSectionStr)
}

func formatTime(now *time.Time) string {
	// Formats "Jan _2 15:04:05.000" -> "2015:04:05.000"
	timeStr := now.Format(time.StampMilli)[7:]
	return "[" + timeStr + "]"
}
