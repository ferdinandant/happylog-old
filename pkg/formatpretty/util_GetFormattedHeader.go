package formatpretty

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/colors"
	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/logopts"
)

var levelToLabel = map[core.Level]string{
	10: "TRACE",
	20: "DEBUG",
	30: " INFO",
	40: " WARN",
	50: "ERROR",
	60: "FATAL",
}

func GetFormattedHeader(logOpts *logopts.FormatLogOpts) string {
	now := logOpts.Now
	level := *logOpts.Level
	appName := *logOpts.AppName
	colorScheme := *logOpts.ColorScheme()

	// Determine color
	bgThemeBoldColor := colorScheme.BgBold
	fgThemeBoldColor := colorScheme.FgBold
	fgBlackBoldColor := colors.FlagColorFgFollowBlack

	// Return string
	// e.g. "| ERROR | 04:19:34.552 | [webacd-desktop]""
	ribbonText := " " + levelToLabel[level] + " | " + formatTime(now) + " "
	resultStr := bgThemeBoldColor + (fgThemeBoldColor + "|") + (fgBlackBoldColor + ribbonText) + (fgThemeBoldColor + "|") + colors.FlagReset
	if appName != "" {
		tagsText := " [" + appName + "]"
		resultStr += fgThemeBoldColor + tagsText + colors.FlagReset
	}
	return resultStr
}

func formatTime(now *time.Time) string {
	// Formats "Jan _2 15:04:05.000" -> "2015:04:05.000"
	return now.Format(time.StampMilli)[7:]
}
