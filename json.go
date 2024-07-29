package convert

import (
	"encoding/json"
	"fmt"
)

// JsonConvert is a function type that converts any value to a pointer to JSON byte slice.
// If the conversion fails, it returns nil.
// This allows for custom conversion logic to be implemented and passed to conversion functions.
//
// Examples:
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//
//	result := ToJson(myCustomValue, customConverter)
type JsonConvert func(value interface{}) *[]byte

// ToJson converts any type of value to JSON byte slice.
// It uses ToJsonE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to JSON.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A JSON byte slice representation of the input value.
//
// Example:
//
//	b := ToJson(map[string]interface{}{"key": "value"})
//	fmt.Println(string(b)) // Output: {"key":"value"}
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//	customJson := ToJson(myCustomValue, customConverter)
func ToJson(value interface{}, converters ...JsonConvert) []byte {
	res, _ := ToJsonE(value, converters...)
	return res
}

// ToJsonOrDefault converts any type of value to JSON byte slice.
// If the conversion fails, it returns the provided default value.
//
// Parameters:
//   - value: The value to be converted to JSON.
//   - defaultValue: The default value to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A JSON byte slice representation of the input value or the default value if conversion fails.
//
// Example:
//
//	b := ToJsonOrDefault(map[string]interface{}{"key": "value"}, []byte("{}"))
//	fmt.Println(string(b)) // Output: {"key":"value"}
//
//	b = ToJsonOrDefault(nil, []byte("{}"))
//	fmt.Println(string(b)) // Output: {}
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonOrDefault(myCustomValue, []byte("{}"), customConverter)
func ToJsonOrDefault(value interface{}, defaultValue []byte, converters ...JsonConvert) []byte {
	if value == nil {
		return defaultValue
	}
	res, err := ToJsonE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToJsonE converts any type of value to JSON byte slice or returns an error.
//
// Parameters:
//   - value: The value to be converted to JSON.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - A JSON byte slice representation of the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	b, err := ToJsonE(map[string]interface{}{"key": "value"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(b)) // Output: {"key":"value"}
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//	customJson, err := ToJsonE(myCustomValue, customConverter)
func ToJsonE(value interface{}, converters ...JsonConvert) ([]byte, error) {
	if value == nil {
		return nil, fmt.Errorf("nil value cannot be converted to JSON")
	}
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	b, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ToJsonString converts any type of value to JSON string.
// It uses ToJsonStringE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to JSON string.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A JSON string representation of the input value.
//
// Example:
//
//	s := ToJsonString(map[string]interface{}{"key": "value"})
//	fmt.Println(s) // Output: {"key":"value"}
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				s := string(b)
//				return &s
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonString(myCustomValue, customConverter)
func ToJsonString(value interface{}, converters ...JsonConvert) string {
	res, _ := ToJsonStringE(value, converters...)
	return res
}

// ToJsonStringOrDefault converts any type of value to JSON string.
// If the conversion fails, it returns the provided default value.
//
// Parameters:
//   - value: The value to be converted to JSON string.
//   - defaultValue: The default value to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	A JSON string representation of the input value or the default value if conversion fails.
//
// Example:
//
//	s := ToJsonStringOrDefault(map[string]interface{}{"key": "value"}, "{}")
//	fmt.Println(s) // Output: {"key":"value"}
//
//	s = ToJsonStringOrDefault(nil, "{}")
//	fmt.Println(s) // Output: {}
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				s := string(b)
//				return &s
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonStringOrDefault(myCustomValue, "{}", customConverter)
func ToJsonStringOrDefault(value interface{}, defaultValue string, converters ...JsonConvert) string {
	if value == nil {
		return defaultValue
	}
	res, err := ToJsonStringE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToJsonStringE converts any type of value to JSON string or returns an error.
//
// Parameters:
//   - value: The value to be converted to JSON string.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - A JSON string representation of the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	s, err := ToJsonStringE(map[string]interface{}{"key": "value"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(s) // Output: {"key":"value"}
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.Marshal(v)
//			if err == nil {
//				s := string(b)
//				return &s
//			}
//		}
//		return nil
//	}
//	customJson, err := ToJsonStringE(myCustomValue, customConverter)
func ToJsonStringE(value interface{}, converters ...JsonConvert) (string, error) {
	b, err := ToJsonE(value, converters...)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToJsonIndent converts any type of value to indented JSON byte slice.
// It uses ToJsonIndentE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to indented JSON.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	An indented JSON byte slice representation of the input value.
//
// Example:
//
//	b := ToJsonIndent(map[string]interface{}{"key": "value"})
//	fmt.Println(string(b)) // Output: {
//	//   "key": "value"
//	// }
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.MarshalIndent(v, "", "  ")
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonIndent(myCustomValue, customConverter)
func ToJsonIndent(value interface{}, converters ...JsonConvert) []byte {
	res, _ := ToJsonIndentE(value, converters...)
	return res
}

// ToJsonIndentOrDefault converts any type of value to indented JSON byte slice.
// If the conversion fails, it returns the provided default value.
//
// Parameters:
//   - value: The value to be converted to indented JSON.
//   - defaultValue: The default value to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	An indented JSON byte slice representation of the input value or the default value if conversion fails.
//
// Example:
//
//	b := ToJsonIndentOrDefault(map[string]interface{}{"key": "value"}, []byte("{}"))
//	fmt.Println(string(b)) // Output: {
//	//   "key": "value"
//	// }
//
//	b = ToJsonIndentOrDefault(nil, []byte("{}"))
//	fmt.Println(string(b)) // Output: {}
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.MarshalIndent(v, "", "  ")
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonIndentOrDefault(myCustomValue, []byte("{}"), customConverter)
func ToJsonIndentOrDefault(value interface{}, defaultValue []byte, converters ...JsonConvert) []byte {
	if value == nil {
		return defaultValue
	}
	res, err := ToJsonIndentE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToJsonIndentE converts any type of value to indented JSON byte slice or returns an error.
//
// Parameters:
//   - value: The value to be converted to indented JSON.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - An indented JSON byte slice representation of the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	b, err := ToJsonIndentE(map[string]interface{}{"key": "value"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(b)) // Output: {
//	//   "key": "value"
//	// }
//
//	customConverter := func(value interface{}) *[]byte {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.MarshalIndent(v, "", "  ")
//			if err == nil {
//				return &b
//			}
//		}
//		return nil
//	}
//	customJson, err := ToJsonIndentE(myCustomValue, customConverter)
func ToJsonIndentE(value interface{}, converters ...JsonConvert) ([]byte, error) {
	if value == nil {
		return nil, fmt.Errorf("nil value cannot be converted to JSON")
	}

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	b, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ToJsonIndentString converts any type of value to indented JSON string.
// It uses ToJsonIndentStringE internally and ignores any error.
//
// Parameters:
//   - value: The value to be converted to indented JSON string.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	An indented JSON string representation of the input value.
//
// Example:
//
//	s := ToJsonIndentString(map[string]interface{}{"key": "value"})
//	fmt.Println(s) // Output: {
//	//   "key": "value"
//	// }
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.MarshalIndent(v, "", "  ")
//			if err == nil {
//				s := string(b)
//				return &s
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonIndentString(myCustomValue, customConverter)
func ToJsonIndentString(value interface{}, converters ...JsonConvert) string {
	res, _ := ToJsonIndentStringE(value, converters...)
	return res
}

// ToJsonIndentStringOrDefault converts any type of value to indented JSON string.
// If the conversion fails, it returns the provided default value.
//
// Parameters:
//   - value: The value to be converted to indented JSON string.
//   - defaultValue: The default value to return if conversion fails.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//
//	An indented JSON string representation of the input value or the default value if conversion fails.
//
// Example:
//
//	s := ToJsonIndentStringOrDefault(map[string]interface{}{"key": "value"}, "{}")
//	fmt.Println(s) // Output: {
//	//   "key": "value"
//	// }
//
//	s = ToJsonIndentStringOrDefault(nil, "{}")
//	fmt.Println(s) // Output: {}
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.MarshalIndent(v, "", "  ")
//			if err == nil {
//				s := string(b)
//				return &s
//			}
//		}
//		return nil
//	}
//	customJson := ToJsonIndentStringOrDefault(myCustomValue, "{}", customConverter)
func ToJsonIndentStringOrDefault(value interface{}, defaultValue string, converters ...JsonConvert) string {
	if value == nil {
		return defaultValue
	}
	res, err := ToJsonIndentStringE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToJsonIndentStringE converts any type of value to indented JSON string or returns an error.
//
// Parameters:
//   - value: The value to be converted to indented JSON string.
//   - converters: Optional custom converters to be applied before the default conversion.
//
// Returns:
//   - An indented JSON string representation of the input value.
//   - An error if the conversion fails.
//
// Example:
//
//	s, err := ToJsonIndentStringE(map[string]interface{}{"key": "value"})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(s) // Output: {
//	//   "key": "value"
//	// }
//
//	customConverter := func(value interface{}) *string {
//		if v, ok := value.(CustomType); ok {
//			b, err := json.MarshalIndent(v, "", "  ")
//			if err == nil {
//				s := string(b)
//				return &s
//			}
//		}
//		return nil
//	}
//	customJson, err := ToJsonIndentStringE(myCustomValue, customConverter)
func ToJsonIndentStringE(value interface{}, converters ...JsonConvert) (string, error) {
	b, err := ToJsonIndentE(value, converters...)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
