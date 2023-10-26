package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatBool(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	valuePtr := traversalCtx.CurrentValuePtr
	value := *valuePtr

	valueCast, ok := value.(bool)
	if !ok {
		err := fmt.Errorf("Cannot cast to bool: %+v", value)
		return FormatParserError(traversalCtx, err, valuePtr)
	}

	if valueCast {
		return colors.FormatTextWithColor(config.ColorMain, "true"), LiteralParseResultCtx
	} else {
		return colors.FormatTextWithColor(config.ColorMain, "false"), LiteralParseResultCtx
	}
}
