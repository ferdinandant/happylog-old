package ctxparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func ParseNumberToColoredString(value interface{}, valueKind reflect.Kind) string {
	// The easiest way to get it done (slow).
	// We have used reflect already anyway.
	valueStr := fmt.Sprintf("%v", value)

	// Print the number, e.g. "12"
	if valueKind == reflect.Int {
		return colors.FormatTextWithColor(ColorRealValue, valueStr)
	}
	// Print the number with the type (e.g)
	typeLiteral := strings.ToLower(valueKind.String())
	return FormatNumberLiteral(valueStr, typeLiteral)
}
