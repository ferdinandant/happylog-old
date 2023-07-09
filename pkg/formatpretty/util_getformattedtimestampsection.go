package formatpretty

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/core"
)

func GetFormattedTimestampSection(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	now := logOpts.Now
	appName := *logOpts.AppName

	// Determine color
	var fgColor core.Color
	switch level {
	case core.LevelTrace:
		fgColor = core.FlagColorFgBoldBrightBlack
	case core.LevelDebug:
		fgColor = core.FlagColorFgBoldBlue
	case core.LevelInfo:
		fgColor = core.FlagColorFgBoldGreen
	case core.LevelWarn:
		fgColor = core.FlagColorFgBoldYellow
	case core.LevelError:
		fgColor = core.FlagColorFgBoldRed
	case core.LevelFatal:
		fgColor = core.FlagColorFgBoldMagenta
	}

	// Create string
	timestampSectionStr := formatTime(now)
	if appName != "" {
		timestampSectionStr += " [" + appName + "]"
	}

	// Return
	return FormatTextWithColor(fgColor, timestampSectionStr)
}

func formatTime(now *time.Time) string {
	// Formats "Jan _2 15:04:05.000" -> "2015:04:05.000"
	timeStr := now.Format(time.StampMilli)[7:]
	return "[" + timeStr + "]"
}
