package ctxparser

import (
	"fmt"
	"strings"
)

// ================================================================================
// MAIN
// ================================================================================

type ItemPrefixSuffixGenerator struct {
	lastItemIdx     int
	itemFirstPrefix string
	itemPrefix      string
	itemSuffix      string
	itemLastSuffix  string
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

func CreateItemPrefixSuffixGenerator(shouldPrintInOneLine bool, childrenDepth int, propertiesLength int) (*ItemPrefixSuffixGenerator, error) {
	if childrenDepth <= 0 {
		return nil, fmt.Errorf("childrenDepth must be positive: %d", childrenDepth)
	}

	// Create the generator
	generator := ItemPrefixSuffixGenerator{}
	generator.lastItemIdx = propertiesLength - 1
	if shouldPrintInOneLine {
		generator.itemFirstPrefix = " "
		generator.itemPrefix = " "
		generator.itemSuffix = ", "
		generator.itemLastSuffix = ", "
	} else {
		padding := strings.Repeat("  ", childrenDepth)
		generator.itemFirstPrefix = "\n" + padding
		generator.itemPrefix = padding
		generator.itemSuffix = ",\n"
		generator.itemLastSuffix = ",\n" + strings.Repeat("  ", childrenDepth-1)
	}

	return &generator, nil
}

func MustCreateItemPrefixSuffixGenerator(shouldPrintInOneLine bool, childrenDepth int, propertiesLength int) *ItemPrefixSuffixGenerator {
	usedChildrenDepth := childrenDepth
	usedShouldPrintInOneLine := shouldPrintInOneLine
	// Force correct values so it doesn't throw
	if childrenDepth <= 0 {
		usedChildrenDepth = 1
		usedShouldPrintInOneLine = true
	}

	generator, _ := CreateItemPrefixSuffixGenerator(usedShouldPrintInOneLine, usedChildrenDepth, propertiesLength)
	return generator
}
