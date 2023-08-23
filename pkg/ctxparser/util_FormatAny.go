package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatAny(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	valuePtr := traversalCtx.CurrentValuePtr
	value := *valuePtr

	// (1) Handle nil
	if value == nil {
		result := config.ColorMain + "nil" + colors.FlagReset
		return result, nil
	}

	// (2) Handle literals
	// - https://stackoverflow.com/a/35791105/5181368
	// - https://pkg.go.dev/reflect#Kind
	valueKind := traversalCtx.CurrentValueKind
	switch valueKind {
	case reflect.Bool:
		return FormatBool(traversalCtx), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Uintptr:
		return FormatRealNumber(traversalCtx), nil
	case reflect.Complex64, reflect.Complex128:
		return FormatComplexNumber(traversalCtx), nil
	case reflect.String:
		return FormatString(traversalCtx), nil
	case reflect.UnsafePointer:
		return FormatUnsafePointer(traversalCtx), nil
	}

	// (3) Handle complex types
	// (need to determine isAllLiteral separately)
	if traversalCtx.Depth > 5 {
		return "...", nil
	}
	switch valueKind {
	case reflect.Array:
		return FormatArraylike(traversalCtx)
	case reflect.Slice:
		return FormatArraylike(traversalCtx)
	case reflect.Struct:
		return FormatStruct(traversalCtx)
	}

	// Unexpected/unhandled kind/flow
	// - https://github.com/golang/go/issues/39268
	valueKindStr := strings.ToLower(valueKind.String())
	err := fmt.Errorf("Unimplemented kind: %s", valueKindStr)
	return FormatParserError(traversalCtx, err, valuePtr), nil
}
