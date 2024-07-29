package convert

import (
	"fmt"
	"strings"
)

// BoolConvert is a function type that converts any value to a pointer to bool.
// If the conversion fails, it returns nil.
// This allows for custom conversion logic to be implemented and passed to conversion functions.
//
// Examples:
//
//	customConverter := func(value interface{}) *bool {
//		if v, ok := value.(CustomType); ok {
//			b := v.CustomToBool()
//			return &b
//		}
//		return nil
//	}
//
//	result := ToBool(myCustomValue, customConverter)
type BoolConvert func(value interface{}) *bool

// ToBool converts any type of value to bool.
// It uses ToBoolE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to bool.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A boolean representation of the input value.
//
// Example:
//
//	b := ToBool("true")
//	fmt.Println(b) // Output: true
//
//	customConverter := func(value interface{}) *bool {
//		if v, ok := value.(CustomType); ok {
//			b := v.CustomToBool()
//			return &b
//		}
//		return nil
//	}
//	customBool := ToBool(myCustomValue, customConverter)
func ToBool(value interface{}, converters ...BoolConvert) bool {
	res, _ := ToBoolE(value, converters...)
	return res
}

// ToBoolOrDefault converts any type of value to bool.
// If the conversion fails, it returns the provided default value.
//
// Parameters:
//   - value: The value to be converted to bool.
//   - defaultValue: The default value to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A boolean representation of the input value or the default value if conversion fails.
//
// Example:
//
//	b := ToBoolOrDefault("true", false)
//	fmt.Println(b) // Output: true
//
//	b = ToBoolOrDefault(nil, true)
//	fmt.Println(b) // Output: true
//
//	customConverter := func(value interface{}) *bool {
//		if v, ok := value.(CustomType); ok {
//			b := v.CustomToBool()
//			return &b
//		}
//		return nil
//	}
//	customBool := ToBoolOrDefault(myCustomValue, false, customConverter)
func ToBoolOrDefault(value interface{}, defaultValue bool, converters ...BoolConvert) bool {
	if value == nil {
		return defaultValue
	}
	res, err := ToBoolE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToBoolE converts any type of value to bool or returns error.
// It handles various types including:
//   - bool, string, nil
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//
// For string type, it considers the following values as true:
// "1", "t", "T", "true", "TRUE", "True", "y", "yes", "Y", "YES", "Yes", "ok", "OK", "on", "ON"
// And the following values as false:
// "0", "f", "F", "false", "FALSE", "False", "n", "no", "N", "NO", "No", "off", "OFF"
//
// Parameters:
//   - value: The value to be converted to bool.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - A boolean representation of the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	b, err := ToBoolE("true")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(b) // Output: true
//
//	customConverter := func(value interface{}) *bool {
//		if v, ok := value.(CustomType); ok {
//			b := v.CustomToBool()
//			return &b
//		}
//		return nil
//	}
//	customBool, err := ToBoolE(myCustomValue, customConverter)
func ToBoolE(value interface{}, converters ...BoolConvert) (bool, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	v := Indirect(value)

	switch b := v.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int, int8, int16, int32, int64:
		return b != 0, nil
	case uint, uint8, uint16, uint32, uint64:
		return b != 0, nil
	case float32, float64:
		return b != 0, nil
	case string:
		switch strings.ToLower(b) {
		case "1", "t", "true", "y", "yes", "ok", "on":
			return true, nil
		case "0", "f", "false", "n", "no", "off", "":
			return false, nil
		default:
			return false, fmt.Errorf("convert: %v to boolean failed", value)
		}
	case []byte:
		return ToBoolE(string(b))
	default:
		return false, fmt.Errorf("convert: %v (%T) to boolean failed", value, value)
	}
}
