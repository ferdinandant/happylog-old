package ctxparser

import "github.com/ferdinandant/happylog/pkg/colors"

// ColorRealValue is used for most purposes (for the "real values"):
// e.g. ctx field's value, separators, type literals
const ColorRealValue = colors.FlagColorFgFaintBrightWhite

// ColorPlaceholderValue is used to denote "placeholder values":
// e.g. errors, function type
const ColorPlaceholderValue = colors.FlagColorFgFaintBrightBlack
