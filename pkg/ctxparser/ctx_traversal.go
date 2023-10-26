package ctxparser

import (
	"reflect"
)

// ================================================================================
// TYPES/CONST
// ================================================================================

type ParsingAsEnum int

type ParentFieldInfo struct {
	Kind   reflect.Kind
	KeyPtr *interface{}
}

const (
	// The current ctx appears in value
	ParsingAsValue ParsingAsEnum = 0
	// The current ctx appears in key (e.g. of a map)
	ParsingAsKey ParsingAsEnum = 1
)

// SpecialTraversalDereferencing just marks that we are doing deferencing operation.
// (This is just a trick to get a random memory address)
var SpecialTraversalDereferencingPtr *interface{} = new(interface{})

// ================================================================================
// MAIN
// ================================================================================

type TraversalCtx struct {
	Config    *ParseConfig
	ParsingAs ParsingAsEnum

	// ----- Parent context -----
	// Depth shows how deep we have accessed fields/indices/dereferences
	Depth int
	// IndentLevel are to be used for view only
	IndentLevel  int
	ParentFields []ParentFieldInfo

	// ----- Current node context -----
	// CurrentValuePtr can never be nil!
	CurrentValuePtr *interface{}
	// CurrentValueType and CurrentValueKind shall be nil if CurrentValuePtr is nil.
	// They are never nil otherwise.
	CurrentValueType *reflect.Type
	CurrentValueKind reflect.Kind
}

// ================================================================================
// HELPERS
// ================================================================================

func CreateTraversalCtx(config *ParseConfig, currentValuePtr *interface{}) TraversalCtx {
	// Parse value type
	currentValue := *currentValuePtr
	valueType, valueKind := getTypeAndValue(currentValue)

	// Return result
	return TraversalCtx{
		Config:           config,
		ParsingAs:        ParsingAsValue,
		Depth:            0,
		IndentLevel:      0,
		ParentFields:     []ParentFieldInfo{},
		CurrentValuePtr:  currentValuePtr,
		CurrentValueType: valueType,
		CurrentValueKind: valueKind,
	}
}

func ExtendTraversalCtx(parentTraversalCtx *TraversalCtx, childrenKeyPtr *interface{}, childrenValuePtr *interface{}) TraversalCtx {
	// Parse key type
	newParentFields := make([]ParentFieldInfo, len(parentTraversalCtx.ParentFields))
	copy(newParentFields, parentTraversalCtx.ParentFields)
	// We might traverse an object deeper, but we might not need to increase the "indent level"
	var indentLevelIncrement int
	if childrenKeyPtr == SpecialTraversalDereferencingPtr {
		indentLevelIncrement = 0
	} else {
		childrenKey := *childrenKeyPtr
		_, childrenKeyKind := getTypeAndValue(childrenKey)
		indentLevelIncrement = 1
		newParentFields = append(newParentFields, ParentFieldInfo{
			Kind:   childrenKeyKind,
			KeyPtr: childrenKeyPtr,
		})
	}

	// Parse value type
	childrenValue := *childrenValuePtr
	childrenValueType, childrenValueKind := getTypeAndValue(childrenValue)

	// Return result
	return TraversalCtx{
		Config:           parentTraversalCtx.Config,
		ParsingAs:        parentTraversalCtx.ParsingAs,
		Depth:            parentTraversalCtx.Depth + 1,
		IndentLevel:      parentTraversalCtx.IndentLevel + indentLevelIncrement,
		ParentFields:     newParentFields,
		CurrentValuePtr:  childrenValuePtr,
		CurrentValueType: childrenValueType,
		CurrentValueKind: childrenValueKind,
	}
}

func getTypeAndValue(obj interface{}) (*reflect.Type, reflect.Kind) {
	if obj == nil {
		return nil, reflect.Invalid
	}
	tempType := reflect.TypeOf(obj)
	tempKind := tempType.Kind()
	return &tempType, tempKind
}
