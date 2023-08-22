package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatBool(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	valuePtr := traversalCtx.CurrentValuePtr
	value := *valuePtr

	valueCast, ok := value.(bool)
	if !ok {
		err := fmt.Errorf("Cannot cast to bool: %+v", value)
		return FormatParserError(traversalCtx, err, valuePtr)
	}

	if valueCast {
		return colors.FormatTextWithColor(config.ColorMain, "true")
	} else {
		return colors.FormatTextWithColor(config.ColorMain, "false")
	}
}
