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

	// (3) Handle other cases
	// - https://stackoverflow.com/a/35791105/5181368
	// - https://pkg.go.dev/reflect#Kind
	valueKind := traversalCtx.CurrentValueKind
	switch valueKind {
	// --- These are all literals ---
	case reflect.Bool:
		return FormatBool(traversalCtx), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return FormatInteger(traversalCtx), nil
	case reflect.String:
		return FormatString(traversalCtx), nil
	// --- These are all complex types ---
	// (need to determine isAllLiteral separately)
	case reflect.Array:
		return FormatArraylike(traversalCtx)
	case reflect.Slice:
		return FormatArraylike(traversalCtx)
	}

	// Unexpected/unhandled kind/flow
	// https://github.com/golang/go/issues/39268
	valueKindStr := strings.ToLower(valueKind.String())
	err := fmt.Errorf("Unimplemented kind: %s", valueKindStr)
	return FormatParserError(traversalCtx, err, valuePtr), nil
}
