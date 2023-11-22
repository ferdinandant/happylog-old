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
	// Set prefix
	if itemIdx == 0 {
		prefix = this.itemFirstPrefix
	} else {
		prefix = this.itemPrefix
	}
	// Set suffix
	if itemIdx == this.lastItemIdx {
		suffix = this.itemLastSuffix
	} else {
		suffix = this.itemSuffix
	}
	return
}

// ================================================================================
// HELPERS
// ================================================================================

func CreateItemPrefixSuffixGenerator(
	shouldPrintInOneLine bool,
	childrenIndentLevel int,
	propertyCount int,
	hasOmittedProperty bool,
) (*ItemPrefixSuffixGenerator, error) {
	if childrenIndentLevel <= 0 {
		return nil, fmt.Errorf("childrenIndentLevel must be positive: %d", childrenIndentLevel)
	}

	// If it has ommited fields, there is one extra item
	// (for the ellipsis, e.g. "... 10 hidden items")
	usedPropertyCount := propertyCount
	if hasOmittedProperty {
		usedPropertyCount += 1
	}

	// Create the generator
	generator := ItemPrefixSuffixGenerator{}
	generator.lastItemIdx = usedPropertyCount - 1
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

func MustCreateItemPrefixSuffixGenerator(
	shouldPrintInOneLine bool,
	childrenIndentLevel int,
	propertiesLength int,
	hasOmittedProperty bool,
) *ItemPrefixSuffixGenerator {
	usedChildrenIndentLevel := childrenIndentLevel
	usedShouldPrintInOneLine := shouldPrintInOneLine

	// Force correct values so it doesn't throw
	if childrenIndentLevel <= 0 {
		usedChildrenIndentLevel = 1
		usedShouldPrintInOneLine = true
	}

	generator, _ := CreateItemPrefixSuffixGenerator(
		usedShouldPrintInOneLine,
		usedChildrenIndentLevel,
		propertiesLength,
		hasOmittedProperty,
	)
	return generator
}
