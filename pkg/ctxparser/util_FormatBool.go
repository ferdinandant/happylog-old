package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatBool(traversalCtx TraversalCtx) string {
	valuePtr := traversalCtx.CurrentValuePtr
	value := *valuePtr

	valueCast, ok := value.(bool)
	if !ok {
		err := fmt.Errorf("Cannot cast to bool: %+v", value)
		return FormatParserError(err, valuePtr)
	}

	if valueCast {
		return colors.FormatTextWithColor(ColorRealValue, "true")
	} else {
		return colors.FormatTextWithColor(ColorRealValue, "false")
	}
}
