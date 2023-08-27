package ctxparser

import "github.com/ferdinandant/happylog/pkg/colors"

// ================================================================================
// MAIN
// ================================================================================

func FormatFunction(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	// fgColor := config.ColorScheme.FgFaint
	valueType := *traversalCtx.CurrentValueType
	valuePtr := traversalCtx.CurrentValuePtr

	// Return result
	// Currently we don't use resultCtx for functions
	typeStr := valueType.String()
	addressStr := GetAddressString(*valuePtr)
	return formatFunctionWithType(config, typeStr, addressStr), nil
}

// ================================================================================
// HELPERS
// ================================================================================

func formatFunctionWithType(config *ParseConfig, typeStr string, addressStr string) string {
	typeSegment := config.ColorType + typeStr + " "
	addressSegment := config.ColorMain + "<" + addressStr + ">"
	return typeSegment + addressSegment + colors.FlagReset
}
