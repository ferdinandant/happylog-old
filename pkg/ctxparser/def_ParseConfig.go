package ctxparser

import (
	"github.com/ferdinandant/happylog/pkg/colors"
)

type ParseConfig struct {
	ColorScheme           *colors.ColorScheme
	ColorRealValue        colors.Color
	ColorPlaceholderValue colors.Color
	ColorType             colors.Color
}
