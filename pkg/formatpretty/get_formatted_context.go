package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/ctxparser"
	"github.com/ferdinandant/happylog/pkg/logopts"
)

func GetFormattedContext(logOpts *logopts.FormatLogOpts) string {
	ctxPtr := logOpts.CtxPtr
	colorScheme := logOpts.ColorScheme()

	// Create config
	config := ctxparser.CreateParseConfig(colorScheme, nil)

	// Return string
	return ctxparser.ParseToColoredString(config, ctxPtr)
}
