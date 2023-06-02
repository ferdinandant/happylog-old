package formatpretty

import (
	"fmt"
	"os"

	"github.com/ferdinandant/happylog/pkg/types"
)

func Log(logOpts *types.LogOpts) {

	// Print header line
	formattedLabelTag := GetFormattedLabelTag(logOpts)
	formattedTimestampSection := GetFormattedTimestampSection(logOpts)
	formattedHeaderLine := formattedLabelTag + " " + formattedTimestampSection
	fmt.Fprint(os.Stdout, formattedHeaderLine, "\n")

	// Print message
	formattedMessage := GetFormattedMessage(logOpts)
	fmt.Fprint(os.Stdout, formattedMessage, "\n")
}
