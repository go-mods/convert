package convert

import (
	"time"

	"github.com/golang-module/carbon/v2"
)

// TimeConverter is a function type that converts any value to a pointer to time.Time.
// If the conversion fails, it returns nil.
// This allows for custom conversion logic to be implemented and passed to conversion functions.
//
// Examples:
//
//	customConverter := func(value interface{}) *time.Time {
//		if v, ok := value.(CustomType); ok {
//			t := v.CustomToTime()
//			return &t
//		}
//		return nil
//	}
//
//	result := ToTimeE(myCustomValue, customConverter)
type TimeConverter func(value interface{}) *time.Time

// DurationConverter is a function type that converts any value to a pointer to time.Duration.
// If the conversion fails, it returns nil.
// This allows for custom conversion logic to be implemented and passed to conversion functions.
//
// Examples:
//
//	customConverter := func(value interface{}) *time.Duration {
//		if v, ok := value.(CustomType); ok {
//			d := v.CustomToDuration()
//			return &d
//		}
//		return nil
//	}
//
//	result := ToDurationE(myCustomValue, customConverter)
type DurationConverter func(value interface{}) *time.Duration

// ToTime converts any type of value to time.Time, ignoring errors.
func ToTime(value interface{}, converters ...TimeConverter) time.Time {
	res, _ := ToTimeE(value, converters...)
	return res
}

// ToTimeOrDefault converts any type of value to time.Time or returns the provided default value if conversion fails.
func ToTimeOrDefault(value interface{}, defaultValue time.Time, converters ...TimeConverter) time.Time {
	res, err := ToTimeE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToTimeE converts any type of value to time.Time or returns an error.
func ToTimeE(value interface{}, converters ...TimeConverter) (time.Time, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case time.Time:
		return v, nil
	case *time.Time:
		return *v, nil
	case string:
		c := carbon.Parse(v)
		if c.Error != nil {
			return time.Time{}, c.Error
		}
		return c.StdTime(), nil
	default:
		valueStr := ToString(v)
		return ToTimeE(valueStr, converters...)
	}
}

// ToLayoutTime converts any type of value to time.Time with a layout applied, ignoring errors.
func ToLayoutTime(layout string, value interface{}, converters ...TimeConverter) time.Time {
	res, _ := ToLayoutTimeE(layout, value, converters...)
	return res
}

// ToLayoutTimeOrDefault converts any type of value to time.Time with a layout applied or returns the provided default value if conversion fails.
func ToLayoutTimeOrDefault(layout string, value interface{}, defaultValue time.Time, converters ...TimeConverter) time.Time {
	res, err := ToLayoutTimeE(layout, value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToLayoutTimeE converts any type of value to time.Time with a layout applied or returns an error.
func ToLayoutTimeE(layout string, value interface{}, converters ...TimeConverter) (time.Time, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case time.Time:
		return v, nil
	case *time.Time:
		return *v, nil
	case string:
		c := carbon.ParseByFormat(v, layout)
		if c.Error != nil {
			return time.Time{}, c.Error
		}
		return c.StdTime(), nil
	default:
		valueStr := ToString(v)
		c := carbon.ParseByFormat(valueStr, layout)
		if c.Error != nil {
			return time.Time{}, c.Error
		}
		return c.StdTime(), nil
	}
}

// ToTimeString converts a time.Time to a string with an optional format, ignoring errors.
func ToTimeString(t time.Time, format ...string) string {
	res, _ := ToTimeStringE(t, format...)
	return res
}

// ToTimeStringE converts a time.Time to a string with an optional format or returns an error.
func ToTimeStringE(t time.Time, format ...string) (string, error) {
	c := carbon.CreateFromStdTime(t)
	if c.Error != nil {
		return "", c.Error
	}
	if len(format) > 0 {
		r := t.Format(format[len(format)-1])
		if r != format[len(format)-1] {
			return r, nil
		}
		return c.Format(format[len(format)-1]), nil
	}
	return c.String(), nil
}

// ToDuration converts any type of value to time.Duration, ignoring errors.
func ToDuration(value interface{}, converters ...DurationConverter) time.Duration {
	res, _ := ToDurationE(value, converters...)
	return res
}

// ToDurationOrDefault converts any type of value to time.Duration or returns the provided default value if conversion fails.
func ToDurationOrDefault(value interface{}, defaultValue time.Duration, converters ...DurationConverter) time.Duration {
	res, err := ToDurationE(value, converters...)
	if err != nil {
		return defaultValue
	}
	return res
}

// ToDurationE converts any type of value to time.Duration or returns an error.
func ToDurationE(value interface{}, converters ...DurationConverter) (time.Duration, error) {
	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := value.(type) {
	case time.Duration:
		return v, nil
	case *time.Duration:
		return *v, nil
	case string:
		d, err := time.ParseDuration(v)
		if err != nil {
			return 0, err
		}
		return d, nil
	default:
		valueStr := ToString(v)
		return ToDurationE(valueStr, converters...)
	}
}
