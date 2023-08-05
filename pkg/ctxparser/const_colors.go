package ctxparser

import "github.com/ferdinandant/happylog/pkg/colors"

// ColorRealValue is used for most purposes (for the "real values"):
// e.g. ctx field's value, separators, type literals
var ColorRealValue = colors.FlagColorFgFaintBrightWhite

// ColorPlaceholderValue is used to denote "placeholder values":
// e.g. parser errors, function type (you can't copy these as-is to your code)
var ColorPlaceholderValue = colors.FlagColorFgFaintBrightWhite

// ColorType is used to write types,
// e.g. "[3]string" or "[]struct { a int; b int; c struct { d int } }"
var ColorType = colors.FlagColorFgFaintBrightBlack
