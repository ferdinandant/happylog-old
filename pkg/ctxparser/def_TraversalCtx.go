package ctxparser

import "reflect"

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
	Kind *reflect.Kind
	Name interface{}
}

// ================================================================================
// MAIN
// ================================================================================

type TraversalCtx struct {
	Config    *ParseConfig
	ParsingAs ParsingAsEnum

	// ----- Parent context -----
	Depth        int
	ParentFields []ParentFieldInfo

	// ----- Current node context -----
	// CurrentValuePtr can never be nil!
	CurrentValuePtr *interface{}
	// CurrentValueType and CurrentValueKind shall be nil if CurrentValuePtr is nil.
	// They are never nil otherwise.
	CurrentValueType *reflect.Type
	CurrentValueKind *reflect.Kind
}

// ================================================================================
// HELPERS
// ================================================================================

func CreateTraversalCtx(config *ParseConfig, currentValuePtr *interface{}) TraversalCtx {
	currentValue := *currentValuePtr

	var valueType *reflect.Type
	var valueKind *reflect.Kind
	if currentValue != nil {
		tempType := reflect.TypeOf(currentValue)
		tempKind := tempType.Kind()
		valueType = &tempType
		valueKind = &tempKind
	}

	return TraversalCtx{
		Config:           config,
		ParsingAs:        ParsingAsValue,
		Depth:            0,
		ParentFields:     []ParentFieldInfo{},
		CurrentValuePtr:  currentValuePtr,
		CurrentValueType: valueType,
		CurrentValueKind: valueKind,
	}
}
