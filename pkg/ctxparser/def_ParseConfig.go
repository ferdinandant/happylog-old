package ctxparser

import (
	"github.com/ferdinandant/happylog/pkg/colors"
)

// ================================================================================
// MAIN
// ================================================================================

type ParseConfig struct {
	// ------------------------------------------------------------
	// Colors
	// ------------------------------------------------------------

	// Color scheme for this error level
	ColorScheme *colors.ColorScheme

	// ColorMain is used for most purposes (for the "real values"):
	// e.g. ctx field's value, separators, type literals
	ColorMain colors.Color

	// ColorType is used to write types,
	// e.g. "[3]string" or "[]struct { a int; b int; c struct { d int } }"
	ColorType colors.Color

	// ------------------------------------------------------------
	// General Behavior
	// ------------------------------------------------------------

	// How deep should we traverse fields of objects
	MaxDepth int

	// How deep we should dereference pointer (to show its referenced value)
	MaxDereferencingDepth int

	// Allow printign array items or struct fields in one line if it only contains literals
	AllowPrintItemsInline bool

	// Whether or not to print struct methods
	PrintMethods bool

	// ------------------------------------------------------------
	// Max Items
	// ------------------------------------------------------------

	MaxFieldCount int

	MaxItemCount int
}

// ================================================================================
// HELPERS
// ================================================================================

func CreateParseConfig(colorScheme *colors.ColorScheme, overrides *ParseConfig) *ParseConfig {
	var defaultParseConfig *ParseConfig = &ParseConfig{
		// --- colors ---
		ColorScheme: colorScheme,
		ColorMain:   colors.FlagColorFgFaintBrightWhite,
		ColorType:   colors.FlagColorFgFaintBrightBlack,
		// --- general behavior ---
		MaxDepth:              5,
		MaxDereferencingDepth: 2,
		AllowPrintItemsInline: true,
		PrintMethods:          true,
		// --- max items ---
		MaxFieldCount: 50,
		MaxItemCount:  100,
	}

	return defaultParseConfig
}
