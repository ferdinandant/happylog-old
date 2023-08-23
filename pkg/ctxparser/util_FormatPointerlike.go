package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatUnsafePointer(traversalCtx TraversalCtx) string {
	config := traversalCtx.Config
	value := *traversalCtx.CurrentValuePtr
	valueStr := fmt.Sprintf("%v", value)
	return config.ColorMain + valueStr + colors.FlagReset
}
