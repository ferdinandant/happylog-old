package ctxparser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// ================================================================================
// MAIN
// ================================================================================

func FormatArraylike(
	traversalCtx TraversalCtx,
) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	fgColor := config.ColorScheme.FgFaint
	valueType := *traversalCtx.CurrentValueType
	valuePtr := traversalCtx.CurrentValuePtr

	// Using slice of interface to standardize
	valueSlice, err := convertInterfaceToSlice(valuePtr)
	if err != nil {
		return FormatParserError(err, valuePtr), nil
	}
	// Iterate slice
	var itemValueStrList []string
	tempResultCtx := ParseResultCtx{
		isAllLiteral: true,
	}
	for i, itemValue := range valueSlice {
		var childKey interface{} = i
		childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, &childKey, &itemValue)
		itemResult, itemResultCtx := FormatAny(childrenTraversalCtx)
		if itemResultCtx != nil && !itemResultCtx.isAllLiteral {
			tempResultCtx.isAllLiteral = false
		}
		itemValueStrList = append(itemValueStrList, itemResult)
	}
	// Format values
	valueStrResult := ColorRealValue
	valueStrLastIdx := len(itemValueStrList) - 1
	childrenItemDepth := traversalCtx.Depth + 1
	itemFirstPrefix, itemPrefix, itemSuffix, itemLastSuffix := getItemPrefixSuffix(false, childrenItemDepth)
	for i, itemValueStr := range itemValueStrList {
		keyStr := strconv.FormatInt(int64(i), 10) + ": "
		var usedPrefix string
		var usedSuffix string
		if i == 0 {
			usedPrefix = itemFirstPrefix
			usedSuffix = itemSuffix
		} else if i == valueStrLastIdx {
			usedPrefix = itemPrefix
			usedSuffix = itemLastSuffix
		} else {
			usedPrefix = itemPrefix
			usedSuffix = itemSuffix
		}
		formattedValueStr := colors.FormatTextWithColor(fgColor, keyStr) + ColorRealValue + itemValueStr
		valueStrResult += usedPrefix + formattedValueStr + ColorRealValue + usedSuffix
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
	return formatArraylikeWithType(valueTypeStr, valueStrResult), &tempResultCtx
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

func formatArraylikeWithType(typeStr string, valueStr string) string {
	return ColorType + typeStr + ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}

func getItemPrefixSuffix(shouldPrintInOneLine bool, depth int) (
	itemFirstPrefix string, itemPrefix string, itemSuffix string, itemLastSuffix string,
) {
	if shouldPrintInOneLine {
		itemFirstPrefix = " "
		itemPrefix = " "
		itemSuffix = ", "
		itemLastSuffix = ", "
	} else {
		padding := strings.Repeat("  ", depth)
		itemFirstPrefix = "\n" + padding
		itemPrefix = padding
		itemSuffix = ",\n"
		itemLastSuffix = ",\n" + strings.Repeat("  ", depth-1)
	}
	return
}
