package ctxparser

import (
	"fmt"
	"reflect"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// ================================================================================
// MAIN
// ================================================================================

func FormatArraylike(
	traversalCtx TraversalCtx,
) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	// fgColor := config.ColorScheme.FgFaint
	valueType := *traversalCtx.CurrentValueType
	valuePtr := traversalCtx.CurrentValuePtr

	// Format values
	// Using slice of interface to standardize
	valueSlice, err := convertInterfaceToSlice(valuePtr)
	if err != nil {
		return FormatParserError(err, valuePtr), nil
	}
	tempResultCtx := ParseResultCtx{
		isAllLiteral: true,
	}
	valueStr := ""
	// Iterate slice
	for i, itemValue := range valueSlice {
		println(i, itemValue)
	}
	// valueStr := strings.Join([]string{
	// 	ColorRealValue,
	// 	"  " + colors.FormatTextWithColor(fgColor, "0:") + ColorRealValue + " dsfsdf,",
	// 	"  " + colors.FormatTextWithColor(fgColor, "1:") + ColorRealValue + " dsfsdf,",
	// 	"  " + colors.FormatTextWithColor(fgColor, "2:") + ColorRealValue + " dsfsdf,",
	// 	"",
	// }, "\n")

	// Return result
	// We should use `reflect.TypeOf(...).String()` so it uses the struct name
	valueTypeStr := valueType.String()
	return formatArraylikeWithType(valueTypeStr, valueStr, config), &tempResultCtx
}

// ================================================================================
// HELPERS
// ================================================================================

// Using valuePtr so we are not copying the array's values to here.
func convertInterfaceToSlice(valuePtr *interface{}) ([]interface{}, error) {
	reflectValue := reflect.ValueOf(*valuePtr)
	kind := reflectValue.Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		// ChatGPT: "go how to convert interface of array to []interface{}"
		// ChatGPT: "what if you don't know the type/length of the array in compile-time?"
		resultSlice := make([]interface{}, reflectValue.Len())
		for i := 0; i < reflectValue.Len(); i++ {
			resultSlice[i] = reflectValue.Index(i).Interface()
		}
		return resultSlice, nil
	}
	return nil, fmt.Errorf("Not an array-like")
}

func formatArraylikeWithType(typeStr string, valueStr string, config *ParseConfig) string {
	fgColor := config.ColorScheme.FgFaint
	return fgColor + typeStr + ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}
