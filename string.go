package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// StringConvert is a function type that converts any value to a pointer to string.
// If the conversion fails, it returns nil.
// This allows for custom conversion logic to be implemented and passed to conversion functions.
//
// Examples:
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			s := v.CustomToString()
//			return &s
//		}
//		return nil
//	}
//
//	result := ToString(myCustomValue, customConverter)
type StringConvert func(value interface{}) *string

// ToString converts any type of value to string.
// It uses ToStringE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to string.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A string representation of the input value.
//
// Example:
//
//	str := ToString(42)
//	fmt.Println(str) // Output: "42"
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(time.Time); ok {
//			s := v.Format("2006-01-02")
//			return &s
//		}
//		return nil
//	}
//	dateStr := ToString(time.Now(), customConverter)
//	fmt.Println(dateStr) // Output: "2023-04-15" (example date)
func ToString(value interface{}, converters ...StringConvert) string {
	s, _ := ToStringE(value, converters...)
	return s
}

// ToStringE converts any type of value to string or returns an error.
// It handles various types including:
//   - string, []byte, nil, bool
//   - integer types (int, int8, int16, int32, int64)
//   - unsigned integer types (uint, uint8, uint16, uint32, uint64)
//   - float types (float32, float64)
//   - fmt.Stringer, fmt.GoStringer, and error interfaces
//
// For other types, it uses fmt.Sprintf("%v", s).
//
// Parameters:
//   - value: The value to be converted to string.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - A string representation of the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	str, err := ToStringE(42)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(str) // Output: "42"
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.([]int); ok {
//			s := fmt.Sprintf("%v", v)
//			return &s
//		}
//		return nil
//	}
//	arrStr, err := ToStringE([]int{1, 2, 3}, customConverter)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(arrStr) // Output: "[1 2 3]"
func ToStringE(value interface{}, converters ...StringConvert) (string, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	i := Indirect(value)

	switch s := i.(type) {
	case string:
		return s, nil
	case []byte:
		return string(s), nil
	case nil:
		return "", nil
	case bool:
		return strconv.FormatBool(s), nil
	case int:
		return strconv.Itoa(s), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int32:
		return strconv.FormatInt(int64(s), 10), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(s, 10), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case fmt.Stringer:
		return s.String(), nil
	case fmt.GoStringer:
		return s.GoString(), nil
	default:
		{
			if e, ok := value.(error); ok {
				return e.Error(), nil
			}
			return fmt.Sprintf("%v", s), nil
		}
	}
}

// ToStringOrDefault converts any type of value to string.
// If the conversion fails, it returns the provided default value.
//
// Parameters:
//   - value: The value to be converted to string.
//   - defaultValue: The default value to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A string representation of the input value or the default value if conversion fails.
//
// Example:
//
//	str := ToStringOrDefault(42, "default")
//	fmt.Println(str) // Output: "42"
//
//	str = ToStringOrDefault(nil, "default")
//	fmt.Println(str) // Output: "default"
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(time.Time); ok {
//			s := v.Format("2006-01-02")
//			return &s
//		}
//		return nil
//	}
//	dateStr := ToStringOrDefault(time.Now(), "default", customConverter)
//	fmt.Println(dateStr) // Output: "2023-04-15" (example date)
func ToStringOrDefault(value interface{}, defaultValue string, converters ...StringConvert) string {
	s, err := ToStringE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return s
}

// ToStringArray converts any type of value to an array of strings.
// It uses ToStringArrayE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to an array of strings.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	An array of strings representing the input value.
//
// Example:
//
//	arr := ToStringArray([]int{1, 2, 3})
//	fmt.Println(arr) // Output: ["1", "2", "3"]
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(time.Time); ok {
//			s := v.Format("2006-01-02")
//			return &s
//		}
//		return nil
//	}
//	dates := ToStringArray([]time.Time{time.Now(), time.Now().AddDate(0, 0, 1)}, customConverter)
//	fmt.Println(dates) // Output: ["2023-04-15", "2023-04-16"] (example dates)
func ToStringArray(value interface{}, converters ...StringConvert) []string {
	result, _ := ToStringArrayE(value, converters...)
	return result
}

// ToStringArrayOrDefault converts any type of value to an array of strings.
// If the conversion fails, it returns the provided default array.
//
// Parameters:
//   - value: The value to be converted to an array of strings.
//   - defaultValue: The default array to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	An array of strings representing the input value or the default array if conversion fails.
//
// Example:
//
//	arr := ToStringArrayOrDefault([]int{1, 2, 3}, []string{"default"})
//	fmt.Println(arr) // Output: ["1", "2", "3"]
//
//	arr = ToStringArrayOrDefault(nil, []string{"default"})
//	fmt.Println(arr) // Output: ["default"]
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(time.Time); ok {
//			s := v.Format("2006-01-02")
//			return &s
//		}
//		return nil
//	}
//	dates := ToStringArrayOrDefault([]time.Time{time.Now(), time.Now().AddDate(0, 0, 1)}, []string{"default"}, customConverter)
//	fmt.Println(dates) // Output: ["2023-04-15", "2023-04-16"] (example dates)
func ToStringArrayOrDefault(value interface{}, defaultValue []string, converters ...StringConvert) []string {
	result, err := ToStringArrayE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return result
}

// ToStringArrayE converts any type of value to an array of strings or returns an error.
// It handles slices, arrays, and single values.
//
// Parameters:
//   - value: The value to be converted to an array of strings.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - An array of strings representing the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	arr, err := ToStringArrayE([]int{1, 2, 3})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(arr) // Output: ["1", "2", "3"]
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			s := v.CustomToString()
//			return &s
//		}
//		return nil
//	}
//	customArr, err := ToStringArrayE([]CustomType{custom1, custom2}, customConverter)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(customArr) // Output: ["custom1", "custom2"]
func ToStringArrayE(value interface{}, converters ...StringConvert) ([]string, error) {
	if value == nil {
		return []string{}, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Array, reflect.Slice:
		length := v.Len()
		result := make([]string, length)
		for i := 0; i < length; i++ {
			str, err := ToStringE(v.Index(i).Interface(), converters...)
			if err != nil {
				return nil, fmt.Errorf("error converting element at index %d: %v", i, err)
			}
			result[i] = str
		}
		return result, nil
	default:
		str, err := ToStringE(value, converters...)
		if err != nil {
			return nil, err
		}
		return []string{str}, nil
	}
}
