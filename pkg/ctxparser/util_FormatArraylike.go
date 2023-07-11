package ctxparser

import (
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatArray(value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string) string {
	fgColor := config.ColorScheme.FgNormal
	typeStr := valueType.String()
	valueStr := strings.Join([]string{
		ColorRealValue,
		"  " + colors.FormatTextWithColor(fgColor, "0:") + ColorRealValue + " dsfsdf,",
		"  " + colors.FormatTextWithColor(fgColor, "1:") + ColorRealValue + " dsfsdf,",
		"  " + colors.FormatTextWithColor(fgColor, "2:") + ColorRealValue + " dsfsdf,",
		"",
	}, "\n")
	return formatArraylikeWithType(typeStr, valueStr, config)
}

func FormatSlice(value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string) string {
	typeStr := valueType.String()
	valueStr := ""
	return formatArraylikeWithType(typeStr, valueStr, config)
}

func formatArraylikeWithType(typeStr string, valueStr string, config *ParseToColoredStringConfig) string {
	fgColor := config.ColorScheme.FgFaint
	return fgColor + typeStr + ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}
