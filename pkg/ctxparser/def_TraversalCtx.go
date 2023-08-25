package ctxparser

import (
	"reflect"
)

// ================================================================================
// TYPES/CONST
// ================================================================================

type ParsingAsEnum int

const (
	// The current ctx appears in value
	ParsingAsValue ParsingAsEnum = 0
	// The current ctx appears in key (e.g. of a map)
	ParsingAsKey ParsingAsEnum = 1
)

type ParentFieldInfo struct {
	Kind   reflect.Kind
	KeyPtr *interface{}
}

// ================================================================================
// MAIN
// ================================================================================

type TraversalCtx struct {
	Config    *ParseConfig
	ParsingAs ParsingAsEnum

	// ----- Parent context -----
	Depth        int
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
	currentValue := *currentValuePtr
	// Parse value type
	var valueType *reflect.Type
	var valueKind reflect.Kind
	if currentValue != nil {
		tempType := reflect.TypeOf(currentValue)
		tempKind := tempType.Kind()
		valueType = &tempType
		valueKind = tempKind
	}
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
	newParentFields := append(parentTraversalCtx.ParentFields, ParentFieldInfo{})
	// Parse value type
	var childrenValueType *reflect.Type = nil
	var childrenValueKind reflect.Kind = reflect.Invalid
	childrenValue := *childrenValuePtr
	if childrenValue != nil {
		tempType := reflect.TypeOf(childrenValue)
		tempKind := tempType.Kind()
		childrenValueType = &tempType
		childrenValueKind = tempKind
	}
	// println(childrenValueKind)
	// Return result
	return TraversalCtx{
		Config:           parentTraversalCtx.Config,
		ParsingAs:        parentTraversalCtx.ParsingAs,
		Depth:            parentTraversalCtx.Depth + 1,
		IndentLevel:      parentTraversalCtx.IndentLevel + 1,
		ParentFields:     newParentFields,
		CurrentValuePtr:  childrenValuePtr,
		CurrentValueType: childrenValueType,
		CurrentValueKind: childrenValueKind,
	}
}
