package logopts

import (
	"time"

	"github.com/ferdinandant/happylog/pkg/colors"
	"github.com/ferdinandant/happylog/pkg/core"
)

// FormatLogOpts represents what happylog formatters needs to know about a log event.
type FormatLogOpts struct {
	// E.g. Level.LevelError
	Level *core.Level
	// The name of the running application (optional)
	AppName *string
	// When the log function is called
	Now *time.Time
	// The message for the log event
	Msg *string
	// Pointer to the context object
	CtxPtr *interface{}
}

func (this *FormatLogOpts) ColorScheme() *colors.ColorScheme {
	return colors.ColorSchemeMap[*this.Level]
}
