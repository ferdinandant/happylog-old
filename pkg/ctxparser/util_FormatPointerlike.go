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

func FormatUnsafePointer(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	value := *traversalCtx.CurrentValuePtr
	valueStr := fmt.Sprintf("%p", value)
	return config.ColorType + "uintptr" + config.ColorMain + "(" + valueStr + ")" + colors.FlagReset
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
			result = FormatParserError(traversalCtx, err, &provisionalValue)
		}
	}()

	// Parse pointer types
	tempResultCtx := ParseResultCtx{
		isAllLiteral: true,
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
		isPointer := reflectKind == reflect.Pointer
		if currentAddressOrValue == nil || (isPointer && reflectValue.IsNil()) {
			isValueFound = true
			targetValue = nil
			break
		}
		// Check currentAddressOrValue: is is a pointer or a value?
		// (1) If it's a pointer, then prepare next iteration and add to addrSpecChain
		// (2) Otherwise, store the value to targetValue
		if isPointer {
			addrSpecChain = append(addrSpecChain, PointerAddressSpec{
				Address: GetAddressString(currentAddressOrValue),
				Type:    reflectValue.Type().String(),
			})
			pointedAddressReflectValue := reflectValue.Elem().Interface()
			currentAddressOrValue = pointedAddressReflectValue
		} else {
			isValueFound = true
			targetValue = currentAddressOrValue
			break
		}
	}

	// Parse pointed value
	var valueStr = ""
	if isValueFound {
		// Format children
		childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, SpecialTraversalDereferencingPtr, &targetValue)
		pointedResult, pointedResultCtx := FormatAny(childrenTraversalCtx)
		if pointedResultCtx != nil && !pointedResultCtx.isAllLiteral {
			tempResultCtx.isAllLiteral = false
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
	if !hasNilAddress {
		valueSegment = arrowStr + (config.ColorMain + valueStr)
	}
	return addressChainSegment + valueSegment + colors.FlagReset
}
