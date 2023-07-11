package ctxparser

import (
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatArray(
	value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string,
) (result string, resultCtx *ParseResultCtx) {
	fgColor := config.ColorScheme.FgFaint
	typeStr := valueType.String()

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
	return formatArraylikeWithType(typeStr, valueStr, config), &tempResultCtx
}

func FormatSlice(
	value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string,
) (result string, resultCtx *ParseResultCtx) {
	typeStr := valueType.String()

	// Prepare resultCtx
	tempResultCtx := ParseResultCtx{
		isAllLiteral: true,
	}

	// Format values
	valueStr := ""

	// Return result
	return formatArraylikeWithType(typeStr, valueStr, config), &tempResultCtx
}

func formatArraylikeWithType(typeStr string, valueStr string, config *ParseToColoredStringConfig) string {
	fgColor := config.ColorScheme.FgFaint
	return fgColor + typeStr + ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}
