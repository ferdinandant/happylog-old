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
		return FormatParserError(traversalCtx, err, valuePtr)
	}

	// Iterate slice indices
	// itemValueStrList contains formatting result per index
	isAllItemLiteral := true
	iteratedItemCount := 0
	var itemValueStrList []string
	for i, itemValue := range valueSlice {
		iteratedItemCount++
		if iteratedItemCount > config.MaxItemCount {
			break
		}
		// Process item
		var itemKey interface{} = i
		childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, &itemKey, &itemValue)
		itemResult, itemResultCtx := FormatAny(childrenTraversalCtx)
		if itemResultCtx != nil && !itemResultCtx.isLiteral {
			isAllItemLiteral = false
		}
		itemValueStrList = append(itemValueStrList, itemResult)
	}

	// Combine result
	valueStrResult := config.ColorMain
	childrenIndentLevel := traversalCtx.IndentLevel + 1
	itemCount := len(valueSlice)
	hasOmittedItems := itemCount > config.MaxItemCount
	shouldPrintInline := CheckShouldPrintInline(config, traversalCtx.Depth, isAllItemLiteral)
	itemPsGenerator, err := CreateItemPrefixSuffixGenerator(shouldPrintInline, childrenIndentLevel, itemCount, hasOmittedItems)
	if err != nil {
		return FormatParserError(traversalCtx, err, valuePtr)
	}
	// Print the fields
	for i, itemValueStr := range itemValueStrList {
		keyStr := strconv.FormatInt(int64(i), 10) + ": "
		usedPrefix, usedSuffix := itemPsGenerator.GetPrefixSuffix(i)
		formattedValueStr := colors.FormatTextWithColor(fgColor, keyStr) + config.ColorMain + itemValueStr
		valueStrResult += usedPrefix + formattedValueStr + config.ColorMain + usedSuffix
	}
	// Print the ellipsis
	if hasOmittedItems {
		usedPrefix, usedSuffix := itemPsGenerator.GetPrefixSuffix(itemCount)
		numberOfHiddenFields := itemCount - config.MaxFieldCount
		formattedValueStr := "... " + strconv.Itoa(numberOfHiddenFields) + " hidden field(s)"
		valueStrResult += usedPrefix + formattedValueStr + config.ColorMain + usedSuffix
	}

	// Return result
	// We should use `reflect.TypeOf(...).String()` so it uses the struct name
	valueTypeStr := valueType.String()
	return formatArraylikeWithType(config, valueTypeStr, valueStrResult), StructParseResultCtx
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

func formatArraylikeWithType(config *ParseConfig, typeStr string, valueStr string) string {
	return config.ColorType + typeStr + config.ColorMain + " {" + valueStr + "}" + colors.FlagReset
}
