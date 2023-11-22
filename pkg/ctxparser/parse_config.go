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

	// Allow printing array items or struct fields in one line if it only contains literals
	AllowPrintItemsInline bool

	// Whether or not to print EXPORTED struct methods and receiver functions.
	// (Exploring unexported methods look tricky: http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/)
	PrintPublicMethods bool

	// ------------------------------------------------------------
	// Max Items
	// ------------------------------------------------------------

	// Maximum number of fields/methods to show.
	MaxFieldCount int

	// Maximum number of array/slice items to show
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
		PrintPublicMethods:    true,
		// --- max items ---
		MaxFieldCount: 100,
		MaxItemCount:  100,
	}

	return defaultParseConfig
}

func CheckShouldPrintInline(config *ParseConfig, depth int, isAllLiteral bool) bool {
	// On `depth == 0`, just print in separate lines because we want to direct the reading flow downwards
	return config.AllowPrintItemsInline && depth > 1 && isAllLiteral
}
