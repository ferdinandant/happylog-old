package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// FormatParserErrorToColoredString formats happylog's internal error
func FormatParserError(err error, valuePtr *interface{}) string {
	errMsg := "<ParserErr> [Error]: " + err.Error()
	if valuePtr != nil {
		value := *valuePtr
		errMsg += fmt.Sprintf(" -- [Object]: %+v", value)
	}

	return colors.FormatTextWithColor(ColorPlaceholderValue, errMsg)
}
