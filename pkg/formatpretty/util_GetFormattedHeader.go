package formatpretty

import (
	"time"

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

func GetFormattedHeader(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	now := logOpts.Now
	appName := *logOpts.AppName

	// Determine color
	var bgColor colors.Color
	var fgThemeColor colors.Color
	var fgThemeBoldColor colors.Color
	fgBlackColor := colors.FlagColorFgBlack
	switch level {
	case core.LevelTrace:
		fgThemeColor = colors.FlagColorFgBrightBlack
		fgThemeBoldColor = colors.FlagColorFgBoldBrightBlack
		bgColor = colors.FlagColorBgBoldBrightBlack
	case core.LevelDebug:
		fgThemeColor = colors.FlagColorFgBlue
		fgThemeBoldColor = colors.FlagColorFgBoldBlue
		bgColor = colors.FlagColorBgBoldBlue
	case core.LevelInfo:
		fgThemeColor = colors.FlagColorFgGreen
		fgThemeBoldColor = colors.FlagColorFgBoldGreen
		bgColor = colors.FlagColorBgBoldGreen
	case core.LevelWarn:
		fgThemeColor = colors.FlagColorFgYellow
		fgThemeBoldColor = colors.FlagColorFgBoldYellow
		bgColor = colors.FlagColorBgBoldYellow
	case core.LevelError:
		fgThemeColor = colors.FlagColorFgRed
		fgThemeBoldColor = colors.FlagColorFgBoldRed
		bgColor = colors.FlagColorBgBoldRed
	case core.LevelFatal:
		fgThemeColor = colors.FlagColorFgMagenta
		fgThemeBoldColor = colors.FlagColorFgBoldMagenta
		bgColor = colors.FlagColorBgBoldMagenta
	}

	// Create string
	ribbonText := " " + levelToLabel[level] + " | " + formatTime(now) + " "
	formattedHeader := bgColor + (fgThemeColor + "|") + (fgBlackColor + ribbonText) + (fgThemeColor + "|") + colors.FlagReset
	if appName != "" {
		tagsText := " [" + appName + "]"
		formattedHeader += fgThemeBoldColor + tagsText + colors.FlagReset
	}
	return formattedHeader
}

func formatTime(now *time.Time) string {
	// Formats "Jan _2 15:04:05.000" -> "2015:04:05.000"
	return now.Format(time.StampMilli)[7:]
}
