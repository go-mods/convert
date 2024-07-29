package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// IntConverter is a function type for custom int conversion.
// It takes any value and returns a pointer to an int if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customIntConverter := func(value interface{}) *int {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt(someValue, customIntConverter)
type IntConverter func(value interface{}) *int

// Int8Converter is a function type for custom int8 conversion.
// It takes any value and returns a pointer to an int8 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt8Converter := func(value interface{}) *int8 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt8()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt8(someValue, customInt8Converter)
type Int8Converter func(value interface{}) *int8

// Int16Converter is a function type for custom int16 conversion.
// It takes any value and returns a pointer to an int16 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt16Converter := func(value interface{}) *int16 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt16()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt16(someValue, customInt16Converter)
type Int16Converter func(value interface{}) *int16

// Int32Converter is a function type for custom int32 conversion.
// It takes any value and returns a pointer to an int32 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt32Converter := func(value interface{}) *int32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt32()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt32(someValue, customInt32Converter)
type Int32Converter func(value interface{}) *int32

// Int64Converter is a function type for custom int64 conversion.
// It takes any value and returns a pointer to an int64 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt64Converter := func(value interface{}) *int64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt64(someValue, customInt64Converter)
type Int64Converter func(value interface{}) *int64

// UintConverter is a function type for custom uint conversion.
// It takes any value and returns a pointer to a uint if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUintConverter := func(value interface{}) *uint {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint(someValue, customUintConverter)
type UintConverter func(value interface{}) *uint

// Uint8Converter is a function type for custom uint8 conversion.
// It takes any value and returns a pointer to a uint8 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint8Converter := func(value interface{}) *uint8 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint8()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint8(someValue, customUint8Converter)
type Uint8Converter func(value interface{}) *uint8

// Uint16Converter is a function type for custom uint16 conversion.
// It takes any value and returns a pointer to a uint16 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint16Converter := func(value interface{}) *uint16 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint16()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint16(someValue, customUint16Converter)
type Uint16Converter func(value interface{}) *uint16

// Uint32Converter is a function type for custom uint32 conversion.
// It takes any value and returns a pointer to a uint32 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint32Converter := func(value interface{}) *uint32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint32()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint32(someValue, customUint32Converter)
type Uint32Converter func(value interface{}) *uint32

// Uint64Converter is a function type for custom uint64 conversion.
// It takes any value and returns a pointer to a uint64 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint64Converter := func(value interface{}) *uint64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint64(someValue, customUint64Converter)
type Uint64Converter func(value interface{}) *uint64

// Float32Converter is a function type for custom float32 conversion.
// It takes any value and returns a pointer to a float32 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customFloat32Converter := func(value interface{}) *float32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToFloat32()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToFloat32(someValue, customFloat32Converter)
type Float32Converter func(value interface{}) *float32

// Float64Converter is a function type for custom float64 conversion.
// It takes any value and returns a pointer to a float64 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customFloat64Converter := func(value interface{}) *float64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToFloat64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToFloat64(someValue, customFloat64Converter)
type Float64Converter func(value interface{}) *float64

// IntArrayConverter is a function type for custom int array conversion.
// It takes any value and returns a pointer to an array of int if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customIntArrayConverter := func(value interface{}) *[]int {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToIntArray()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToIntArray(someValue, customIntArrayConverter)
type IntArrayConverter func(value interface{}) *[]int

// Int8ArrayConverter is a function type for custom int8 array conversion.
// It takes any value and returns a pointer to an array of int8 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt8ArrayConverter := func(value interface{}) *[]int8 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt8Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt8Array(someValue, customInt8ArrayConverter)
type Int8ArrayConverter func(value interface{}) *[]int8

// Int16ArrayConverter is a function type for custom int16 array conversion.
// It takes any value and returns a pointer to an array of int16 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt16ArrayConverter := func(value interface{}) *[]int16 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt16Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt16Array(someValue, customInt16ArrayConverter)
type Int16ArrayConverter func(value interface{}) *[]int16


// Int32ArrayConverter is a function type for custom int32 array conversion.
// It takes any value and returns a pointer to an array of int32 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt32ArrayConverter := func(value interface{}) *[]int32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt32Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt32Array(someValue, customInt32ArrayConverter)
type Int32ArrayConverter func(value interface{}) *[]int32

// Int64ArrayConverter is a function type for custom int64 array conversion.
// It takes any value and returns a pointer to an array of int64 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customInt64ArrayConverter := func(value interface{}) *[]int64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToInt64Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToInt64Array(someValue, customInt64ArrayConverter)
type Int64ArrayConverter func(value interface{}) *[]int64

// UintArrayConverter is a function type for custom uint array conversion.
// It takes any value and returns a pointer to an array of uint if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUintArrayConverter := func(value interface{}) *[]uint {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUintArray()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUintArray(someValue, customUintArrayConverter)
type UintArrayConverter func(value interface{}) *[]uint

// Uint8ArrayConverter is a function type for custom uint8 array conversion.
// It takes any value and returns a pointer to an array of uint8 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint8ArrayConverter := func(value interface{}) *[]uint8 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint8Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint8Array(someValue, customUint8ArrayConverter)
type Uint8ArrayConverter func(value interface{}) *[]uint8

// Uint16ArrayConverter is a function type for custom uint16 array conversion.
// It takes any value and returns a pointer to an array of uint16 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint16ArrayConverter := func(value interface{}) *[]uint16 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint16Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint16Array(someValue, customUint16ArrayConverter)
type Uint16ArrayConverter func(value interface{}) *[]uint16

// Uint32ArrayConverter is a function type for custom uint32 array conversion.
// It takes any value and returns a pointer to an array of uint32 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint32ArrayConverter := func(value interface{}) *[]uint32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint32Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint32Array(someValue, customUint32ArrayConverter)
type Uint32ArrayConverter func(value interface{}) *[]uint32

// Uint64ArrayConverter is a function type for custom uint64 array conversion.
// It takes any value and returns a pointer to an array of uint64 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customUint64ArrayConverter := func(value interface{}) *[]uint64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToUint64Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToUint64Array(someValue, customUint64ArrayConverter)
type Uint64ArrayConverter func(value interface{}) *[]uint64

// Float32ArrayConverter is a function type for custom float32 array conversion.
// It takes any value and returns a pointer to an array of float32 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customFloat32ArrayConverter := func(value interface{}) *[]float32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToFloat32Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToFloat32Array(someValue, customFloat32ArrayConverter)
type Float32ArrayConverter func(value interface{}) *[]float32

// Float64ArrayConverter is a function type for custom float64 array conversion.
// It takes any value and returns a pointer to an array of float64 if conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
//
// Example usage:
//
//	customFloat64ArrayConverter := func(value interface{}) *[]float64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToFloat64Array()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToFloat64Array(someValue, customFloat64ArrayConverter)
type Float64ArrayConverter func(value interface{}) *[]float64

// ToIntE converts any type of value to int or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToIntE(value interface{}, converters ...IntConverter) (int, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return int(res64), nil
		} else if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return int(resU64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return int(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToIntE(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to int failed", v)
		}
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return int(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return int(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToIntE(valueStr)
	}
}

// ToInt converts any type of value to int, ignoring errors.
func ToInt(value interface{}, converters ...IntConverter) int {
	res, _ := ToIntE(value, converters...)
	return res
}

// ToIntOrDefault converts any type of value to int or returns the provided default value if conversion fails.
func ToIntOrDefault(value interface{}, defaultValue int, converters ...IntConverter) int {
	if value == nil {
		return defaultValue
	}
	res, err := ToIntE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt8E converts any type of value to int8 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt8E(value interface{}, converters ...Int8Converter) (int8, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return int8(res64), nil
		} else if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return int8(resU64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return int8(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToInt8E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to int8 failed", v)
		}
	case int, int8, int16, int32, int64:
		return int8(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return int8(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return int8(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToInt8E(valueStr)
	}
}

// ToInt8 converts any type of value to int8, ignoring errors.
func ToInt8(value interface{}, converters ...Int8Converter) int8 {
	res, _ := ToInt8E(value, converters...)
	return res
}

// ToInt8OrDefault converts any type of value to int8 or returns the provided default value if conversion fails.
func ToInt8OrDefault(value interface{}, defaultValue int8, converters ...Int8Converter) int8 {
	if value == nil {
		return defaultValue
	}
	res, err := ToInt8E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt16E converts any type of value to int16 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt16E(value interface{}, converters ...Int16Converter) (int16, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return int16(res64), nil
		} else if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return int16(resU64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return int16(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToInt16E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to int16 failed", v)
		}
	case int, int8, int16, int32, int64:
		return int16(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return int16(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return int16(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToInt16E(valueStr)
	}
}

// ToInt16 converts any type of value to int16, ignoring errors.
func ToInt16(value interface{}, converters ...Int16Converter) int16 {
	res, _ := ToInt16E(value, converters...)
	return res
}

// ToInt16OrDefault converts any type of value to int16 or returns the provided default value if conversion fails.
func ToInt16OrDefault(value interface{}, defaultValue int16, converters ...Int16Converter) int16 {
	if value == nil {
		return defaultValue
	}
	res, err := ToInt16E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt32E converts any type of value to int32 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt32E(value interface{}, converters ...Int32Converter) (int32, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return int32(res64), nil
		} else if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return int32(resU64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return int32(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToInt32E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to int32 failed", v)
		}
	case int, int8, int16, int32, int64:
		return int32(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return int32(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return int32(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToInt32E(valueStr)
	}
}

// ToInt32 converts any type of value to int32, ignoring errors.
func ToInt32(value interface{}, converters ...Int32Converter) int32 {
	res, _ := ToInt32E(value, converters...)
	return res
}

// ToInt32OrDefault converts any type of value to int32 or returns the provided default value if conversion fails.
func ToInt32OrDefault(value interface{}, defaultValue int32, converters ...Int32Converter) int32 {
	if value == nil {
		return defaultValue
	}
	res, err := ToInt32E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt64E converts any type of value to int64 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt64E(value interface{}, converters ...Int64Converter) (int64, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return res64, nil
		} else if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return int64(resU64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return int64(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToInt64E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to int64 failed", v)
		}
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int(), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return int64(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return int64(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToInt64E(valueStr)
	}
}

// ToInt64 converts any type of value to int64, ignoring errors.
func ToInt64(value interface{}, converters ...Int64Converter) int64 {
	res, _ := ToInt64E(value, converters...)
	return res
}

// ToInt64OrDefault converts any type of value to int64 or returns the provided default value if conversion fails.
func ToInt64OrDefault(value interface{}, defaultValue int64, converters ...Int64Converter) int64 {
	if value == nil {
		return defaultValue
	}
	res, err := ToInt64E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUintE converts any type of value to uint or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUintE(value interface{}, converters ...UintConverter) (uint, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return uint(resU64), nil
		} else if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return uint(res64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return uint(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToUintE(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to uint failed", v)
		}
	case int, int8, int16, int32, int64:
		return uint(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return uint(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return uint(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToUintE(valueStr)
	}
}

// ToUint converts any type of value to uint, ignoring errors.
func ToUint(value interface{}, converters ...UintConverter) uint {
	res, _ := ToUintE(value, converters...)
	return res
}

// ToUintOrDefault converts any type of value to uint or returns the provided default value if conversion fails.
func ToUintOrDefault(value interface{}, defaultValue uint, converters ...UintConverter) uint {
	if value == nil {
		return defaultValue
	}
	res, err := ToUintE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint8E converts any type of value to uint8 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint8E(value interface{}, converters ...Uint8Converter) (uint8, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return uint8(resU64), nil
		} else if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return uint8(res64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return uint8(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToUint8E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to uint8 failed", v)
		}
	case int, int8, int16, int32, int64:
		return uint8(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return uint8(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return uint8(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToUint8E(valueStr)
	}
}

// ToUint8 converts any type of value to uint8, ignoring errors.
func ToUint8(value interface{}, converters ...Uint8Converter) uint8 {
	res, _ := ToUint8E(value, converters...)
	return res
}

// ToUint8OrDefault converts any type of value to uint8 or returns the provided default value if conversion fails.
func ToUint8OrDefault(value interface{}, defaultValue uint8, converters ...Uint8Converter) uint8 {
	if value == nil {
		return defaultValue
	}
	res, err := ToUint8E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint16E converts any type of value to uint16 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint16E(value interface{}, converters ...Uint16Converter) (uint16, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return uint16(resU64), nil
		} else if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return uint16(res64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return uint16(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToUint16E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to uint16 failed", v)
		}
	case int, int8, int16, int32, int64:
		return uint16(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return uint16(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return uint16(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToUint16E(valueStr)
	}
}

// ToUint16 converts any type of value to uint16, ignoring errors.
func ToUint16(value interface{}, converters ...Uint16Converter) uint16 {
	res, _ := ToUint16E(value, converters...)
	return res
}

// ToUint16OrDefault converts any type of value to uint16 or returns the provided default value if conversion fails.
func ToUint16OrDefault(value interface{}, defaultValue uint16, converters ...Uint16Converter) uint16 {
	if value == nil {
		return defaultValue
	}
	res, err := ToUint16E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint32E converts any type of value to uint32 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint32E(value interface{}, converters ...Uint32Converter) (uint32, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return uint32(resU64), nil
		} else if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return uint32(res64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return uint32(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToUint32E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to uint32 failed", v)
		}
	case int, int8, int16, int32, int64:
		return uint32(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return uint32(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return uint32(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToUint32E(valueStr)
	}
}

// ToUint32 converts any type of value to uint32, ignoring errors.
func ToUint32(value interface{}, converters ...Uint32Converter) uint32 {
	res, _ := ToUint32E(value, converters...)
	return res
}

// ToUint32OrDefault converts any type of value to uint32 or returns the provided default value if conversion fails.
func ToUint32OrDefault(value interface{}, defaultValue uint32, converters ...Uint32Converter) uint32 {
	if value == nil {
		return defaultValue
	}
	res, err := ToUint32E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint64E converts any type of value to uint64 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint64E(value interface{}, converters ...Uint64Converter) (uint64, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resU64, err := strconv.ParseUint(v, 0, 64); err == nil {
			return resU64, nil
		} else if res64, err := strconv.ParseInt(v, 0, 64); err == nil {
			return uint64(res64), nil
		} else if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return uint64(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToUint64E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to uint64 failed", v)
		}
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return uint64(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return uint64(reflect.ValueOf(v).Float()), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToUint64E(valueStr)
	}
}

// ToUint64 converts any type of value to uint64, ignoring errors.
func ToUint64(value interface{}, converters ...Uint64Converter) uint64 {
	res, _ := ToUint64E(value, converters...)
	return res
}

// ToUint64OrDefault converts any type of value to uint64 or returns the provided default value if conversion fails.
func ToUint64OrDefault(value interface{}, defaultValue uint64, converters ...Uint64Converter) uint64 {
	if value == nil {
		return defaultValue
	}
	res, err := ToUint64E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToFloat32E converts any type of value to float32 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToFloat32E(value interface{}, converters ...Float32Converter) (float32, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return float32(resF64), nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToFloat32E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to float32 failed", v)
		}
	case int, int8, int16, int32, int64:
		return float32(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return float32(reflect.ValueOf(v).Uint()), nil
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToFloat32E(valueStr)
	}
}

// ToFloat32 converts any type of value to float32, ignoring errors.
func ToFloat32(value interface{}, converters ...Float32Converter) float32 {
	res, _ := ToFloat32E(value, converters...)
	return res
}

// ToFloat32OrDefault converts any type of value to float32 or returns the provided default value if conversion fails.
func ToFloat32OrDefault(value interface{}, defaultValue float32, converters ...Float32Converter) float32 {
	if value == nil {
		return defaultValue
	}
	res, err := ToFloat32E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToFloat64E converts any type of value to float64 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToFloat64E(value interface{}, converters ...Float64Converter) (float64, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case string:
		if len(v) == 0 {
			return 0, nil
		}
		if resF64, err := strconv.ParseFloat(v, 64); err == nil {
			return resF64, nil
		} else if resBool, err := strconv.ParseBool(v); err == nil {
			return ToFloat64E(resBool)
		} else {
			return 0, fmt.Errorf("convert: string \"%s\" to float64 failed", v)
		}
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return float64(reflect.ValueOf(v).Uint()), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		valueStr := ToString(v)
		return ToFloat64E(valueStr)
	}
}

// ToFloat64 converts any type of value to float64, ignoring errors.
func ToFloat64(value interface{}, converters ...Float64Converter) float64 {
	res, _ := ToFloat64E(value, converters...)
	return res
}

// ToFloat64OrDefault converts any type of value to float64 or returns the provided default value if conversion fails.
func ToFloat64OrDefault(value interface{}, defaultValue float64, converters ...Float64Converter) float64 {
	if value == nil {
		return defaultValue
	}
	res, err := ToFloat64E(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToIntArrayE converts any type of value to an array of int or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToIntArrayE(value interface{}, converters ...IntArrayConverter) ([]int, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]int, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToIntE(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToIntArray converts any type of value to an array of int, ignoring errors.
func ToIntArray(value interface{}, converters ...IntArrayConverter) []int {
	res, _ := ToIntArrayE(value, converters...)
	return res
}

// ToIntArrayOrDefault converts any type of value to an array of int or returns the provided default array if conversion fails.
func ToIntArrayOrDefault(value interface{}, defaultValue []int, converters ...IntArrayConverter) []int {
	res, err := ToIntArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt8ArrayE converts any type of value to an array of int8 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt8ArrayE(value interface{}, converters ...Int8ArrayConverter) ([]int8, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]int8, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToInt8E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToInt8Array convertit n'importe quel type de valeur en un tableau d'int8, en ignorant les erreurs.
func ToInt8Array(value interface{}, converters ...Int8ArrayConverter) []int8 {
	res, _ := ToInt8ArrayE(value, converters...)
	return res
}

// ToInt8ArrayOrDefault converts any type of value to an array of int8 or returns the provided default array if conversion fails.
func ToInt8ArrayOrDefault(value interface{}, defaultValue []int8, converters ...Int8ArrayConverter) []int8 {
	res, err := ToInt8ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt16ArrayE converts any type of value to an array of int16 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt16ArrayE(value interface{}, converters ...Int16ArrayConverter) ([]int16, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]int16, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToInt16E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToInt16Array converts any type of value to an array of int16, ignoring errors.
func ToInt16Array(value interface{}, converters ...Int16ArrayConverter) []int16 {
	res, _ := ToInt16ArrayE(value, converters...)
	return res
}

// ToInt16ArrayOrDefault converts any type of value to an array of int16 or returns the provided default array if conversion fails.
func ToInt16ArrayOrDefault(value interface{}, defaultValue []int16, converters ...Int16ArrayConverter) []int16 {
	res, err := ToInt16ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt32ArrayE converts any type of value to an array of int32 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt32ArrayE(value interface{}, converters ...Int32ArrayConverter) ([]int32, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]int32, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToInt32E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToInt32Array converts any type of value to an array of int32, ignoring errors.
func ToInt32Array(value interface{}, converters ...Int32ArrayConverter) []int32 {
	res, _ := ToInt32ArrayE(value, converters...)
	return res
}

// ToInt32ArrayOrDefault converts any type of value to an array of int32 or returns the provided default array if conversion fails.
func ToInt32ArrayOrDefault(value interface{}, defaultValue []int32, converters ...Int32ArrayConverter) []int32 {
	res, err := ToInt32ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToInt64ArrayE converts any type of value to an array of int64 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToInt64ArrayE(value interface{}, converters ...Int64ArrayConverter) ([]int64, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]int64, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToInt64E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToInt64Array converts any type of value to an array of int64, ignoring errors.
func ToInt64Array(value interface{}, converters ...Int64ArrayConverter) []int64 {
	res, _ := ToInt64ArrayE(value, converters...)
	return res
}

// ToInt64ArrayOrDefault converts any type of value to an array of int64 or returns the provided default array if conversion fails.
func ToInt64ArrayOrDefault(value interface{}, defaultValue []int64, converters ...Int64ArrayConverter) []int64 {
	res, err := ToInt64ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUintArrayE converts any type of value to an array of uint or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUintArrayE(value interface{}, converters ...UintArrayConverter) ([]uint, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array,
		reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]uint, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToUintE(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToUintArray converts any type of value to an array of uint, ignoring errors.
func ToUintArray(value interface{}, converters ...UintArrayConverter) []uint {
	res, _ := ToUintArrayE(value, converters...)
	return res
}

// ToUintArrayOrDefault converts any type of value to an array of uint or returns the provided default array if conversion fails.
func ToUintArrayOrDefault(value interface{}, defaultValue []uint, converters ...UintArrayConverter) []uint {
	res, err := ToUintArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint8ArrayE converts any type of value to an array of uint8 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint8ArrayE(value interface{}, converters ...Uint8ArrayConverter) ([]uint8, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]uint8, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToUint8E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToUint8Array converts any type of value to an array of uint8, ignoring errors.
func ToUint8Array(value interface{}, converters ...Uint8ArrayConverter) []uint8 {
	res, _ := ToUint8ArrayE(value, converters...)
	return res
}

// ToUint8ArrayOrDefault converts any type of value to an array of uint8 or returns the provided default array if conversion fails.
func ToUint8ArrayOrDefault(value interface{}, defaultValue []uint8, converters ...Uint8ArrayConverter) []uint8 {
	res, err := ToUint8ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint16ArrayE converts any type of value to an array of uint16 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint16ArrayE(value interface{}, converters ...Uint16ArrayConverter) ([]uint16, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]uint16, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToUint16E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToUint16Array converts any type of value to an array of uint16, ignoring errors.
func ToUint16Array(value interface{}, converters ...Uint16ArrayConverter) []uint16 {
	res, _ := ToUint16ArrayE(value, converters...)
	return res
}

// ToUint16ArrayOrDefault converts any type of value to an array of uint16 or returns the provided default array if conversion fails.
func ToUint16ArrayOrDefault(value interface{}, defaultValue []uint16, converters ...Uint16ArrayConverter) []uint16 {
	res, err := ToUint16ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint32ArrayE converts any type of value to an array of uint32 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint32ArrayE(value interface{}, converters ...Uint32ArrayConverter) ([]uint32, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]uint32, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToUint32E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToUint32Array converts any type of value to an array of uint32, ignoring errors.
func ToUint32Array(value interface{}, converters ...Uint32ArrayConverter) []uint32 {
	res, _ := ToUint32ArrayE(value, converters...)
	return res
}

// ToUint32ArrayOrDefault converts any type of value to an array of uint32 or returns the provided default array if conversion fails.
func ToUint32ArrayOrDefault(value interface{}, defaultValue []uint32, converters ...Uint32ArrayConverter) []uint32 {
	res, err := ToUint32ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToUint64ArrayE converts any type of value to an array of uint64 or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
func ToUint64ArrayE(value interface{}, converters ...Uint64ArrayConverter) ([]uint64, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		v := reflect.ValueOf(value)
		resArray := make([]uint64, v.Len())
		for i := 0; i < v.Len(); i++ {
			res, err := ToUint64E(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("convert: cannot convert %v at index %d", v.Index(i).Interface(), i)
			}
			resArray[i] = res
		}
		return resArray, nil
	default:
		return nil, fmt.Errorf("convert: %T is not an array or slice", value)
	}
}

// ToUint64Array converts any type of value to an array of uint64, ignoring errors.
func ToUint64Array(value interface{}, converters ...Uint64ArrayConverter) []uint64 {
	res, _ := ToUint64ArrayE(value, converters...)
	return res
}

// ToUint64ArrayOrDefault converts any type of value to an array of uint64 or returns the provided default array if conversion fails.
func ToUint64ArrayOrDefault(value interface{}, defaultValue []uint64, converters ...Uint64ArrayConverter) []uint64 {
	res, err := ToUint64ArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}
