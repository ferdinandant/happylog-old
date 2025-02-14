package ctxparser

import (
	"fmt"
)

// ParseToColoredString returns the formatted/colored string of valuePtr.
// We always return string so... we force everything to be printed as much as possible
// (don't just print one single error when only one object property is failing).
func ParseToColoredString(config *ParseConfig, valuePtr *interface{}) string {
	// Create traversalCtx
	traversalCtx := CreateTraversalCtx(config, valuePtr)
	if valuePtr == nil {
		err := fmt.Errorf("valuePtr is nil")
		resultStr, _ := FormatParserError(traversalCtx, err, valuePtr)
		return resultStr
	}

	// Generate result for valid values
	result, _ := FormatAny(traversalCtx)
	return result
}
