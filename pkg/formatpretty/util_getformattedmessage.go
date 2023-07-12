package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/colors"
	"github.com/ferdinandant/happylog/pkg/logopts"
)

func GetFormattedMessage(logOpts *logopts.FormatLogOpts) string {
	msg := *logOpts.Msg
	colorScheme := *logOpts.ColorScheme()

	// Return string
	fgColor := colorScheme.FgNormal
	return colors.FormatTextWithColor(fgColor, msg)
}
