package ctxparser

import "github.com/ferdinandant/happylog/pkg/colors"

func FormatParserError(err error) string {
	errMsg := "<ParserErr> " + err.Error()
	return colors.FormatTextWithColor(ColorPlaceholderValue, errMsg)
}

func FormatNumberLiteral(valueStr string, typeLiteral string) string {
	return ColorRealValue + typeLiteral + "(" + valueStr + ")" + colors.FlagReset
}
