package ctxparser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatString(value interface{}, shouldEscape bool) string {
	unescapedValueStr, ok := value.(string)
	if !ok {
		err := fmt.Errorf("Cannot cast to string: %+v", value)
		return FormatParserError(err)
	}
	// Here we just want to print the string as-is, just escape the "`"
	// (Used when `depth == 0`)
	if !shouldEscape {
		escapedValueStr := strings.ReplaceAll(unescapedValueStr, "`", "\\`")
		formattedValueStr := "`" + escapedValueStr + "`"
		return colors.FormatTextWithColor(ColorRealValue, formattedValueStr)
	}
	// Here we replace characters like '<newline>' and '<quote>' to "\n" and "\"".
	// (Used when `depth > 1`, because it means this is a key to something)
	escapedValueStr := strconv.Quote(unescapedValueStr)
	return colors.FormatTextWithColor(ColorRealValue, escapedValueStr)
}