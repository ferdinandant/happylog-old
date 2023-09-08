package ctxparser

import (
	"github.com/ferdinandant/happylog/pkg/colors"
)

// ColorMain is used for most purposes (for the "real values"):
// e.g. ctx field's value, separators, type literals
var DefaultColorMain = colors.FlagColorFgFaintBrightWhite

// ColorType is used to write types,
// e.g. "[3]string" or "[]struct { a int; b int; c struct { d int } }"
var DefaultColorType = colors.FlagColorFgFaintBrightBlack

var DefaultMaxDepth = 5

var DefaultMaxDereferencingDepth = 3

// ================================================================================
// MAIN
// ================================================================================

type ParseConfig struct {
	// ----- Colors -----
	ColorScheme *colors.ColorScheme
	ColorMain   colors.Color
	ColorType   colors.Color

	// ----- Logging behavior -----
	MaxDepth              int
	MaxDereferencingDepth int
}
