package formatpretty

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/core"
	"github.com/ferdinandant/happylog/pkg/ctxparser"
)

func GetFormattedContext(logOpts *core.FormatLogOpts) string {
	level := *logOpts.Level
	ctxPtr := logOpts.CtxPtr

	// Determine color
	var coloredFgColor core.Color
	// neutralFgColor := FlagColorFgFaintBrightBlack
	switch level {
	case core.LevelTrace:
		coloredFgColor = core.FlagColorFgFaintBrightBlack
	case core.LevelDebug:
		coloredFgColor = core.FlagColorFgFaintBlue
	case core.LevelInfo:
		coloredFgColor = core.FlagColorFgFaintGreen
	case core.LevelWarn:
		coloredFgColor = core.FlagColorFgFaintYellow
	case core.LevelError:
		coloredFgColor = core.FlagColorFgFaintRed
	case core.LevelFatal:
		coloredFgColor = core.FlagColorFgFaintMagenta
	}

	formattedCtx, err := ctxparser.ParseToColoredString(ctxPtr, &ctxparser.ParseToColoredStringConfig{
		KeyFgColor: coloredFgColor,
	})
	if err != nil {
		return fmt.Sprintf("Error parsing ctx: %s", err.Error())
	}
	return formattedCtx
}

// Create string
// Using `reflect.TypeOf(ctx).String()` so it uses the struct name
// https://stackoverflow.com/a/35791105/5181368
// ctxType := reflect.TypeOf(ctx)
// var ctxTypeName string
// var ctxTypeKind string
// if ctxType != nil {
// 	ctxTypeName = ctxType.String()
// 	ctxTypeKind = ctxType.Kind().String()
// } else {
// 	ctxTypeName = "<nil>"
// 	ctxTypeKind = "<nil>"
// }
// return coloredFgColor +
// 	fmt.Sprintf("[%+v][%+v] %+v", ctxTypeName, ctxTypeKind, ctx) + FlagReset +
// 	"\n" + FlagColorFgFaintBrightWhite + "Aku mau mencoba seperti ini?" + FlagReset +
// 	"\n" + FlagColorFgFaintBrightBlack + "Aku mau mencoba seperti ini?" + FlagReset
