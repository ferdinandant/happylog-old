package formatpretty

import (
	"fmt"
	"reflect"

	"github.com/ferdinandant/happylog/pkg/types"
)

func GetFormattedContext(logOpts *types.FormatLogOpts) string {
	level := *logOpts.Level
	ctx := *logOpts.CtxPtr

	// Determine color
	var coloredFgColor Color
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
	// Using `reflect.TypeOf(ctx).String()` so it uses the struct name
	// https://stackoverflow.com/a/35791105/5181368
	ctxType := reflect.TypeOf(ctx)
	var ctxTypeName string
	var ctxTypeKind string
	if ctxType != nil {
		ctxTypeName = ctxType.String()
		ctxTypeKind = ctxType.Kind().String()
	} else {
		ctxTypeName = "<nil>"
		ctxTypeKind = "<nil>"
	}
	return coloredFgColor +
		fmt.Sprintf("[%+v][%+v] %+v", ctxTypeName, ctxTypeKind, ctx) + FlagReset +
		"\n" + FlagColorFgFaintBrightWhite + "Aku mau mencoba seperti ini?" + FlagReset +
		"\n" + FlagColorFgFaintBrightBlack + "Aku mau mencoba seperti ini?" + FlagReset
}
