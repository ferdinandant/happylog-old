package ctxparser

import "github.com/ferdinandant/happylog/pkg/colors"

// FormatParserErrorToColoredString formats happylog's internal error
func FormatParserError(err error) string {
	errMsg := "<ParserErr> " + err.Error()
	return colors.FormatTextWithColor(ColorPlaceholderValue, errMsg)
}
