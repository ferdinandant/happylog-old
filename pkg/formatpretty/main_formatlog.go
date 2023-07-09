package formatpretty

import (
	"fmt"
	"io"
	"os"

	"github.com/ferdinandant/happylog/pkg/core"
)

func FormatLog(logOpts *core.FormatLogOpts) {
	var w io.Writer = os.Stdout
	if *logOpts.Level >= core.LevelError {
		w = os.Stderr
	}

	// Print header line
	formattedLabelTag := GetFormattedLabelTag(logOpts)
	formattedTimestampSection := GetFormattedTimestampSection(logOpts)
	formattedHeaderLine := formattedLabelTag + " " + formattedTimestampSection
	fmt.Fprint(w, formattedHeaderLine, "\n")

	// Print message
	formattedMessage := GetFormattedMessage(logOpts)
	fmt.Fprint(w, formattedMessage, "\n")

	// Print context
	if logOpts.CtxPtr != nil {
		formattedContext := GetFormattedContext(logOpts)
		fmt.Fprint(w, formattedContext, "\n")
	}
}
