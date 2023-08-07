package ctxparser

// FormatParserErrorToColoredString formats happylog's internal error in object style.
func FormatParserError(config *ParseConfig, err error, valuePtr *interface{}) string {
	return ""
	// errMsg := "<ParserErr> [Error]: " + err.Error()
	// if valuePtr != nil {
	// 	value := *valuePtr
	// 	errMsg += fmt.Sprintf(" -- [Object]: %+v", value)
	// }

	// return colors.FormatTextWithColor(ColorPlaceholderValue, errMsg)
}
