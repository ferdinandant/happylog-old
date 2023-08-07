package ctxparser

import "github.com/ferdinandant/happylog/pkg/colors"

// ColorMain is used for most purposes (for the "real values"):
// e.g. ctx field's value, separators, type literals
var ColorMain = colors.FlagColorFgFaintBrightWhite

// ColorType is used to write types,
// e.g. "[3]string" or "[]struct { a int; b int; c struct { d int } }"
var ColorType = colors.FlagColorFgFaintBrightBlack
