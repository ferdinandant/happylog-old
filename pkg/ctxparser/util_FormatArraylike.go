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
	valueKind := *traversalCtx.CurrentValueKind
	valuePtr := traversalCtx.CurrentValuePtr

	// Format values
	// Using slice of interface to standardize
	valueSlice, err := convertInterfaceToSlice(valuePtr, valueKind, valueType)
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
func convertInterfaceToSlice(valuePtr *interface{}, k reflect.Kind, t reflect.Type) ([]interface{}, error) {
	if k == reflect.Slice {
		return nil, fmt.Errorf("Unimplemented")
	} else if k == reflect.Array {
		arrayLength := t.Len()
		resultSlice := make([]interface{}, arrayLength)
		reflectValue := reflect.ValueOf(*valuePtr)
		for i := 0; i < arrayLength; i++ {
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
