package formatpretty

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/types"
)

func GetFormattedContext(logOpts *types.FormatLogOpts) string {
	level := *logOpts.Level
	ctx := *logOpts.CtxPtr

	// Determine color
	var coloredFgColor string
	// neutralFgColor := FlagColorFgFaintBrightBlack
	switch level {
	case types.LevelTrace:
		coloredFgColor = FlagColorFgFaintBrightBlack
	case types.LevelDebug:
		coloredFgColor = FlagColorFgFaintBlue
	case types.LevelInfo:
		coloredFgColor = FlagColorFgFaintGreen
	case types.LevelWarn:
		coloredFgColor = FlagColorFgFaintYellow
	case types.LevelError:
		coloredFgColor = FlagColorFgFaintRed
	case types.LevelFatal:
		coloredFgColor = FlagColorFgFaintMagenta
	}

	// Create string
	return coloredFgColor + fmt.Sprintf("%+v", ctx) + FlagReset
}
