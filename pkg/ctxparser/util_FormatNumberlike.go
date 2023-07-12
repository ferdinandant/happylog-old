package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatInteger(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	value := *traversalCtx.CurrentValuePtr
	valueKind := *traversalCtx.CurrentValueKind

	// We have used reflect already anyway.
	// So we don't have to worry if this is being slow.
	valueStr := fmt.Sprintf("%v", value)

	// Case 1: Print the number, e.g. "12"
	if valueKind == reflect.Int {
		return colors.FormatTextWithColor(ColorRealValue, valueStr)
	}
	// Case 2: Print the number with the type, e.g. "uint(12)"
	typeStr := strings.ToLower(valueKind.String())
	return formatIntegerLiteralWithType(typeStr, valueStr, config)
}

func formatIntegerLiteralWithType(typeStr string, valueStr string, config *ParseConfig) string {
	fgColor := config.ColorScheme.FgFaint
	return fgColor + typeStr + ColorRealValue + "(" + valueStr + ")" + colors.FlagReset
}
