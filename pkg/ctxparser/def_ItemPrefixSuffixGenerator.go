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

func CreateItemPrefixSuffixGenerator(shouldPrintInOneLine bool, childrenIndentLevel int, propertiesLength int) (*ItemPrefixSuffixGenerator, error) {
	if childrenIndentLevel <= 0 {
		return nil, fmt.Errorf("childrenIndentLevel must be positive: %d", childrenIndentLevel)
	}

	// Create the generator
	generator := ItemPrefixSuffixGenerator{}
	generator.lastItemIdx = propertiesLength - 1
	if shouldPrintInOneLine {
		generator.itemFirstPrefix = " "
		generator.itemPrefix = " "
		generator.itemSuffix = ","
		generator.itemLastSuffix = " "
	} else {
		padding := strings.Repeat("  ", childrenIndentLevel)
		generator.itemFirstPrefix = "\n" + padding
		generator.itemPrefix = padding
		generator.itemSuffix = ",\n"
		generator.itemLastSuffix = ",\n" + strings.Repeat("  ", childrenIndentLevel-1)
	}

	return &generator, nil
}

func MustCreateItemPrefixSuffixGenerator(shouldPrintInOneLine bool, childrenIndentLevel int, propertiesLength int) *ItemPrefixSuffixGenerator {
	usedChildrenIndentLevel := childrenIndentLevel
	usedShouldPrintInOneLine := shouldPrintInOneLine
	// Force correct values so it doesn't throw
	if childrenIndentLevel <= 0 {
		usedChildrenIndentLevel = 1
		usedShouldPrintInOneLine = true
	}

	generator, _ := CreateItemPrefixSuffixGenerator(usedShouldPrintInOneLine, usedChildrenIndentLevel, propertiesLength)
	return generator
}
