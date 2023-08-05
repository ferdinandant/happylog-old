package ctxparser

import (
	"fmt"
	"strings"
)

// ================================================================================
// MAIN
// ================================================================================

type ItemPrefixSuffixGenerator struct {
	itemFirstPrefix string
	itemPrefix      string
	itemSuffix      string
	itemLastSuffix  string
	lastItemIdx     int
}

func (this *ItemPrefixSuffixGenerator) GetPrefixSuffix(itemIdx int) (prefix string, suffix string) {
	if itemIdx == 0 {
		prefix = this.itemFirstPrefix
		suffix = this.itemSuffix
	} else if itemIdx == this.lastItemIdx {
		prefix = this.itemPrefix
		suffix = this.itemLastSuffix
	} else {
		prefix = this.itemPrefix
		suffix = this.itemSuffix
	}
	return
}

// ================================================================================
// HELPERS
// ================================================================================

func CreateItemPrefixSuffixGenerator(shouldPrintInOneLine bool, childrenDepth int) (*ItemPrefixSuffixGenerator, error) {
	if childrenDepth <= 0 {
		return nil, fmt.Errorf("childrenDepth must be positive: %d", childrenDepth)
	}

	return nil, nil
}

func GetItemPrefixSuffix(shouldPrintInOneLine bool, childrenDepth int) (
	itemFirstPrefix string, itemPrefix string, itemSuffix string, itemLastSuffix string,
) {
	// Don't wanna deal with potential panic of `strings.Repeat` when the count is negative.
	// childrenDepth should be 1 or higher; this is just a defensive code.
	if shouldPrintInOneLine {
		itemFirstPrefix = " "
		itemPrefix = " "
		itemSuffix = ", "
		itemLastSuffix = ", "
	} else {
		padding := strings.Repeat("  ", childrenDepth)
		itemFirstPrefix = "\n" + padding
		itemPrefix = padding
		itemSuffix = ",\n"
		itemLastSuffix = ",\n" + strings.Repeat("  ", childrenDepth-1)
	}
	return
}
