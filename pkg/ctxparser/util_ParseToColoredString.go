package ctxparser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/ferdinandant/happylog/pkg/core"
)

// brightFgColor is used to color values
const brightFgColor = core.FlagColorFgFaintBrightWhite

// darkFgColor is used to color separators, parantheses, type, etc.
const darkFgColor = core.FlagColorFgFaintBrightBlack

type ParseToColoredStringConfig struct {
	// KeyFgColor is used to color struct fields
	KeyFgColor core.Color
}

// ================================================================================
// MAIN
// ================================================================================

func ParseToColoredString(valuePtr *interface{}, config *ParseToColoredStringConfig) (string, error) {
	if valuePtr == nil {
		return "", fmt.Errorf("valuePtr is nil")
	}
	return implParseToColoredString(valuePtr, config, 0, []string{})
}

func implParseToColoredString(valuePtr *interface{}, config *ParseToColoredStringConfig, depth int, propsPath []string) (string, error) {
	value := *valuePtr
	// (1) Handle nil
	if value == nil {
		result := brightFgColor + "nil" + core.FlagReset
		return result, nil
	}
	// (2) Handle other cases
	// Using `reflect.TypeOf(ctx).String()` so it uses the struct name
	// https://stackoverflow.com/a/35791105/5181368
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()
	valueKindStr := valueKind.String()
	// https://pkg.go.dev/reflect#Kind
	if strings.Contains(valueKindStr, "Int") {
		return formatNumberToColoredString(value, valueKind)
	} else {
		return "", fmt.Errorf("Unhandled kind: %s", valueKind)
	}

	// result = "TODO//" + valueType.String() + "//" + valueKind.String()
	// return &result, nil
}

// ================================================================================
// TYPE PARSERS (formatXxxToColoredString)
// ================================================================================

func formatNumberToColoredString(value interface{}, valueKind reflect.Kind, propsPath []string) (string, error) {
	var result string
	valueCasted, ok := value.(int64)
	if !ok {
		propsPathStr := getPropsPathStr(propsPath)
		return "", fmt.Errorf("Cannot cast %v to int64 (at %s)", value, propsPathStr)
	}
	if valueKind == reflect.Int {

		valueStr := strconv.FormatInt(valueCasted, 10)
		result = brightFgColor + valueStr + core.FlagReset
	}
	return result, nil
}

// ================================================================================
// HELPERS
// ================================================================================

func getPropsPathStr(propsPath []string) string {

}
