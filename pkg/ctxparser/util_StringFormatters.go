package ctxparser

import (
	"strconv"
	"strings"
)

func WrapStringWithQuotes(str string) string {
	return strconv.Quote(str)
}

func WrapStringWithBackquotes(str string) string {
	escapedValueStr := strings.ReplaceAll(str, "`", "\\`")
	return "`" + escapedValueStr + "`"
}
