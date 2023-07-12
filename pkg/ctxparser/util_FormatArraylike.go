package ctxparser

import (
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatArray(
	traversalCtx TraversalCtx,
	// value interface{}, valueType reflect.Type, config *ParseConfig, currentDepth int, propsPath []string,
) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	fgColor := config.ColorScheme.FgFaint
	valueType := *traversalCtx.CurrentValueType

	// Prepare resultCtx
	tempResultCtx := ParseResultCtx{
		isAllLiteral: true,
	}

	// Format values
	valueStr := strings.Join([]string{
		ColorRealValue,
		"  " + colors.FormatTextWithColor(fgColor, "0:") + ColorRealValue + " dsfsdf,",
		"  " + colors.FormatTextWithColor(fgColor, "1:") + ColorRealValue + " dsfsdf,",
		"  " + colors.FormatTextWithColor(fgColor, "2:") + ColorRealValue + " dsfsdf,",
		"",
	}, "\n")

	// Return result
	// We should use `reflect.TypeOf(...).String()` so it uses the struct name
	valueTypeStr := valueType.String()
	return formatArraylikeWithType(valueTypeStr, valueStr, config), &tempResultCtx
}

func FormatSlice(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	return "Unhandled", nil
}

func formatArraylikeWithType(typeStr string, valueStr string, config *ParseConfig) string {
	fgColor := config.ColorScheme.FgFaint
	return fgColor + typeStr + ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}
