package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

type ParseToColoredStringConfig struct {
	// KeyFgColor is used to color struct fields
	KeyFgColor colors.Color
}

// ParseToColoredString returns the formatted/colored string of valuePtr.
// We always return string so... we force everything to be printed as much as possible
// (don't just print one single error when only one object property is failing).
func ParseToColoredString(valuePtr *interface{}, config *ParseToColoredStringConfig) string {
	if valuePtr == nil {
		return FormatParserError(fmt.Errorf("valuePtr is nil"))
	}
	return implParseToColoredString(valuePtr, config, 0, []string{})
}

func implParseToColoredString(valuePtr *interface{}, config *ParseToColoredStringConfig, depth int, propsPath []string) string {
	value := *valuePtr
	// (1) Handle nil
	if value == nil {
		result := ColorRealValue + "nil" + colors.FlagReset
		return result
	}
	// (2) Handle other cases
	// Using `reflect.TypeOf(ctx).String()` so it uses the struct name
	// https://stackoverflow.com/a/35791105/5181368
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()
	valueKindStr := strings.ToLower(valueKind.String())
	// https://pkg.go.dev/reflect#Kind
	if strings.Contains(valueKindStr, "int") {
		return ParseNumberToColoredString(value, valueKind)
	} else {
		return FormatParserError(fmt.Errorf("Unimplemented kind: %s", valueKindStr))
	}
}
