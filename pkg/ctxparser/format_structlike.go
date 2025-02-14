package ctxparser

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/ferdinandant/happylog/pkg/colors"
)

func FormatStruct(traversalCtx TraversalCtx) (result string, resultCtx *ParseResultCtx) {
	config := traversalCtx.Config
	fgColor := config.ColorScheme.FgFaint
	valueType := *traversalCtx.CurrentValueType
	valuePtr := traversalCtx.CurrentValuePtr

	// Get fields information
	// - https://stackoverflow.com/a/66511341/5181368
	// - https://pkg.go.dev/reflect#StructField
	structFields := reflect.VisibleFields(valueType)
	reflectValue := reflect.ValueOf(*valuePtr)
	isAllFieldLiteral := true
	iteratedFieldCount := 0
	var itemKeyStrList []string
	var itemValueStrList []string

	// Iterate fields
	// structField.PackagePath is empty IFF the field is exported.
	// - https://pkg.go.dev/reflect#StructField
	for _, structField := range structFields {
		iteratedFieldCount++
		if iteratedFieldCount > config.MaxFieldCount {
			break
		}
		// Process field
		isFieldExported := structField.PkgPath == ""
		fieldIndexPath := structField.Index
		var itemKey interface{} = structField.Name
		var itemValue interface{}
		var itemResult string
		var itemErr error
		var itemResultCtx *ParseResultCtx
		// Get formatted value
		if isFieldExported {
			itemValue, itemErr = getExportedFieldValue(reflectValue, fieldIndexPath)
		} else {
			itemValue, itemErr = getUnexportedFieldValue(reflectValue, fieldIndexPath)
		}
		if itemErr != nil {
			var provisionalValue interface{} = reflectValue.FieldByIndex(fieldIndexPath)
			childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, &itemKey, &provisionalValue)
			itemResult, itemResultCtx = FormatParserError(childrenTraversalCtx, itemErr, &provisionalValue)
		} else {
			childrenTraversalCtx := ExtendTraversalCtx(&traversalCtx, &itemKey, &itemValue)
			itemResult, itemResultCtx = FormatAny(childrenTraversalCtx)
		}
		// Maintain state
		if itemResultCtx != nil && !itemResultCtx.isLiteral {
			isAllFieldLiteral = false
		}
		// Append to temp storage
		itemKeyStrList = append(itemKeyStrList, structField.Name)
		itemValueStrList = append(itemValueStrList, itemResult)
	}

	// Iterate methods
	numMethods := 0
	if config.PrintPublicMethods {
		numMethods = valueType.NumMethod()
		for i := 0; i < numMethods; i++ {
			iteratedFieldCount++
			if iteratedFieldCount > config.MaxFieldCount {
				break
			}
			// Process method
			method := valueType.Method(i)
			methodName := method.Name + "()"
			methodType := method.Type.String()
			itemResult := formatMethodFieldWithType(config, methodType)
			// Append to temp storage
			itemKeyStrList = append(itemKeyStrList, methodName)
			itemValueStrList = append(itemValueStrList, itemResult)
		}
	}

	// Combine result
	valueStrResult := config.ColorMain
	childrenIndentLevel := traversalCtx.IndentLevel + 1
	childrenCount := len(structFields) + numMethods
	hasOmittedFields := iteratedFieldCount > config.MaxFieldCount
	shouldPrintInline := CheckShouldPrintInline(config, traversalCtx.Depth, isAllFieldLiteral)
	itemPsGenerator, err := CreateItemPrefixSuffixGenerator(shouldPrintInline, childrenIndentLevel, childrenCount, hasOmittedFields)
	if err != nil {
		return FormatParserError(traversalCtx, err, valuePtr)
	}
	// Print the fields
	for i, itemValueStr := range itemValueStrList {
		keyStr := itemKeyStrList[i] + ": "
		usedPrefix, usedSuffix := itemPsGenerator.GetPrefixSuffix(i)
		formattedValueStr := colors.FormatTextWithColor(fgColor, keyStr) + config.ColorMain + itemValueStr
		valueStrResult += usedPrefix + formattedValueStr + config.ColorMain + usedSuffix
	}
	// Print the ellipsis
	if hasOmittedFields {
		usedPrefix, usedSuffix := itemPsGenerator.GetPrefixSuffix(childrenCount)
		numberOfHiddenFields := childrenCount - config.MaxFieldCount
		formattedValueStr := "... " + strconv.Itoa(numberOfHiddenFields) + " hidden field(s)"
		valueStrResult += usedPrefix + formattedValueStr + config.ColorMain + usedSuffix
	}

	// Return result
	// We should use `reflect.TypeOf(...).String()` so it uses the struct name
	valueTypeStr := valueType.String()
	return formatStructWithType(config, valueTypeStr, valueStrResult), StructParseResultCtx
}

// ================================================================================
// HELPERS
// ================================================================================

func formatStructWithType(config *ParseConfig, typeStr string, valueStr string) string {
	return config.ColorType + typeStr + config.ColorMain + " {" + valueStr + "}" + colors.FlagReset
}

func formatMethodFieldWithType(config *ParseConfig, typeStr string) string {
	return config.ColorType + typeStr + colors.FlagReset
}

func getExportedFieldValue(structReflectValue reflect.Value, fieldIndexPath []int) (result interface{}, err error) {
	// Do not want to panic, e.g. when stepping through nil value
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			panicErrMsg := fmt.Sprintf("%+v", panicErr)
			err = fmt.Errorf("Panic: %s", panicErrMsg)
		}
	}()
	// We can only call `reflectValue.FieldByIndex(fieldIndexPath).Interface()` if the field is EXPORTED.
	// Otherwise calling `.Interface()` will panic.
	itemReflectValue := structReflectValue.FieldByIndex(fieldIndexPath)
	return itemReflectValue.Interface(), nil
}

func getUnexportedFieldValue(structReflectValue reflect.Value, fieldIndexPath []int) (result interface{}, err error) {
	// Do not want to panic, e.g. when accessing unaddressible address
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			err = fmt.Errorf("Panic: %+v", panicErr)
		}
	}()
	// Using some trick to access unexported field value
	// Need to copy the entire struct to handle "reflect: reflect.Value.Set using unaddressable value"
	// (copying just the item's reflect.Value somehow does not work)
	// - https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields
	addressibleStructReflectValue := reflect.New(structReflectValue.Type()).Elem()
	addressibleStructReflectValue.Set(structReflectValue)
	itemReflectValue := addressibleStructReflectValue.FieldByIndex(fieldIndexPath)
	itemReflectValue = reflect.NewAt(itemReflectValue.Type(), unsafe.Pointer(itemReflectValue.UnsafeAddr())).Elem()
	return itemReflectValue.Interface(), nil
}
