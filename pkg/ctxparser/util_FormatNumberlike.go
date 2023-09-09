package ctxparser

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// ================================================================================
// MAIN
// ================================================================================

func FormatRealNumber(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	valueKind := traversalCtx.CurrentValueKind
	value := *traversalCtx.CurrentValuePtr

	// We have used reflect already anyway.
	// So we don't have to worry if this is being slow.
	valueStr := fmt.Sprintf("%v", value)

	// Case 1: Print the number, e.g. "12"
	if valueKind == reflect.Int {
		return colors.FormatTextWithColor(config.ColorMain, valueStr), LiteralParseResultCtx
	}
	// Case 2: Print the number with the type, e.g. "uint(12)"
	typeStr := strings.ToLower(valueKind.String())
	return formatNumberLiteralWithType(config, typeStr, valueStr), LiteralParseResultCtx
}

func FormatComplexNumber(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	valueKind := traversalCtx.CurrentValueKind
	value := *traversalCtx.CurrentValuePtr

	// Parse complex number
	var valueStr string
	var castOk bool = false
	typeStr := strings.ToLower(valueKind.String())
	if valueKind == reflect.Complex64 {
		valueCast, ok := value.(complex64)
		if ok {
			valueStr = formatComplexNumber(complex128(valueCast))
			castOk = true
		}
	} else {
		valueCast, ok := value.(complex128)
		valueStr = fmt.Sprintf("%v %+vi", real(valueCast), imag(valueCast))
		if ok {
			valueStr = formatComplexNumber(valueCast)
			castOk = true
		}
	}
	if !castOk {
		err := fmt.Errorf("Cannot cast to %s: %+v", typeStr, value)
		return FormatParserError(traversalCtx, err, traversalCtx.CurrentValuePtr)
	}

	// We have used reflect already anyway.
	// So we don't have to worry if this is being slow.
	return formatNumberLiteralWithType(config, typeStr, valueStr), LiteralParseResultCtx
}

// ================================================================================
// HELPERS
// ================================================================================

func formatNumberLiteralWithType(config *ParseConfig, typeStr string, valueStr string) string {
	return config.ColorType + typeStr + config.ColorMain + "(" + valueStr + ")" + colors.FlagReset
}

// As time of writing, using generics for complex number did not work.
// "(variable of type T constrained by complex64 | complex128) not supported as argument to imag for go1.18 (see issue #50937)"
func formatComplexNumber(cplx complex128) string {
	realPart := real(cplx)
	imagPart := imag(cplx)
	return fmt.Sprintf("%v %s %+vi", realPart, getSign(imagPart), math.Abs(imagPart))
}

func getSign(number float64) string {
	if number < 0 {
		return "-"
	}
	return "+"
}
