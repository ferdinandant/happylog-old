package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

// FormatParserErrorToColoredString formats happylog's internal error in object style.
func FormatParserError(traversalCtx TraversalCtx, err error, valuePtr *interface{}) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	fgColor := config.ColorScheme.FgFaint

	// Parse value
	// Normally: we want to at least log the error message
	parserErrorPropertyPairs := [][2]string{
		{"ErrorMsg", WrapStringWithBackquotes(err.Error())},
	}
	// Otherwise: we want to log error message and values
	if valuePtr != nil {
		value := *valuePtr
		valueStr := fmt.Sprintf("%+v", value)
		quotedEscapedValueStr := WrapStringWithBackquotes(valueStr)
		valueCtxPropertyPair := [2]string{"ValueStr", quotedEscapedValueStr}
		parserErrorPropertyPairs = append(parserErrorPropertyPairs, valueCtxPropertyPair)
	}

	// Format values
	valueStrResult := config.ColorMain
	childrenIndentLevel := traversalCtx.IndentLevel + 1
	itemPsGenerator := MustCreateItemPrefixSuffixGenerator(
		false, childrenIndentLevel, len(parserErrorPropertyPairs), false,
	)
	// Print key name/value mappings
	for propertyIdx, parserErrorPropertiesPair := range parserErrorPropertyPairs {
		usedPrefix, usedSuffix := itemPsGenerator.GetPrefixSuffix(propertyIdx)
		propertyKeyStr := parserErrorPropertiesPair[0] + ": "
		propertyValueStr := parserErrorPropertiesPair[1]
		formattedValueStr := colors.FormatTextWithColor(fgColor, propertyKeyStr) + config.ColorMain + propertyValueStr
		valueStrResult += usedPrefix + formattedValueStr + config.ColorMain + usedSuffix
	}

	// Return result
	typeStr := "!!!ERROR!!!"
	resultStr := config.ColorType + typeStr + config.ColorMain + " {" + valueStrResult + "}" + colors.FlagReset
	return resultStr, ErrorParseResultCtx
}
