package ctxparser

import (
	"fmt"
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

// GetAddressString returns address as in "0x123456" or "nil"
func GetAddressString(value interface{}) string {
	addressStr := fmt.Sprintf("%p", value)
	if addressStr == "0x0" {
		return "nil"
	}
	return addressStr
}
