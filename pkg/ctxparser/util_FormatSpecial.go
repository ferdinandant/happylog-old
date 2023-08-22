package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// FormatParserErrorToColoredString formats happylog's internal error in object style.
func FormatParserError(traversalCtx TraversalCtx, err error, valuePtr *interface{}) string {
	config := traversalCtx.Config

	typeStr := "<ParserErr>"
	valueStr := ""
	if valuePtr != nil {
		value := *valuePtr
		valueStr = fmt.Sprintf("%+v", value)
	}

	// childrenItemDepth := traversalCtx.Depth + 1
	// itemPsGenerator, err := CreateItemPrefixSuffixGenerator(false, childrenItemDepth, 2)
	return config.ColorType + typeStr + config.ColorMain + " {" + valueStr + "}" + colors.FlagReset
}
