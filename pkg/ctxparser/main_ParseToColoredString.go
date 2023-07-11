package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

type ParseToColoredStringConfig struct {
	// KeyFgColor is used to color struct fields
	ColorScheme *colors.ColorScheme
}

// ParseToColoredString returns the formatted/colored string of valuePtr.
// We always return string so... we force everything to be printed as much as possible
// (don't just print one single error when only one object property is failing).
func ParseToColoredString(valuePtr *interface{}, config *ParseToColoredStringConfig) string {
	if valuePtr == nil {
		return FormatParserError(fmt.Errorf("valuePtr is nil"))
	}
	result, _ := implParseToColoredString(valuePtr, config, 0, []string{})
	return result
}

func implParseToColoredString(
	valuePtr *interface{}, config *ParseToColoredStringConfig, depth int, propsPath []string,
) (result string, resultCtx *ParseResultCtx) {
	value := *valuePtr

	// (1) Handle nil
	if value == nil {
		result := ColorRealValue + "nil" + colors.FlagReset
		return result, nil
	}

	// (3) Handle other cases
	// We should use `reflect.TypeOf(ctx).String()` so it uses the struct name
	// https://stackoverflow.com/a/35791105/5181368
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()
	// See the different types of `valueType.Kind()` here:
	// https://pkg.go.dev/reflect#Kind
	switch valueKind {
	// --- These are all literals ---
	case reflect.Bool:
		return FormatBool(value), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return FormatInteger(value, valueKind, config), nil
	case reflect.String:
		shouldEscape := depth > 0
		return FormatString(value, shouldEscape), nil
	// --- These are all complex types ---
	// (need to determine isAllLiteral separately)
	case reflect.Array:
		return FormatArray(value, valueType, config, depth, propsPath)
	case reflect.Slice:
		return FormatSlice(value, valueType, config, depth, propsPath)
	}

	// Unexpected/unhandled kind/flow
	// https://github.com/golang/go/issues/39268
	valueKindStr := strings.ToLower(valueKind.String())
	return FormatParserError(fmt.Errorf("Unimplemented kind: %s", valueKindStr)), nil
}
