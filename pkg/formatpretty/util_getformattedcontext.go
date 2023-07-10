package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/colors"
	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/ctxparser"
)

func GetFormattedContext(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	ctxPtr := logOpts.CtxPtr

	// Determine color
	var coloredFgColor colors.Color
	// neutralFgColor := FlagColorFgFaintBrightBlack
	switch level {
	case core.LevelTrace:
		coloredFgColor = colors.FlagColorFgFaintBrightBlack
	case core.LevelDebug:
		coloredFgColor = colors.FlagColorFgFaintBlue
	case core.LevelInfo:
		coloredFgColor = colors.FlagColorFgFaintGreen
	case core.LevelWarn:
		coloredFgColor = colors.FlagColorFgFaintYellow
	case core.LevelError:
		coloredFgColor = colors.FlagColorFgFaintRed
	case core.LevelFatal:
		coloredFgColor = colors.FlagColorFgFaintMagenta
	}

	return ctxparser.ParseToColoredString(ctxPtr, &ctxparser.ParseToColoredStringConfig{
		KeyFgColor: coloredFgColor,
	})
}
