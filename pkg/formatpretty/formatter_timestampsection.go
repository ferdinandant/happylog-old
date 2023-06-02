package formatpretty

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/types"
)

func GetFormattedTimestampSection(logOpts *types.LogOpts) string {
	level := *logOpts.Level
	now := logOpts.Now
	appName := *logOpts.AppName

	// Determine color
	var fgColor string
	switch level {
	case types.LevelTrace:
		fgColor = FlagColorFgBoldBrightBlack
	case types.LevelDebug:
		fgColor = FlagColorFgBoldBlue
	case types.LevelInfo:
		fgColor = FlagColorFgBoldGreen
	case types.LevelWarn:
		fgColor = FlagColorFgBoldYellow
	case types.LevelError:
		fgColor = FlagColorFgBoldRed
	case types.LevelFatal:
		fgColor = FlagColorFgBoldMagenta
	}

	// Create string
	timestampSectionStr := formatTime(now)
	if appName != "" {
		timestampSectionStr += " [" + appName + "]"
	}

	// Return
	return fgColor + timestampSectionStr + FlagReset
}

func formatTime(now *time.Time) string {
	// Formats "Jan _2 15:04:05.000" -> "15:04:05.000"
	return now.Format(time.StampMilli)[7:]
}
