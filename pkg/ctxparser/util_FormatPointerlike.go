package ctxparser

import (
	"fmt"
	"reflect"

	"github.com/ferdinandant/happylog/pkg/colors"
)

type PointerAddressSpec struct {
	Address string //=> "nil" or "0x123abcdef"
	Type    string //=> e.g. "**int" or "*string"
}

// ================================================================================
// MAIN
// ================================================================================

func FormatUnsafePointer(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	value := *traversalCtx.CurrentValuePtr
	valueStr := fmt.Sprintf("%p", value)
	return config.ColorType + "uintptr" + config.ColorMain + "(" + valueStr + ")" + colors.FlagReset, LiteralParseResultCtx
}

func FormatPointer(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	valuePtr := traversalCtx.CurrentValuePtr
	value := *valuePtr

	// Do not want to panic, e.g. when accessing unaddressible address
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			err := fmt.Errorf("Panic: %+v", panicErr)
			var provisionalValue interface{} = fmt.Sprintf("%+v", value)
			tmpResult, tmpResultCtx := FormatParserError(traversalCtx, err, &provisionalValue)
			result = tmpResult
			resultCtx = tmpResultCtx
		}
	}()

	// Parse pointer types
	tempResultCtx := ParseResultCtx{
		isAllDescendantLiteral: true,
	}
	isValueFound := false
	var targetValue interface{} = nil
	var addrSpecChain []PointerAddressSpec
	// Traverse pointed value via dereferencing until we get a non-pointer value
	// - https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1
	currentAddressOrValue := *valuePtr
	for dereferencingDepth := 0; dereferencingDepth <= config.MaxDereferencingDepth; dereferencingDepth++ {
		reflectValue := reflect.ValueOf(currentAddressOrValue)
		reflectKind := reflectValue.Kind()
		if currentAddressOrValue == nil {
			isValueFound = true
			targetValue = nil
			break
		}
		// Check currentAddressOrValue: is is a pointer or a value?
		// (1) If it's NOT a pointer, then we've found the referenced value
		if reflectKind != reflect.Pointer {
			isValueFound = true
			targetValue = currentAddressOrValue
			break
		}
		// (2) Otherwise, prepare for the next iteration
		currentAddressStr := GetAddressString(currentAddressOrValue)
		addrSpecChain = append(addrSpecChain, PointerAddressSpec{
			Address: GetAddressString(currentAddressOrValue),
			Type:    reflectValue.Type().String(),
		})
		if currentAddressStr == "nil" {
			break
		}
		pointedAddressReflectValue := reflectValue.Elem().Interface()
		currentAddressOrValue = pointedAddressReflectValue
	}

	// Parse pointed value
	var valueStr = ""
	if isValueFound {
		// Format children
		childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, SpecialTraversalDereferencingPtr, &targetValue)
		pointedResult, pointedResultCtx := FormatAny(childrenTraversalCtx)
		if pointedResultCtx != nil && !pointedResultCtx.isAllDescendantLiteral {
			tempResultCtx.isAllDescendantLiteral = false
		}
		valueStr = pointedResult
	}

	// Return result
	return formatPointerWithType(config, addrSpecChain, valueStr), &tempResultCtx
}

// ================================================================================
// HELPERS
// ================================================================================

func formatPointerWithType(config *ParseConfig, addrSpecChain []PointerAddressSpec, valueStr string) string {
	hasNilAddress := false
	arrowStr := (config.ColorType + " => ")

	// Print addresses
	var addressChainSegment string = ""
	for i, addrSpec := range addrSpecChain {
		currentSegmentStr := ""
		// Arrow padding
		if i > 0 {
			currentSegmentStr += arrowStr
		}
		// Type and address
		currentSegmentStr += config.ColorType + addrSpec.Type + " "
		currentSegmentStr += config.ColorMain + "<" + addrSpec.Address + ">"
		// State update
		addressChainSegment += currentSegmentStr
		if addrSpec.Address == "nil" {
			hasNilAddress = true
			break
		}
	}

	// Print value
	var valueSegment string = ""
	if !hasNilAddress && valueStr != "" {
		valueSegment = arrowStr + (config.ColorMain + valueStr)
	}
	return addressChainSegment + valueSegment + colors.FlagReset
}
