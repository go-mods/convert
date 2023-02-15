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
	if val, err := ToBool(value); err == nil {
		return val
	}
	return def
}

// ToBoolOrPanic converts any type of value to bool or panics
func ToBoolOrPanic(value interface{}) bool {
	if res, err := ToBool(value); err == nil {
		return res
	} else {
		panic(err)
	}
}
