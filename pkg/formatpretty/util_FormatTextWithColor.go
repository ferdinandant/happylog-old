package formatpretty

import (
	"github.com/ferdinandant/happylog/pkg/core"
)

func FormatTextWithColor(color core.Color, str string) string {
	return color + str + core.FlagReset
}
