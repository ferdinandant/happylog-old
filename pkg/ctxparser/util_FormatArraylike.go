package ctxparser

import (
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatArray(value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string) string {
	typeStr := valueType.String()
	valueStr := strings.Join([]string{
		"",
		"  " + colors.FormatTextWithColor(config.KeyFgColor, "0:") + " dsfsdf,",
		"  " + colors.FormatTextWithColor(config.KeyFgColor, "1:") + " dsfsdf,",
		"  " + colors.FormatTextWithColor(config.KeyFgColor, "2:") + " dsfsdf,",
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
	return config.KeyFgColor + typeStr + ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}
