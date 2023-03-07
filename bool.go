package convert

import (
	"fmt"
	"strings"
)

// ToBool converts any type of value to bool or returns error
func ToBool(value interface{}) (res bool, err error) {
	res = false
	err = nil
	switch v := value.(type) {
	case string:
		{
			str := ToValidString(value)
			switch strings.ToLower(str) {
			case "1", "t", "true", "y", "yes", "ok", "on":
				{
					res = true
				}
			case "0", "f", "false", "n", "no", "off", "":
				{
					res = false
				}
			default:
				err = fmt.Errorf("convert: %v (%v) to boolean failed", value, v)
			}
		}
	case bool:
		{
			res = value.(bool)
		}
	default:
		valueStr := ToValidString(value)
		res, err = ToBool(valueStr)
	}
	return
}

// ToBoolDef converts any type of value to bool or returns default value
func ToBoolDef(value interface{}, def bool) bool {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToBool(value)
	if err != nil {
		return def
	}
	return res
}

// ToBoolOrPanic converts any type of value to bool or panics
func ToBoolOrPanic(value interface{}) bool {
	if res, err := ToBool(value); err == nil {
		return res
	} else {
		panic(err)
	}
}
