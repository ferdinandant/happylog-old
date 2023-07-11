package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatInteger(value interface{}, valueKind reflect.Kind, config *ParseToColoredStringConfig) string {
	// The easiest way to get it done (slow).
	// We have used reflect already anyway.
	valueStr := fmt.Sprintf("%v", value)

	// Print the number, e.g. "12"
	if valueKind == reflect.Int {
		return colors.FormatTextWithColor(ColorRealValue, valueStr)
	}
	// Print the number with the type (e.g)
	typeStr := strings.ToLower(valueKind.String())
	return formatIntegerLiteralWithType(typeStr, valueStr, config)
}

func formatIntegerLiteralWithType(typeStr string, valueStr string, config *ParseToColoredStringConfig) string {
	fgColor := config.ColorScheme.FgFaint
	return fgColor + typeStr + ColorRealValue + "(" + valueStr + ")" + colors.FlagReset
}
