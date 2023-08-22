package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatString(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	value := *traversalCtx.CurrentValuePtr
	shouldUseBacktick := traversalCtx.Depth == 0

	// Cast string value
	unescapedValueStr, ok := value.(string)
	if !ok {
		err := fmt.Errorf("Cannot cast to string: %+v", value)
		return FormatParserError(traversalCtx, err, traversalCtx.CurrentValuePtr)
	}

	// Here we just want to print the string as-is, just escape the "`"
	// (Used when `depth == 0`, just to make strings more readable)
	if shouldUseBacktick {
		quotedEscapedValueStr := WrapStringWithBackquotes(unescapedValueStr)
		return colors.FormatTextWithColor(config.ColorMain, quotedEscapedValueStr)
	}
	// Here we replace characters like '<newline>' and '<quote>' to "\n" and "\"".
	// (Used when `depth > 1`, because it means this is a key to something)
	quotedEscapedValueStr := WrapStringWithQuotes(unescapedValueStr)
	return colors.FormatTextWithColor(config.ColorMain, quotedEscapedValueStr)
}
