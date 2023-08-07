package ctxparser

import (
	"fmt"
	"reflect"
	"strconv"

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
	valueStrResult := config.ColorRealValue
	childrenItemDepth := traversalCtx.Depth + 1
	childrenCount := len(itemValueStrList)
	itemPsGenerator, err := CreateItemPrefixSuffixGenerator(false, childrenItemDepth, childrenCount)
	if err != nil {
		return FormatParserError(err, valuePtr), nil
	}
	for i, itemValueStr := range itemValueStrList {
		keyStr := strconv.FormatInt(int64(i), 10) + ": "
		usedPrefix, usedSuffix := itemPsGenerator.GetPrefixSuffix(i)
		formattedValueStr := colors.FormatTextWithColor(fgColor, keyStr) + config.ColorRealValue + itemValueStr
		valueStrResult += usedPrefix + formattedValueStr + config.ColorRealValue + usedSuffix
	}

	// Return result
	// We should use `reflect.TypeOf(...).String()` so it uses the struct name
	valueTypeStr := valueType.String()
	return formatArraylikeWithType(valueTypeStr, valueStrResult, config), &tempResultCtx
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
	return config.ColorType + typeStr + config.ColorRealValue + " {" + valueStr + "}" + colors.FlagReset
}
