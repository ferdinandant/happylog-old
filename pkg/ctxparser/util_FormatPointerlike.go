package ctxparser

import (
	"fmt"
	"reflect"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// ================================================================================
// MAIN
// ================================================================================

func FormatUnsafePointer(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	value := *traversalCtx.CurrentValuePtr
	valueStr := fmt.Sprintf("%p", value)
	return config.ColorType + "uintptr" + config.ColorMain + "(" + valueStr + ")" + colors.FlagReset
}

func FormatPointer(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	valuePtr := traversalCtx.CurrentValuePtr
	valueType := *traversalCtx.CurrentValueType
	value := *valuePtr

	// Do not want to panic, e.g. when accessing unaddressible address
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			err := fmt.Errorf("Panic: %+v", panicErr)
			var provisionalValue interface{} = fmt.Sprintf("%+v", value)
			result = FormatParserError(traversalCtx, err, &provisionalValue)
		}
	}()

	// Get pointer type
	typeStr := valueType.String()
	addressStr := GetAddressString(value)
	tempResultCtx := ParseResultCtx{
		isAllLiteral: true,
	}
	// Parse pointed value
	var valueStr string
	if value != nil && addressStr != "nil" {
		// Get pointed value
		reflectValue := reflect.ValueOf(*valuePtr)
		pointedReflectValue := reflectValue.Elem()
		var pointedValue interface{} = pointedReflectValue.Interface()
		// Format children
		childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, SpecialTraversalDereferencingPtr, &pointedValue)
		pointedResult, pointedResultCtx := FormatAny(childrenTraversalCtx)
		if pointedResultCtx != nil && !pointedResultCtx.isAllLiteral {
			tempResultCtx.isAllLiteral = false
		}
		valueStr = pointedResult
	}

	// Return result
	return formatPointerWithType(config, typeStr, addressStr, valueStr), &tempResultCtx
}

// ================================================================================
// HELPERS
// ================================================================================

func formatPointerWithType(config *ParseConfig, typeStr string, addressStr string, valueStr string) string {
	typeSegment := config.ColorType + typeStr + " "
	addressSegment := config.ColorMain + "<" + addressStr + ">"
	if addressStr == "nil" {
		return typeSegment + addressSegment
	}
	valueSegment := (config.ColorType + " => ") + (config.ColorMain + valueStr + colors.FlagReset)
	return typeSegment + addressSegment + valueSegment
}
