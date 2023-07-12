package ctxparser

import (
	"fmt"
)

// ParseToColoredString returns the formatted/colored string of valuePtr.
// We always return string so... we force everything to be printed as much as possible
// (don't just print one single error when only one object property is failing).
func ParseToColoredString(valuePtr *interface{}, config *ParseConfig) string {
	if valuePtr == nil {
		return FormatParserError(fmt.Errorf("valuePtr is nil"))
	}

	traversalCtx := CreateTraversalCtx(config, valuePtr)
	result, _ := FormatAny(traversalCtx)
	return result
}
