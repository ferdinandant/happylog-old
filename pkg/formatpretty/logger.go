package formatpretty

import (
	"fmt"
	"os"
	"time"

	"github.com/ferdinandant/happylog/pkg/levels"
)

func Log(level levels.Level, now time.Time, msg string, ctx ...interface{}) {
	formattedLabelTag := GetFormattedLabelTag(level)
	timestamp := formatTime(now)

	formattedHeaderLine := (formattedLabelTag + " " + timestamp)

	// Print message
	fmt.Fprint(os.Stdout, formattedHeaderLine, "\n")
	fmt.Fprint(os.Stdout, msg, "\n")
}

func formatTime(now time.Time) string {
	// Formats "Jan _2 15:04:05.000" -> "15:04:05.000"
	return now.Format(time.StampMilli)[7:]
}
