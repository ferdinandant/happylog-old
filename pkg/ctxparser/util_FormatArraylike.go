package ctxparser

import (
	"reflect"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatArray(value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string) string {
	return valueType.String()
}

func FormatSlice(value interface{}, valueType reflect.Type, config *ParseToColoredStringConfig, currentDepth int, propsPath []string) string {
	return valueType.String()
}

func formatArraylikeWithType(valueStr string, typeStr string, config *ParseToColoredStringConfig) string {
	return ColorPlaceholderValue + typeStr + ColorRealValue + "{" + valueStr + "}" + colors.FlagReset
}
