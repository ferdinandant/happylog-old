package ctxparser

import (
	"fmt"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatBool(value interface{}) string {
	valueBool, ok := value.(bool)
	if !ok {
		err := fmt.Errorf("Cannot cast to bool: %+v", value)
		return FormatParserError(err)
	}

	if valueBool {
		return colors.FormatTextWithColor(ColorRealValue, "true")
	} else {
		return colors.FormatTextWithColor(ColorRealValue, "false")
	}
}
