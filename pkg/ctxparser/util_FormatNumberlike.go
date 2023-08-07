package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatInteger(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	valueKind := traversalCtx.CurrentValueKind
	value := *traversalCtx.CurrentValuePtr

	// We have used reflect already anyway.
	// So we don't have to worry if this is being slow.
	valueStr := fmt.Sprintf("%v", value)

	// Case 1: Print the number, e.g. "12"
	if valueKind == reflect.Int {
		return colors.FormatTextWithColor(config.ColorMain, valueStr)
	}
	// Case 2: Print the number with the type, e.g. "uint(12)"
	typeStr := strings.ToLower(valueKind.String())
	return formatIntegerLiteralWithType(config, typeStr, valueStr)
}

func formatIntegerLiteralWithType(config *ParseConfig, typeStr string, valueStr string) string {
	return config.ColorType + typeStr + config.ColorMain + "(" + valueStr + ")" + colors.FlagReset
}
