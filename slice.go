package convert

import (
	"encoding/json"
	"fmt"
	"time"
)

// SliceStringConverter is a function type for custom conversion of []string.
// It takes any value and returns a pointer to a []string if the conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
type SliceStringConverter func(value interface{}) *[]string

// SliceInterfaceConverter is a function type for custom conversion of []interface{}.
// It takes any value and returns a pointer to a []interface{} if the conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
type SliceInterfaceConverter func(value interface{}) *[]interface{}

// SliceBoolConverter is a function type for custom conversion of []bool.
// It takes any value and returns a pointer to a []bool if the conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
type SliceBoolConverter func(value interface{}) *[]bool

// SliceIntConverter is a function type for custom conversion of []int.
// It takes any value and returns a pointer to a []int if the conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
type SliceIntConverter func(value interface{}) *[]int

// SliceTimeConverter is a function type for custom conversion of []time.Time.
// It takes any value and returns a pointer to a []time.Time if the conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
type SliceTimeConverter func(value interface{}) *[]time.Time

// SliceDurationConverter is a function type for custom conversion of []time.Duration.
// It takes any value and returns a pointer to a []time.Duration if the conversion succeeds, or nil if it fails.
// This allows for flexible, user-defined conversion logic.
type SliceDurationConverter func(value interface{}) *[]time.Duration

// ToSliceString converts any type of value to []string.
// It takes a value of any type and a variable number of custom converters.
// If the conversion fails, it returns an empty slice.
func ToSliceString(value interface{}, converters ...SliceStringConverter) []string {
	res, _ := ToSliceStringE(value, converters...)
	return res
}

// ToSliceStringOrDefault converts any type of value to []string or returns a default value.
// It takes a value of any type, a default value of type []string, and a variable number of custom converters.
// If the input value is nil or if the conversion fails, it returns the default value.
func ToSliceStringOrDefault(value interface{}, defaultValue []string, converters ...SliceStringConverter) []string {
	if value == nil {
		return defaultValue
	}
	res, _ := ToSliceStringE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToSliceStringE converts any type of value to []string.
// It takes a value of any type and a variable number of custom converters.
// If the conversion succeeds, it returns the resulting []string and a nil error.
// If the conversion fails, it returns nil and an error describing the problem.
func ToSliceStringE(value interface{}, converters ...SliceStringConverter) ([]string, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case []string:
		return v, nil
	case []interface{}:
		res := make([]string, len(v))
		for i, val := range v {
			res[i] = ToString(val)
		}
		return res, nil
	case string:
		var res []string
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		var res []string
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// ToSliceInterface converts any type of value to []interface{}.
// It takes a value of any type and a variable number of custom converters.
// If the conversion fails, it returns an empty slice.
func ToSliceInterface(value interface{}, converters ...SliceInterfaceConverter) []interface{} {
	res, _ := ToSliceInterfaceE(value, converters...)
	return res
}

// ToSliceInterfaceOrDefault converts any type of value to []interface{} or returns a default value.
// It takes a value of any type, a default value of type []interface{}, and a variable number of custom converters.
// If the input value is nil or if the conversion fails, it returns the default value.
func ToSliceInterfaceOrDefault(value interface{}, defaultValue []interface{}, converters ...SliceInterfaceConverter) []interface{} {
	if value == nil {
		return defaultValue
	}
	res, _ := ToSliceInterfaceE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToSliceInterfaceE converts any type of value to []interface{}.
// It takes a value of any type and a variable number of custom converters.
// If the conversion succeeds, it returns the resulting []interface{} and a nil error.
// If the conversion fails, it returns nil and an error describing the problem.
func ToSliceInterfaceE(value interface{}, converters ...SliceInterfaceConverter) ([]interface{}, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case []interface{}:
		return v, nil
	case string:
		var res []interface{}
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		var res []interface{}
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// ToSliceBool converts any type of value to []bool.
// It takes a value of any type and a variable number of custom converters.
// If the conversion fails, it returns an empty slice.
func ToSliceBool(value interface{}, converters ...SliceBoolConverter) []bool {
	res, _ := ToSliceBoolE(value, converters...)
	return res
}

// ToSliceBoolOrDefault converts any type of value to []bool or returns a default value.
// It takes a value of any type, a default value of type []bool, and a variable number of custom converters.
// If the input value is nil or if the conversion fails, it returns the default value.
func ToSliceBoolOrDefault(value interface{}, defaultValue []bool, converters ...SliceBoolConverter) []bool {
	if value == nil {
		return defaultValue
	}
	res, _ := ToSliceBoolE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToSliceBoolE converts any type of value to []bool.
// It takes a value of any type and a variable number of custom converters.
// If the conversion succeeds, it returns the resulting []bool and a nil error.
// If the conversion fails, it returns nil and an error describing the problem.
func ToSliceBoolE(value interface{}, converters ...SliceBoolConverter) ([]bool, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	case []interface{}:
		res := make([]bool, len(v))
		for i, val := range v {
			boolVal, err := ToBoolE(val)
			if err != nil {
				return nil, fmt.Errorf("unable to convert element %d to bool: %v", i, err)
			}
			res[i] = boolVal
		}
		return res, nil
	case string:
		var res []bool
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		var res []bool
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// ToSliceInt converts any type of value to []int.
// It takes a value of any type and a variable number of custom converters.
// If the conversion fails, it returns an empty slice.
func ToSliceInt(value interface{}, converters ...SliceIntConverter) []int {
	res, _ := ToSliceIntE(value, converters...)
	return res
}

// ToSliceIntOrDefault converts any type of value to []int or returns a default value.
// It takes a value of any type, a default value of type []int, and a variable number of custom converters.
// If the input value is nil or if the conversion fails, it returns the default value.
func ToSliceIntOrDefault(value interface{}, defaultValue []int, converters ...SliceIntConverter) []int {
	if value == nil {
		return defaultValue
	}
	res, _ := ToSliceIntE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToSliceIntE converts any type of value to []int.
// It takes a value of any type and a variable number of custom converters.
// If the conversion succeeds, it returns the resulting []int and a nil error.
// If the conversion fails, it returns nil and an error describing the problem.
func ToSliceIntE(value interface{}, converters ...SliceIntConverter) ([]int, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	case []interface{}:
		res := make([]int, len(v))
		for i, val := range v {
			intVal, err := ToIntE(val)
			if err != nil {
				return nil, fmt.Errorf("unable to convert element %d to int: %v", i, err)
			}
			res[i] = intVal
		}
		return res, nil
	case string:
		var res []int
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		var res []int
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// ToSliceTime converts any type of value to []time.Time.
// It takes a value of any type and a variable number of custom converters.
// If the conversion fails, it returns an empty slice.
func ToSliceTime(value interface{}, converters ...SliceTimeConverter) []time.Time {
	res, _ := ToSliceTimeE(value, converters...)
	return res
}

// ToSliceTimeOrDefault converts any type of value to []time.Time or returns a default value.
// It takes a value of any type, a default value of type []time.Time, and a variable number of custom converters.
// If the input value is nil or if the conversion fails, it returns the default value.
func ToSliceTimeOrDefault(value interface{}, defaultValue []time.Time, converters ...SliceTimeConverter) []time.Time {
	if value == nil {
		return defaultValue
	}
	res, _ := ToSliceTimeE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToSliceTimeE converts any type of value to []time.Time.
// It takes a value of any type and a variable number of custom converters.
// If the conversion succeeds, it returns the resulting []time.Time and a nil error.
// If the conversion fails, it returns nil and an error describing the problem.
func ToSliceTimeE(value interface{}, converters ...SliceTimeConverter) ([]time.Time, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case []time.Time:
		return v, nil
	case []interface{}:
		res := make([]time.Time, len(v))
		for i, val := range v {
			timeVal, err := ToTimeE(val)
			if err != nil {
				return nil, fmt.Errorf("unable to convert element %d to time.Time: %v", i, err)
			}
			res[i] = timeVal
		}
		return res, nil
	case string:
		var res []time.Time
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		var res []time.Time
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// ToSliceDuration converts any type of value to []time.Duration.
// It takes a value of any type and a variable number of custom converters.
// If the conversion fails, it returns an empty slice.
func ToSliceDuration(value interface{}, converters ...SliceDurationConverter) []time.Duration {
	res, _ := ToSliceDurationE(value, converters...)
	return res
}

// ToSliceDurationOrDefault converts any type of value to []time.Duration or returns a default value.
// It takes a value of any type, a default value of type []time.Duration, and a variable number of custom converters.
// If the input value is nil or if the conversion fails, it returns the default value.
func ToSliceDurationOrDefault(value interface{}, defaultValue []time.Duration, converters ...SliceDurationConverter) []time.Duration {
	if value == nil {
		return defaultValue
	}
	res, _ := ToSliceDurationE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToSliceDurationE converts any type of value to []time.Duration.
// It takes a value of any type and a variable number of custom converters.
// If the conversion succeeds, it returns the resulting []time.Duration and a nil error.
// If the conversion fails, it returns nil and an error describing the problem.
func ToSliceDurationE(value interface{}, converters ...SliceDurationConverter) ([]time.Duration, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case []time.Duration:
		return v, nil
	case []interface{}:
		res := make([]time.Duration, len(v))
		for i, val := range v {
			durationVal, err := ToDurationE(val)
			if err != nil {
				return nil, fmt.Errorf("unable to convert element %d to time.Duration: %v", i, err)
			}
			res[i] = durationVal
		}
		return res, nil
	case string:
		var res []time.Duration
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case []byte:
		var res []time.Duration
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}
